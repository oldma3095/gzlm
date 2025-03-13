package request

type OpenRtpServer struct {
	Base
	Port     uint   `json:"port"`      // 绑定的端口，0时为随机端口
	TcpMode  uint   `json:"tcp_mode"`  // tcp模式，0时为不启用tcp监听，1时为启用tcp监听，2时为tcp主动连接模式
	StreamId string `json:"stream_id"` // 该端口绑定的流id

	ReUsePort uint `json:"re_use_port,omitempty"` // 是否重用端口，默认为0，非必选参数
	SSRC      uint `json:"ssrc,omitempty"`        // 是否指定收流的rtp ssrc, 十进制数字，不指定或指定0时则不过滤rtp，非必选参数
	OnlyAudio uint `json:"only_audio,omitempty"`  // 是否为单音频track，用于语音对讲
}
