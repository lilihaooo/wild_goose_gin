package request

import "fmt"

type PaginationReq struct {
	CurrentPage int          `json:"current_page"`
	PageSize    int          `json:"page_size"`
	Keyword     string       `json:"keyword"`
	Sort        *Sort        `json:"sort"`
	Conditions  *[]Condition `json:"condition"`
}
type Sort struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

type Condition struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NewPaginationReq 初始化分页请求, 设置默认值
func NewPaginationReq() *PaginationReq {
	return &PaginationReq{
		PageSize: 5, // 设置默认值为5
		Sort: &Sort{
			Field: "id",
			Order: "desc",
		},
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

func (p *PaginationReq) GetConditionStr() string {
	conditionStr := ""
	if p.Conditions != nil {
		if len(*p.Conditions) > 0 {
			for _, item := range *p.Conditions {
				conditionStr += fmt.Sprintf(item.Key) + "=" + fmt.Sprintf(item.Value) + " AND "
			}
			conditionStr = conditionStr[:len(conditionStr)-5]
		}
	}
	return conditionStr
}
