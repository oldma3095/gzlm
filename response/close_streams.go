package response

type CloseStreams struct {
	Base
	CountHit    int `json:"count_hit"`    // 筛选命中的流个数
	CountClosed int `json:"count_closed"` // 被关闭的流个数，可能小于count_hit
}
