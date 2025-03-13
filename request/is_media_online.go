package request

type IsMediaOnline struct {
	Base
	Schema string `json:"schema"` // 协议，例如 rtsp或rtmp
	Vhost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	APP    string `json:"app"`    // 应用名，例如 live
	Stream string `json:"stream"` // 流id，例如 test
}
