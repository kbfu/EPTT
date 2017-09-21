package core

import (
"github.com/gin-contrib/cors"
"github.com/gin-gonic/gin"
)

var Router = router()

func router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	return r
}
