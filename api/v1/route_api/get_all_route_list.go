package route_api

import (
	"github.com/gin-gonic/gin"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
)

type GetAllRouteListRes struct {
	ID     uint   `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (RouteApi) GetAllRouteList(c *gin.Context) {
	var model models.Route
	routes, err := model.GetRouteList()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	res := []GetAllRouteListRes{}
	for _, item := range routes {
		one := GetAllRouteListRes{
			ID:     item.ID,
			Path:   item.Path,
			Method: item.Method,
		}
		res = append(res, one)
	}
	response.OkWithData(c, res)
}
