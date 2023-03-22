package data

type Pagination struct {
	PageSize   uint        `json:"-"`
	PageNumber uint        `json:"-"`
	Offset     uint        `json:"-"`
	TotalPage  uint        `json:"total_page"`
	TotalRow   uint        `json:"total_row"`
	Items      interface{} `json:"items"`
}

func Paginate(pageSize, pageNumber, totalRows uint) *Pagination {
	offset := pageSize * (pageNumber - 1)
	pagination := &Pagination{
		PageSize:   pageSize,
		PageNumber: pageNumber,
		TotalRow:   totalRows,
		Offset:     offset,
		TotalPage:  (totalRows + pageSize - 1) / pageSize,
	}
	return pagination
}
