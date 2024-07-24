package router

import (
	"Web_app/controller"
	"Web_app/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r

	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "ok")
	//})
	//return r
}
