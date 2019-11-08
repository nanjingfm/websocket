package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nanjingfm/websocket/service"
)

//内部接口
func RouteIApi(engine *gin.Engine) {
	engine.GET("/", service.Test)
	engine.GET("/send", service.ApiSendMsg)
	engine.GET("/stats", service.ApiServerStats)
}

func RouteWs(engine *gin.Engine) {
	engine.GET("/ws", service.Ws)
}
