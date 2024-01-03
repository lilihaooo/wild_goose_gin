package common_type

import "encoding/json"

type TaskNodeType int

const (
	TaskUnStart      TaskNodeType = 1 // 开工
	TaskStart        TaskNodeType = 2 // 开工
	TaskPreCheck     TaskNodeType = 3 // 预检
	TaskFaultInspect TaskNodeType = 4 // 故检
	TaskEndTest      TaskNodeType = 5 // 测试
	TaskWindUp       TaskNodeType = 6 // 收尾
	TaskPause        TaskNodeType = 7 // 暂停
)

func (t TaskNodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t TaskNodeType) String() string {
	var str string
	switch t {
	case TaskUnStart:
		str = "未开工"
	case TaskStart:
		str = "开工"
	case TaskPreCheck:
		str = "预检"
	case TaskFaultInspect:
		str = "故检"
	case TaskEndTest:
		str = "测试"
	case TaskWindUp:
		str = "收尾"
	case TaskPause:
		str = "暂停"
	}
	return str
}
