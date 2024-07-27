package router

import (
	"Web_app/controller"
	"Web_app/logger"
	"Web_app/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	// 登陆业务路由
	r.POST("/login", controller.LoginHandler)
	r.GET("/ping", func(c *gin.Context) {
		//c.String(http.StatusOK, "pong")
		c.JSON(http.StatusOK, gin.H{
			"version": settings.Conf.Version,
		})
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
