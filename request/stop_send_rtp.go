package request

type StopSendRtp struct {
	Base
	Vhost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	APP    string `json:"app"`    // 应用名，例如 live
	Stream string `json:"stream"` // 流id，例如 obs

	SSRC uint `json:"ssrc,omitempty"` // 根据ssrc关停某路rtp推流，不传时关闭所有推流
}
