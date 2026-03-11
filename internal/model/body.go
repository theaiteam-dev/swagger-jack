package model

import "fmt"

// ExtractBodyFlags parses an OpenAPI requestBody map into a slice of Flags
// with Source=FlagSourceBody. Flat object properties become individual flags;
// nested object properties are expanded using dot notation (e.g., "address.city").
//
// Navigation path: requestBody → content → "application/json" → schema → properties
//
// Returns (nil, nil) when requestBody is nil, and an empty slice (no error)
// when no applicable application/json schema is found.
func ExtractBodyFlags(requestBody map[string]interface{}) ([]Flag, error) {
	if requestBody == nil {
		return nil, nil
	}

	// Check for multipart/form-data first.
	if schema, ok := resolveMultipartSchema(requestBody); ok {
		return extractMultipartFlags(schema)
	}

	schema, ok := resolveJSONSchema(requestBody)
	if !ok {
		return []Flag{}, nil
	}

	return extractFlagsFromSchema(schema, "")
}

// resolveMultipartSchema navigates requestBody → content → "multipart/form-data" → schema
// and returns the schema map if found.
func resolveMultipartSchema(requestBody map[string]interface{}) (map[string]interface{}, bool) {
	content, ok := requestBody["content"].(map[string]interface{})
	if !ok {
		return nil, false
	}

	multipartContent, ok := content["multipart/form-data"].(map[string]interface{})
	if !ok {
		return nil, false
	}

	schema, ok := multipartContent["schema"].(map[string]interface{})
	if !ok {
		return nil, false
	}

	return schema, true
}

// extractMultipartFlags extracts flags from a multipart/form-data schema.
// Properties with type=string and format=binary become FlagTypeFile flags;
// all other properties use their normal mapped types.
func extractMultipartFlags(schema map[string]interface{}) ([]Flag, error) {
	properties, ok := schema["properties"].(map[string]interface{})
	if !ok {
		return []Flag{}, nil
	}

	requiredSet, err := buildRequiredSet(schema)
	if err != nil {
		return nil, err
	}

	flags := []Flag{}
	for key, rawProp := range properties {
		if key == "" {
			continue
		}
		prop, ok := rawProp.(map[string]interface{})
		if !ok {
			continue
		}

		propType, _ := prop["type"].(string)
		propFormat, _ := prop["format"].(string)

		var flagType FlagType
		if propType == "string" && propFormat == "binary" {
			flagType = FlagTypeFile
		} else {
			flagType = mapSchemaType(propType)
		}

		flags = append(flags, Flag{
			Name:     key,
			Type:     flagType,
			Required: requiredSet[key],
			Source:   FlagSourceBody,
		})
	}

	return flags, nil
}

// resolveJSONSchema navigates requestBody → content → "application/json" → schema
// and returns the schema map if found.
func resolveJSONSchema(requestBody map[string]interface{}) (map[string]interface{}, bool) {
	content, ok := requestBody["content"].(map[string]interface{})
	if !ok {
		return nil, false
	}

	jsonContent, ok := content["application/json"].(map[string]interface{})
	if !ok {
		return nil, false
	}

	schema, ok := jsonContent["schema"].(map[string]interface{})
	if !ok {
		return nil, false
	}

	return schema, true
}

// extractFlagsFromSchema recursively walks schema properties, producing flags
// with dot-notated names for nested objects. The prefix argument carries the
// parent key path (e.g., "address.") during recursion; it is empty at the top level.
func extractFlagsFromSchema(schema map[string]interface{}, prefix string) ([]Flag, error) {
	properties, ok := schema["properties"].(map[string]interface{})
	if !ok {
		return []Flag{}, nil
	}

	requiredSet, err := buildRequiredSet(schema)
	if err != nil {
		return nil, err
	}
	flags := []Flag{}

	for key, rawProp := range properties {
		if key == "" {
			continue // skip invalid empty property names
		}

		prop, ok := rawProp.(map[string]interface{})
		if !ok {
			continue
		}

		qualifiedName := prefix + key
		propType, _ := prop["type"].(string)

		if propType == "object" {
			if _, hasNested := prop["properties"]; hasNested {
				nested, err := extractFlagsFromSchema(prop, qualifiedName+".")
				if err != nil {
					return nil, err
				}
				flags = append(flags, nested...)
				continue
			}
		}

		var enumVals []string
		if rawEnum, ok := prop["enum"].([]interface{}); ok {
			for _, v := range rawEnum {
				enumVals = append(enumVals, fmt.Sprintf("%v", v))
			}
		}

		flags = append(flags, Flag{
			Name:     qualifiedName,
			Type:     mapSchemaType(propType),
			Required: requiredSet[key],
			Source:   FlagSourceBody,
			Enum:     enumVals,
		})
	}

	return flags, nil
}

// buildRequiredSet converts the schema "required" array into a set for O(1) lookup.
func buildRequiredSet(schema map[string]interface{}) (map[string]bool, error) {
	requiredSet := map[string]bool{}

	rawRequired, ok := schema["required"]
	if !ok {
		return requiredSet, nil
	}

	// The required field can be a []string (typed) or []interface{} (from JSON unmarshalling).
	switch required := rawRequired.(type) {
	case []string:
		for _, name := range required {
			requiredSet[name] = true
		}
	case []interface{}:
		for _, raw := range required {
			if name, ok := raw.(string); ok {
				requiredSet[name] = true
			}
		}
	default:
		return nil, fmt.Errorf("schema 'required' must be an array, got %T", rawRequired)
	}

	return requiredSet, nil
}
