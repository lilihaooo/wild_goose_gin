package request

type PaginationReq struct {
	CurrentPage int          `json:"current_page"`
	PageSize    int          `json:"page_size"`
	Keyword     string       `json:"keyword"`
	Sort        *Sort        `json:"sort"`
	Conditions  *[]Condition `json:"condition"`
}

// NewPaginationReq 初始化分页请求, 设置默认值
func NewPaginationReq() *PaginationReq {
	return &PaginationReq{
		PageSize:   5, // 设置默认值为5
		Sort:       initSort(),
		Keyword:    "",
		Conditions: nil,
	}
}

func (p *PaginationReq) GetOffset() int {
	offset := (p.CurrentPage - 1) * p.PageSize
	if offset < 0 {
		return 0
	}
	return offset
}
