package request

type GetProxyInfo struct {
	Base
	Key string `json:"key"` // 拉流唯一标识
}
