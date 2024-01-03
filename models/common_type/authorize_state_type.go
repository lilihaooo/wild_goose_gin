package common_type

import "encoding/json"

type AuthorizeStageType int

const (
	AuthorizeStageNormal AuthorizeStageType = 1 // 正常
	AuthorizeStageStop   AuthorizeStageType = 2 // 暂停
)

func (s AuthorizeStageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s AuthorizeStageType) String() bool {
	var res bool
	switch s {
	case AuthorizeStageNormal:
		res = true
	case AuthorizeStageStop:
		res = false
	}
	return res
}
