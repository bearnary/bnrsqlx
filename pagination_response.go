package bnrsqlx

type PaginationResponse struct {
	Page  int64 `json:"page"`
	Count int64 `json:"count"`
}
