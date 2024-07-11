package paginator

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Paginator struct {
	page        int
	pageSize    int
	maxPageSize int
}

func (p *Paginator) GetPage() int {
	page := p.page
	if p.page <= 0 {
		page = 1
	}
	return page
}

func (p *Paginator) GetPageSize() int {
	pageSize := p.pageSize
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	return pageSize
}

func (p *Paginator) GetLimit() int {
	return p.GetPageSize()
}

func (p *Paginator) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

func (p *Paginator) ToPaginatorResult(total int64, len int) *LengthAwarePaginator {
	totalItems := int(total)
	totalPages := totalItems / p.GetPageSize()
	if totalItems%p.GetPageSize() != 0 {
		totalPages++
	}

	from := p.GetOffset() + 1
	to := from + len - 1
	if to > totalItems {
		to = totalItems
	}

	return &LengthAwarePaginator{
		CurrentPage: p.GetPage(),
		PerPage:     p.GetPageSize(),
		From:        from,
		To:          to,
		LastPage:    totalPages,
		TotalItems:  totalItems,
	}
}

func FromGinContext(c *gin.Context, maxPageSize int) *Paginator {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	return &Paginator{
		page:        page,
		pageSize:    pageSize,
		maxPageSize: maxPageSize,
	}
}

type LengthAwarePaginator struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	From        int `json:"from"`
	To          int `json:"to"`
	LastPage    int `json:"last_page"`
	TotalItems  int `json:"total_items"`
}
