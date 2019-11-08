package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nanjingfm/websocket/logic"
)

func Ws(c *gin.Context) {
	uid := c.Query("uid")
	logic.ServeWs(logic.GetDefaultBucketGroup(), c.Writer, c.Request, uid)
}
