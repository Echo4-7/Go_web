package logic

import (
	"Web_app/dao/mysql"
	"Web_app/pkg/snowflake"
)

// 存放业务逻辑代码

func SignUp() {
	// 1. 判断用户是否存在
	mysql.QueryUserByUsername()
	// 2. 生成UID
	snowflake.GenID()
	// 3. 密码加密
	// 4. 保存进数据库
	mysql.InsertUser()
}
