package controller

import (
	"gold/logic"
	"gold/model/request"

	"github.com/gin-gonic/gin"
)

type TradeUserController struct {
}

// 列出用户
func (t *TradeUserController) ListTradeUser(c *gin.Context) {
	req := new(request.ListTradeUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.ListTradeUser(c, req)
	})
}

func (t *TradeUserController) DeleteTradeUser(c *gin.Context) {
	req := new(request.DeleteTradeUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.DeleteTradeUser(c, req)
	})
}

func (t *TradeUserController) CreateTradeUser(c *gin.Context) {
	req := new(request.CreateTradeUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.CreateTradeUser(c, req)
	})
}

func (t *TradeUserController) UpdateTradeUser(c *gin.Context) {
	req := new(request.UpdateTradeUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.UpdateTradeUser(c, req)
	})
}
