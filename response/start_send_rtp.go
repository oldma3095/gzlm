package response

type StartSendRtp struct {
	Base
	LocalPort int `json:"local_port"` // 使用的本地端口号
}
