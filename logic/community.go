package logic

import (
	"Web_app/dao/mysql"
	"Web_app/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()

}
