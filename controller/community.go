package controller

import (
	"Web_app/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ------跟社区相关的------

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name）以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}