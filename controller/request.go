package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const ContextUserIdKey = "userId"

var ErrorUserNotLogin = errors.New("用户未登录")

// getRequestUser 获取当前登陆用户ID
func getRequestUserID(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(ContextUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
