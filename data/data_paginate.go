package data

type Pagination struct {
	PageSize   int         `json:"-"`
	PageNumber int         `json:"-"`
	Offset     int         `json:"-"`
	TotalPage  int64       `json:"total_page"`
	TotalRow   int64       `json:"total_row"`
	Items      interface{} `json:"items"`
}

func Paginate(pageSize, pageNumber int, totalRows int64) *Pagination {
	offset := pageSize * (pageNumber - 1)
	pagination := &Pagination{
		PageSize:   pageSize,
		PageNumber: pageNumber,
		TotalRow:   totalRows,
		Offset:     offset,
		TotalPage:  (totalRows + int64(pageSize) - 1) / int64(pageSize),
	}
	return pagination
}
