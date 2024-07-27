package controller

import (
	"Web_app/logic"
	"Web_app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	var p models.ParamSignup
	// 使用 ShouldBindJSON 将请求体中的 JSON 数据绑定到 p 结构体
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误
		zap.L().Error("Signup with invalid param", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//		//"msg": err.Error(),
	//	})
	//	return
	//}

	// 手动对请求参数进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("Signup with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	//fmt.Println(p)

	// 2. 业务处理
	if err := logic.SignUp(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败！",
		})
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func LoginHandler(c *gin.Context) {
	// 1. 获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断errs是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})
		return
	}
	// 2. 业务逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed!", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "登陆失败！",
		})
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success！",
	})
}
