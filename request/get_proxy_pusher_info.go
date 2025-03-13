package request

type GetProxyPusherInfo struct {
	Base
	Key string `json:"key"` // 推流流唯一标识
}
