package request

type KickSessions struct {
	Base
	LocalPort string `json:"local_port,omitempty"` // 筛选本机端口，例如筛选rtsp链接：554
	PeerIP    string `json:"peer_ip,omitempty"`    // 筛选客户端ip
}
