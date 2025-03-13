package response

type GetProxyInfo struct {
	Base
	Data ProxyInfoData `json:"data"`
}

type ProxyInfoData struct {
	LiveSecs    int `json:"liveSecs"`
	RePullCount int `json:"rePullCount"`
	Status      int `json:"status"`
}
