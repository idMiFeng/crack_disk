package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller 结构定义
type Controller struct{}

// DiskInfo 获取磁盘信息
func (c *Controller) DiskInfo(ctx *gin.Context) {
	// 处理获取磁盘信息的逻辑，这里假设返回一个 JSON 响应
	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取磁盘信息成功",
	})
}

// History 获取历史密码记录
func (c *Controller) History(ctx *gin.Context) {
	// 处理获取历史密码记录的逻辑，这里假设返回一个 JSON 响应
	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取历史密码记录成功",
	})
}

// CrackPS 破解密码
func (c *Controller) CrackPS(ctx *gin.Context) {
	// 处理破解密码的逻辑，这里假设从请求中获取密码并进行处理
	password := ctx.PostForm("password")
	// 在实际应用中，这里会有更复杂的逻辑，这里只是一个示例
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "破解密码成功",
		"password": password,
	})
}
