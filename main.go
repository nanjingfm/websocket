package main

import (
	"flag"
	"log"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	logic "github.com/nanjingfm/websocket/logic"
	"github.com/nanjingfm/websocket/route"
)

var addr = flag.String("addr", ":9000", "http logic address")

func main() {
	flag.Parse()
	logic.InitDefaultBucketGroup()
	logic.InitCounter()
	g := gin.Default()
	ginpprof.Wrap(g)
	g.Use(gin.Logger())

	route.RouteWs(g)
	route.RouteIApi(g)
	err := g.Run(*addr)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
