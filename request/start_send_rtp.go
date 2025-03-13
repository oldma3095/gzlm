package request

type StartSendRtp struct {
	Base
	Vhost   string `json:"vhost"`    // 虚拟主机，例如__defaultVhost__
	APP     string `json:"app"`      // 应用名，例如 live
	Stream  string `json:"stream"`   // 流id，例如 obs
	SSRC    uint   `json:"ssrc"`     // rtp推流的ssrc，ssrc不同时，可以推流到多个上级服务器
	DstUrl  string `json:"dst_url"`  // rtp打包采用ps还是es模式，默认采用ps模式，该参数非必选参数
	DstPort uint   `json:"dst_port"` // 目标端口
	IsUdp   uint   `json:"is_udp"`   // 是否为udp模式,否则为tcp模式

	SrcPort        uint   `json:"src_port,omitempty"`         // 指定tcp/udp客户端使用的本地端口，0时为随机端口，该参数非必选参数，不传时为随机端口。
	FromMp4        uint   `json:"from_mp4,omitempty"`         // 是否推送本地MP4录像，该参数非必选参数
	UsePs          uint   `json:"use_ps,omitempty"`           // rtp打包采用ps还是es模式，默认采用ps模式，该参数非必选参数
	PT             uint   `json:"pt,omitempty"`               // rtp payload type，默认96，该参数非必选参数
	OnlyAudio      uint   `json:"only_audio,omitempty"`       // rtp es方式打包时，是否只打包音频，该参数非必选参数
	UdpRtcpTimeout uint   `json:"udp_rtcp_timeout,omitempty"` // udp方式推流时，是否开启rtcp发送和rtcp接收超时判断，开启后(默认关闭)，如果接收rr rtcp超时，将导致主动停止rtp发送
	RecvStreamId   string `json:"recv_stream_id,omitempty"`   // 发送rtp同时接收，一般用于双向语言对讲, 如果不为空，说明开启接收，值为接收流的id
}
