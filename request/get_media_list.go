package request

type GetMediaList struct {
	Base
	Schema string `json:"schema,omitempty"` // 筛选协议，例如 rtsp或rtmp
	Vhost  string `json:"vhost,omitempty"`  // 筛选虚拟主机，例如__defaultVhost__
	APP    string `json:"app,omitempty"`    // 筛选应用名，例如 live
	Stream string `json:"stream,omitempty"` // 筛选流id，例如 test
}
