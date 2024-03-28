package router

import (
	"crackpassword/controller"
	"crackpassword/middlewares"
	"crackpassword/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 ,page not exists!",
	})
}
func InitRouter() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	r.Use(middlewares.Cors())
	r.NoRoute(NotFound)
	//router.LoadHTMLFiles("./ui/index.html")
	//router.Static("/static", "./ui/static")
	//router.GET("/", func(c *gin.Context) {
	//	// c.JSON：返回JSON格式的数据
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})
	ctrl := &controller.Controller{}
	v1 := r.Group("api")
	{

		v1.GET("/disk-info", ctrl.DiskInfo)
		v1.GET("/history", ctrl.History)
		v1.POST("/crackps", ctrl.CrackPS)
	}
	err := r.Run(":" + utils.HttpPort)

	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
