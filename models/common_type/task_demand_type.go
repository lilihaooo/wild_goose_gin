package common_type

import "encoding/json"

type TaskDemandType int

const (
	Inspect  TaskDemandType = 1 // 检测
	Repair   TaskDemandType = 2 // 修理
	Overhaul TaskDemandType = 3 // 翻修
	Claim    TaskDemandType = 4 // 索赔
)

func (t TaskDemandType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t TaskDemandType) String() string {
	var str string
	switch t {
	case Inspect:
		str = "检测"
	case Repair:
		str = "修理"
	case Overhaul:
		str = "翻修"
	case Claim:
		str = "索赔"
	}
	return str
}
