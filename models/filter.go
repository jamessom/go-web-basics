package models

import "math"

type Filter struct {
	Page, PageSize int
}

type Metadata struct {
	CurrentPage, PageSize, FirstPage, LastPage, TotalRecords int
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