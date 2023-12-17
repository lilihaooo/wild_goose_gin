package common_type

type MenuType uint

const (
	IsMenu   MenuType = 1
	IsFolder MenuType = 2
)

//func (m MenuType) MarshalJSON() ([]byte, error) {
//	return json.Marshal(m.String())
//}

//func (m MenuType) String() string {
//	var str string
//	switch m {
//	case IsFolder:
//		str = "菜单夹"
//	case IsMenu:
//		str = "菜单"
//	}
//	return str
//}
