package paginator

import (
	"errors"
)

type Paginator struct {
	Page       int `json:"page"`
	Limit      int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func Populate(page, limit, total int) (*Paginator, error) {
	paginator := &Paginator{
		Page:  page,
		Limit: limit,
		Total: total,
	}

	if totalPage, err := ResolveTotalPages(total, limit); err == nil {
		paginator.TotalPages = totalPage
		return paginator, nil
	} else {
		return nil, err
	}
}

func ResolveOffset(page, limit int) (int, error) {
	if page == 0 {
		return 0, errors.New("page must greater than 0")
	}
	return (page - 1) * limit, nil
}

func ResolveTotalPages(total, limit int) (int, error) {

	if total == 0 {
		return 0, nil
	}

	if limit == 0 {
		return 0, errors.New("limit must greater than 0")
	}

	totalPages := total / limit

	if total%limit > 0 {
		totalPages++
	}

	return totalPages, nil
}
