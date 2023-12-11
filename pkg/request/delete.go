package request

type DeleteRequest struct {
	IDs []int64 `json:"ids" validate:"required"`
}
