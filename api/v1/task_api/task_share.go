package task_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"wild_goose_gin/global"
	"wild_goose_gin/pkg/response"
)

type TaskShareReq struct {
	ID     uint `json:"id" validate:"required"`
	UserID uint `json:"user_id" validate:"required"`
}

func (TaskApi) TaskShare(c *gin.Context) {
	var req []TaskShareReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, err.Error()) // todo 去掉真实的error
		return
	}
	if len(req) == 0 {
		response.FailWithMsg(c, response.INVALID_PARAMS, "请选择")
		return
	}

	var caseClauses []string
	for _, one := range req {
		var userID string
		if one.UserID == 0 {
			userID = "null"
		} else {
			userID = strconv.Itoa(int(one.UserID))
		}
		// 使用 fmt.Sprintf 构造每个 WHEN THEN 子句
		clause := fmt.Sprintf("WHEN id = %v THEN %v", one.ID, userID)
		// 将每个子句添加到切片中
		caseClauses = append(caseClauses, clause)
	}
	// 将所有子句拼接成一个完整的 CASE 语句
	sql := "UPDATE task SET user_id =  CASE " + strings.Join(caseClauses, " ") + " ELSE user_id END, share = CASE WHEN user_id IS NOT NULL THEN 2 ELSE 1 END"
	query := global.DB.Exec(sql)
	if query.Error != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "执行失败") // todo 错误
		return
	}
	if query.RowsAffected == 0 {
		response.FailWithMsg(c, response.FAIL_OPER, "没有数据更新") // todo 错误
		return
	}
	//response.OkWithMsg(c, "添加成功")
	response.OkWithData(c, sql)
}
