package main

import (
	"git.jiayincloud.com/TestDev/pegasus.git/core"
	"log"
	"git.jiayincloud.com/TestDev/pegasus.git/services"
	"flag"
	"fmt"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	port := flag.String("port", "60006", "server port")
	flag.Parse()
	httpGroup := core.Router.Group("/http")
	{
		httpGroup.POST("/load", services.Load)
		httpGroup.POST("/fire", services.Fire)
		httpGroup.GET("/ammos", services.GetAmmos)
		httpGroup.DELETE("/ammos/drop", services.DropAmmos)
	}
	core.Router.Run(fmt.Sprintf(":%s", *port))
}
