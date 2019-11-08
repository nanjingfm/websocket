package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nanjingfm/websocket/logic"
)

func Test(c *gin.Context) {
	c.File("home.html")
}

func ApiSendMsg(c *gin.Context) {
	to := c.Query("to")
	msg := c.Query("msg")
	logic.GetDefaultBucketGroup().GetHub(to).Send(&logic.Msg{
		IdentifyCode: to,
		Msg:          []byte(msg),
	})
}

func ApiServerStats(c *gin.Context) {
	c.JSON(200, logic.GetDefaultCounter())
}
