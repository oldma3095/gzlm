package gzlm

type MediaInfo struct {
	AliveSecond      int         `json:"aliveSecond"`      // 存活时间，单位秒
	App              string      `json:"app"`              // 应用名
	BytesSpeed       int         `json:"bytesSpeed"`       // 数据产生速度，单位byte/s
	CreateStamp      int64       `json:"createStamp"`      // GMT unix系统时间戳，单位秒
	IsRecordingHLS   bool        `json:"isRecordingHLS"`   // 是否正在录制ts
	IsRecordingMP4   bool        `json:"isRecordingMP4"`   // 是否正在录制MP4
	OriginSock       OriginSock  `json:"originSock"`       // 客户端和服务器网络信息，可能为null类型
	OriginType       int         `json:"originType"`       // 产生源类型，包括 unknown = 0,rtmp_push=1,rtsp_push=2,rtp_push=3,pull=4,ffmpeg_pull=5,mp4_vod=6,device_chn=7
	OriginTypeStr    string      `json:"originTypeStr"`    // 产生源类型
	OriginUrl        string      `json:"originUrl"`        // 产生源的url
	ReaderCount      int         `json:"readerCount"`      // 本协议观看人数
	Schema           string      `json:"schema"`           // 协议
	Stream           string      `json:"stream"`           // 流id
	TotalReaderCount int         `json:"totalReaderCount"` // 观看总人数，包括hls/rtsp/rtmp/http-flv/ws-flv
	Tracks           []Track     `json:"tracks"`           // 音视频轨道
	Vhost            string      `json:"vhost"`            // 虚拟主机名
	Transcodes       []Transcode `json:"transcode"`        // pro版本，转码相关
}

type OriginSock struct {
	Identifier string `json:"identifier"`
	LocalIp    string `json:"local_ip"`
	LocalPort  int    `json:"local_port"`
	PeerIp     string `json:"peer_ip"`
	PeerPort   int    `json:"peer_port"`
}

type Track struct {
	CodecId       int     `json:"codec_id"`                  // H264 = 0, H265 = 1, AAC = 2, G711A = 3, G711U = 4
	CodecIdName   string  `json:"codec_id_name"`             // 编码类型名称
	CodecType     int     `json:"codec_type"`                // Video = 0, Audio = 1
	Fps           float32 `json:"fps,omitempty"`             // 视频fps
	Frames        int     `json:"frames"`                    // 累计接收帧数
	KeyFrames     int     `json:"key_frames,omitempty"`      // 累计接收关键帧数
	Ready         bool    `json:"ready"`                     // 轨道是否准备就绪
	Height        int     `json:"height,omitempty"`          // 视频高
	Width         int     `json:"width,omitempty"`           // 视频宽
	Channels      int     `json:"channels,omitempty"`        // 音频通道数
	SampleBit     int     `json:"sample_bit,omitempty"`      // 音频采样位数
	SampleRate    int     `json:"sample_rate,omitempty"`     // 音频采样率
	Loss          float32 `json:"loss,omitempty"`            // -1
	GopIntervalMs int     `json:"gop_interval_ms,omitempty"` // gop间隔时间，单位毫秒
	GopSize       int     `json:"gop_size,omitempty"`        // gop大小，单位帧数
}

type Transcode struct {
	Name    string           `json:"name"`     // 转码名称
	Setting TranscodeSetting `json:"setting"`  // 转码配置信息
	Adec    string           `json:"adec"`     // ⾳频解码器名称
	Aenc    string           `json:"aenc"`     // ⾳频编码器名称
	AencCtx TranscodeAencCtx `json:"aenc_ctx"` // ⾳频AVCodecContext信息
	Vdec    string           `json:"vdec"`     // 视频解码器名称
	Venc    string           `json:"venc"`     // 视频编码器名称
	VencCtx TranscodeVencCtx `json:"venc_ctx"` // 视频AVCodecContext信息
}

type TranscodeSetting struct {
	AdecoderThreads int    `json:"adecoder_threads"` // ⾳频解码器线程数
	AencoderThreads int    `json:"aencoder_threads"` // ⾳频解码器线程数
	HwDecoder       bool   `json:"hw_decoder"`       // 启动硬件解码器
	HwEncoder       bool   `json:"hw_encoder"`       // 启动硬件编码器
	TargetAcodec    string `json:"target_acodec"`    // ⽬标⾳频编码格式
	TargetVcodec    string `json:"target_vcodec"`    // ⽬标视频编码格式
	VdecoderThreads int    `json:"vdecoder_threads"` // 视频解码器线程数
	VencoderThreads int    `json:"vencoder_threads"` // 视频编码器线程数
}

type TranscodeAencCtx struct {
	BitRate     int    `json:"bit_rate"`     // ⽐特率
	Channels    int    `json:"channels"`     // 通道数
	FrameNumber int    `json:"frame_number"` // 已编码帧数
	FrameSize   int    `json:"frame_size"`   // 每帧采样数
	SampleFmt   string `json:"sample_fmt"`   // ⾳频编码输⼊格式
	SampleRate  int    `json:"sample_rate"`  // 编码器采样率
}

type TranscodeVencCtx struct {
	BitRate     int    `json:"bit_rate"`     // ⽐特率
	Fps         int    `json:"fps"`          // 帧率
	FrameNumber int    `json:"frame_number"` // 已编码帧数
	Gop         int    `json:"gop"`          // gop⼤⼩
	HasBFrames  int    `json:"has_b_frames"` // 是否编码b帧
	Height      int    `json:"height"`       // 视频⾼度
	PixFmt      string `json:"pix_fmt"`      // 编码器输⼊图⽚格式
	Width       int    `json:"width"`        // 视频宽度
}
