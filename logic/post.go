package logic

import (
	"Web_app/dao/mysql"
	"Web_app/models"
	"Web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成post_id
	p.ID = snowflake.GenID()
	// 2. 保存到数据库
	return mysql.CreatePost(p)
	// 3.
}
