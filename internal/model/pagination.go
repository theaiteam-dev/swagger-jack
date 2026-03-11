package model

// PaginationType identifies the pagination style used by an endpoint.
type PaginationType string

const (
	// PaginationPageBased indicates page/per_page style pagination.
	PaginationPageBased PaginationType = "page"
	// PaginationOffsetBased indicates limit/offset style pagination.
	PaginationOffsetBased PaginationType = "offset"
	// PaginationCursorBased indicates cursor/after/before style pagination.
	PaginationCursorBased PaginationType = "cursor"
)

// Pagination describes the pagination style detected for a Command.
type Pagination struct {
	// Type is the pagination style.
	Type PaginationType
	// PageParam is the page/offset parameter name (e.g. "page" or "offset").
	PageParam string
	// SizeParam is the page-size/limit parameter name (e.g. "per_page" or "limit").
	SizeParam string
	// CursorParam is the cursor parameter name (e.g. "cursor", "after", "before").
	CursorParam string
}

// detectPagination scans a list of query parameter names and returns a
// Pagination value if a known pagination pattern is found, or nil otherwise.
func detectPagination(queryParamNames []string) *Pagination {
	params := make(map[string]bool, len(queryParamNames))
	for _, n := range queryParamNames {
		params[n] = true
	}

	// Cursor-based: cursor / after / before take highest priority.
	for _, cursorKey := range []string{"cursor", "after", "before"} {
		if params[cursorKey] {
			p := &Pagination{
				Type:        PaginationCursorBased,
				CursorParam: cursorKey,
			}
			if params["limit"] {
				p.SizeParam = "limit"
			}
			return p
		}
	}

	// Offset-based: limit + offset.
	if params["limit"] && params["offset"] {
		return &Pagination{
			Type:      PaginationOffsetBased,
			SizeParam: "limit",
			PageParam: "offset",
		}
	}

	// Page-based: page + one of per_page / perPage / per-page / page_size / pageSize.
	if params["page"] {
		for _, sizeKey := range []string{"per_page", "perPage", "per-page", "page_size", "pageSize"} {
			if params[sizeKey] {
				return &Pagination{
					Type:      PaginationPageBased,
					PageParam: "page",
					SizeParam: sizeKey,
				}
			}
		}
	}

	return nil
}
