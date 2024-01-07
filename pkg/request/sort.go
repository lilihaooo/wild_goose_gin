package request

type Sort struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

func initSort() *Sort {
	return &Sort{
		Field: "id",
		Order: "desc",
	}
}
