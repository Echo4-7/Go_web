package redis

// redis key
// redis key 注意使用，命名空间的方式，方便查询和拆分（如下使用冒号）
const (
	KeyPrefix              = "Web_app:"
	KeyPostTimeZSet        = "post:time"   // zset;帖子及发帖时间 （其中zset是redis的一种数据类型）
	KeyPostScoreZSet       = "post:score"  // zset;帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted:" // zset;记录用户及投票类型;参数是post_id
)

// 给redis key 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
