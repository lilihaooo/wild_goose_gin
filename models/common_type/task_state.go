package common_type

import "encoding/json"

type TaskNodeType int

const (
	TaskStart        TaskNodeType = 1 // 开工
	TaskPreCheck     TaskNodeType = 2 // 预检
	TaskFaultInspect TaskNodeType = 3 // 故检
	TaskEndTest      TaskNodeType = 4 // 测试
	TaskWindUp       TaskNodeType = 5 // 收尾
	TaskPause        TaskNodeType = 6 // 暂停
)

func (t TaskNodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t TaskNodeType) String() string {
	var str string
	switch t {
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
