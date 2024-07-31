package route

import (
	"gold/config"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Conf.System.Mode)
	r := gin.Default()
	group := r.Group("/" + config.Conf.System.UrlPathPrefix)

	InitUserRouter(group)

	return r
}
