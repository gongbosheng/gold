package route

import (
	"gold/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(group *gin.RouterGroup) {
	tradeUser := group.Group("/trade")
	{
		tradeUser.GET("/userList", controller.TradeUser.ListTradeUser)
		tradeUser.POST("/userCreate", controller.TradeUser.CreateTradeUser)
		tradeUser.POST("/userUpdate", controller.TradeUser.UpdateTradeUser)
		tradeUser.POST("/userDelete", controller.TradeUser.DeleteTradeUser)
	}
}
