package request

type AddStreamProxy struct {
	Base
	Vhost  string `json:"vhost"`  // 添加的流的虚拟主机，例如__defaultVhost__
	APP    string `json:"app"`    // 添加的流的应用名，例如live
	Stream string `json:"stream"` // 添加的流的id名，例如test
	URL    string `json:"url"`    // 拉流地址，例如rtmp://live.hkstv.hk.lxdns.com/live/hks2

	RtpType      uint    `json:"rtp_type,omitempty"`       // rtp拉流时，拉流方式，0：tcp，1：udp，2：组播
	TimeoutSec   float32 `json:"timeout_sec,omitempty"`    // 拉流超时时间，单位秒，float类型
	RetryCount   int     `json:"retry_count,omitempty"`    // 拉流重试次数,不传此参数或传值<=0时，则无限重试
	EnableHls    uint    `json:"enable_hls,omitempty"`     // 是否转hls
	EnableMp4    uint    `json:"enable_mp4,omitempty"`     // 是否mp4录制
	EnableRtsp   uint    `json:"enable_rtsp,omitempty"`    // 是否转协议为rtsp/webrtc
	EnableRtmp   uint    `json:"enable_rtmp,omitempty"`    // 是否转协议为rtmp/flv
	EnableTs     uint    `json:"enable_ts,omitempty"`      // 是否转协议为http-ts/ws-ts
	EnableFmp4   uint    `json:"enable_fmp4,omitempty"`    // 是否转协议为http-fmp4/ws-fmp4
	EnableAudio  uint    `json:"enable_audio,omitempty"`   // 转协议是否开启音频
	AddMuteAudio uint    `json:"add_mute_audio,omitempty"` // 转协议无音频时，是否添加静音aac音频
	Mp4SavePath  string  `json:"mp4_save_path,omitempty"`  // mp4录制保存根目录，置空使用默认目录
	Mp4MaxSecond uint    `json:"mp4_max_second,omitempty"` // mp4录制切片大小，单位秒
	HlsSavePath  string  `json:"hls_save_path,omitempty"`  // hls保存根目录，置空使用默认目录
	ModifyStamp  uint    `json:"modify_stamp,omitempty"`   // 是否重新计算时间戳
}
