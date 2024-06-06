package bnrsqlx

import "fmt"

type PaginationRequest struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

func (r *PaginationRequest) ConcatPagination(queryStr string) string {
	if r.Page == 0 || r.Limit == 0 {
		return queryStr
	}
	if queryStr != "" {
		queryStr = fmt.Sprintf("%v&", queryStr)
	}
	queryStr = fmt.Sprintf("%vpage=%v&limit=%v", queryStr, r.Page, r.Limit)
	return queryStr
}
