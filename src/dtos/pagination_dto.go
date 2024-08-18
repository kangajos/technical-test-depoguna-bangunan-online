package dtos

import "strconv"

type Pagination struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func SetPaginationConfig(page, limit string) Pagination {
	var p Pagination
	pages, _ := strconv.Atoi(page)
	if pages < 1 {
		pages = 1
	}
	p.Page = pages

	limits, _ := strconv.Atoi(limit)
	if limits < 1 {
		limits = 10
	}

	p.Limit = limits

	offset := (p.Page - 1) - p.Limit
	p.Offset = offset
	return p
}
