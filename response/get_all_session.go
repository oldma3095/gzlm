package response

type GetAllSession struct {
	Base
	Data []Session `json:"data"`
}

type Session struct {
	Id        string `json:"id"`         // 该tcp链接唯一id
	LocalIp   string `json:"local_ip"`   // 本机网卡ip
	LocalPort int    `json:"local_port"` // 本机端口号(这是个rtmp播放器或推流器)
	PeerIp    string `json:"peer_ip"`    // 客户端ip
	PeerPort  int    `json:"peer_port"`  // 客户端端口号
	Typeid    string `json:"typeid"`     // 客户端TCPSession typeid
}
