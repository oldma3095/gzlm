package response

type GetProxyPusherInfo struct {
	Base
	Data ProxyPusherInfoData `json:"data"`
}

type ProxyPusherInfoData struct {
	LiveSecs       int `json:"liveSecs"`
	RePublishCount int `json:"rePublishCount"`
	Status         int `json:"status"`
}
