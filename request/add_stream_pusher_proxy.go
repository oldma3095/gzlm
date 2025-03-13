package request

type AddStreamPusherProxy struct {
	Base
	Schema string `json:"schema"`  // 推流协议，支持rtsp、rtmp，大小写敏感
	Vhost  string `json:"vhost"`   // 已注册流的虚拟主机，一般为__defaultVhost__
	APP    string `json:"app"`     // 已注册流的应用名，例如live
	Stream string `json:"stream"`  // 已注册流的id名，例如test
	DstUrl string `json:"dst_url"` // 推流地址，需要与schema字段协议一致

	RtpType    uint    `json:"rtp_type,omitempty"`    // rtsp推流时，推流方式，0：tcp，1：udp
	TimeoutSec float32 `json:"timeout_sec,omitempty"` // 推流超时时间，单位秒，float类型
	RetryCount int     `json:"retry_count,omitempty"` // 推流重试次数,不传此参数或传值<=0时，则无限重试
}
