package common_type

import "encoding/json"

type TaskShareType int

const (
	UnShared TaskShareType = 1 // 未分配
	Shared   TaskShareType = 2 // 已分配
)

func (t TaskShareType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t TaskShareType) String() string {
	var str string
	switch t {
	case Shared:
		str = "已分配"
	case UnShared:
		str = "未分配"
	}
	return str
}
