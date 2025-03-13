package response

type KickSessions struct {
	Base
	CountHit int `json:"count_hit"` // 筛选命中客户端个数
}
