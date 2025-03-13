package response

type StartRecord struct {
	Base
	Result bool `json:"result"` // 成功与否
}
