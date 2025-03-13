package response

type IsMediaOnline struct {
	Base
	Online bool `json:"online"` // 是否在线
}
