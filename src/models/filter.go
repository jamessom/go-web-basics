package models

import "math"

type Filter struct {
	Page, PageSize int
}

type Metadata struct {
	CurrentPage int `json:"current_page"`
	PageSize int `json:"page_size"`
	FirstPage int `json:"first_page"`
	LastPage int `json:"last_page"`
	TotalRecords int `json:"total_records"`
}

func (filter Filter) Limit() int {
	return filter.PageSize
}

func (filter Filter) Offset() int {
	return (filter.Page - 1) * filter.PageSize
}

func ComputeMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage: page,
		PageSize: pageSize,
		FirstPage: 1,
		LastPage: int(
			math.Ceil(float64(totalRecords) / float64(pageSize)),
		),
		TotalRecords: totalRecords,
	}
}