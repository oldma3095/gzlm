package response

type StopRecord struct {
	Base
	Result bool `json:"result"` // 成功与否
}
