package response

type GetRtpInfo struct {
	Base
	Exist     bool   `json:"exist"`      // 是否存在
	PeerIp    string `json:"peer_ip"`    // 推流客户端ip
	PeerPort  int    `json:"peer_port"`  // 客户端端口号
	LocalIp   string `json:"local_ip"`   // 本地监听的网卡ip
	LocalPort int    `json:"local_port"` // 本地监听的端口号
}
