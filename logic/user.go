package logic

import (
	"Web_app/dao/mysql"
	"Web_app/models"
	"Web_app/pkg/snowflake"
)

// 存放业务逻辑代码

func SignUp(p *models.ParamSignup) (err error) {
	// 1. 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}

	// 2. 生成UID
	userId := snowflake.GenID()
	// 构造User实例
	user := &models.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 3. 保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) error {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
