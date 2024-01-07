package request

import "fmt"

type Condition struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func GetConditionStr(conditions []Condition) string {
	conditionStr := ""
	if len(conditions) > 0 {
		for _, item := range conditions {
			conditionStr += fmt.Sprintf(item.Key) + "=" + fmt.Sprintf(item.Value) + " AND "
		}
		conditionStr = conditionStr[:len(conditionStr)-5] //去掉最后的额and
	}
	return conditionStr
}
