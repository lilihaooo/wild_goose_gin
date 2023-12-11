package common_type

import "encoding/json"

type ComponentStageType int

const (
	ComponentStageNormal ComponentStageType = 1 // 正常
	ComponentStageStop   ComponentStageType = 2 // 暂停
)

func (s ComponentStageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ComponentStageType) String() bool {
	var res bool
	switch s {
	case ComponentStageNormal:
		res = true
	case ComponentStageStop:
		res = false
	}
	return res
}
