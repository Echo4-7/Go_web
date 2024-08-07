package models

import "time"

// 结构体内存对齐 类型一致的尽量放在一起，缩小整个结构体字段的内存大小

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName       string                    `json:"author_name"`
	*Post                                      // 嵌入帖子结构体
	*CommunityDetail `json:"community_detail"` // 嵌入社区信息
}
