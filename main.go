package main

import (
	"github.com/kbfu/pegasus/core"
	"log"
	"github.com/kbfu/pegasus/services"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	httpGroup := core.Router.Group("/http")
	{
		httpGroup.POST("/load", services.Load)
		httpGroup.POST("/fire", services.Fire)
	}
	core.Router.Run(":60006")
}
