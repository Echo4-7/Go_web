package models

// 定义请求的参数结构体

// ParamSignup 注册请求参数
type ParamSignup struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登陆请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前用户
	PostID    string `json:"post_id,string" binding:"required"`                // 帖子ID
	Direction int8   `json:"direction,string" binding:"required,oneof=1 0 -1"` // 赞成票（1）反对票（-1）取消投票（0）
}
