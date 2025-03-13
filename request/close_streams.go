package request

type CloseStreams struct {
	Base
	Schema string `json:"schema,omitempty"` // 协议，例如 rtsp或rtmp
	Vhost  string `json:"vhost,omitempty"`  // 虚拟主机，例如__defaultVhost__
	APP    string `json:"app,omitempty"`    // 应用名，例如 live
	Stream string `json:"stream,omitempty"` // 流id，例如 test
	Force  uint   `json:"force"`            // 是否强制关闭(有人在观看是否还关闭)
}
