package gzlm

type AddFFMpegSourceReq struct {
	Secret    string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	SrcUrl    string `json:"src_url"`          // FFmpeg拉流地址,支持任意协议或格式(只要FFmpeg支持即可)
	DstUrl    string `json:"dst_url"`          // FFmpeg rtmp推流地址，一般都是推给自己，例如rtmp://127.0.0.1/live/stream_form_ffmpeg
	TimeoutMs uint   `json:"timeout_ms"`       // FFmpeg推流成功超时时间,单位毫秒
	EnableHls uint   `json:"enable_hls"`       // 是否开启hls录制
	EnableMp4 uint   `json:"enable_mp4"`       // 是否开启mp4录制

	FfmpegCmdKey string `json:"ffmpeg_cmd_key,omitempty"` // FFmpeg命名参数模板，置空则采用配置项:ffmpeg.cmd
}

type AddStreamProxyReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Vhost  string `json:"vhost"`            // 添加的流的虚拟主机，例如__defaultVhost__
	App    string `json:"app"`              // 添加的流的应用名，例如live
	Stream string `json:"stream"`           // 添加的流的id名，例如test
	Url    string `json:"url"`              // 拉流地址，例如rtmp://live.hkstv.hk.lxdns.com/live/hks2

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

type AddStreamPusherProxyReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Schema string `json:"schema"`           // 推流协议，支持rtsp、rtmp，大小写敏感
	Vhost  string `json:"vhost"`            // 已注册流的虚拟主机，一般为__defaultVhost__
	App    string `json:"app"`              // 已注册流的应用名，例如live
	Stream string `json:"stream"`           // 已注册流的id名，例如test
	DstUrl string `json:"dst_url"`          // 推流地址，需要与schema字段协议一致

	RtpType    uint    `json:"rtp_type,omitempty"`    // rtsp推流时，推流方式，0：tcp，1：udp
	TimeoutSec float32 `json:"timeout_sec,omitempty"` // 推流超时时间，单位秒，float类型
	RetryCount int     `json:"retry_count,omitempty"` // 推流重试次数,不传此参数或传值<=0时，则无限重试
}

type CloseRtpServerReq struct {
	Secret   string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	StreamId string `json:"stream_id"`        // 该端口绑定的流id
}

type CloseStreamsReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Schema string `json:"schema,omitempty"` // 协议，例如 rtsp或rtmp
	Vhost  string `json:"vhost,omitempty"`  // 虚拟主机，例如__defaultVhost__
	App    string `json:"app,omitempty"`    // 应用名，例如 live
	Stream string `json:"stream,omitempty"` // 流id，例如 test
	Force  uint   `json:"force"`            // 是否强制关闭(有人在观看是否还关闭)
}

type DelFFMpegSourceReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Key    string `json:"key"`              // addFFmpegSource接口返回的key
}

type DelStreamProxyReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Key    string `json:"key"`              // addStreamProxy接口返回的key
}

type DelStreamPusherProxyReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Key    string `json:"key"`              // addStreamPusherProxy接口返回的key
}

type GetAllSessionReq struct {
	Secret    string `json:"secret,omitempty"`     // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	LocalPort string `json:"local_port,omitempty"` // 筛选本机端口，例如筛选rtsp链接：554
	PeerIP    string `json:"peer_ip,omitempty"`    // 筛选客户端ip
}

type GetApiListReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
}

type GetMediaListReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Schema string `json:"schema,omitempty"` // 筛选协议，例如 rtsp或rtmp
	Vhost  string `json:"vhost,omitempty"`  // 筛选虚拟主机，例如__defaultVhost__
	App    string `json:"app,omitempty"`    // 筛选应用名，例如 live
	Stream string `json:"stream,omitempty"` // 筛选流id，例如 test
}

type GetMp4RecordFileReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Vhost  string `json:"vhost"`            // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`              // 应用名，例如 live
	Stream string `json:"stream"`           // 流id，例如 test

	CustomizedPath string `json:"customized_path,omitempty"` // 录像文件保存自定义根目录，为空则采用配置文件设置
	Period         string `json:"period,omitempty"`          // 流的录像日期，格式为2020-02-01,如果不是完整的日期，那么是搜索录像文件夹列表，否则搜索对应日期下的mp4文件列表
}

type GetProxyInfoReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Key    string `json:"key"`              // 拉流唯一标识
}

type GetProxyPusherInfoReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Key    string `json:"key"`              // 推流流唯一标识
}

type GetRtpInfoReq struct {
	Secret   string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	StreamId string `json:"stream_id"`        // 流id
}

type GetServerConfigReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
}

type GetSnapReq struct {
	Secret     string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Url        string `json:"url"`              // 需要截图的url，可以是本机的，也可以是远程主机的
	TimeoutSec uint   `json:"timeout_Sec"`      // 截图失败超时时间，防止FFmpeg一直等待截图
	ExpireSec  uint   `json:"expire_Sec"`       // 截图的过期时间，该时间内产生的截图都会作为缓存返回
}

type GetStatisticReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
}

type GetThreadsLoadReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
}

type GetWorkThreadsLoadReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
}

type IsMediaOnlineReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Schema string `json:"schema"`           // 协议，例如 rtsp或rtmp
	Vhost  string `json:"vhost"`            // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`              // 应用名，例如 live
	Stream string `json:"stream"`           // 流id，例如 test
}

type KickSessionsReq struct {
	Secret    string `json:"secret,omitempty"`     // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	LocalPort string `json:"local_port,omitempty"` // 筛选本机端口，例如筛选rtsp链接：554
	PeerIP    string `json:"peer_ip,omitempty"`    // 筛选客户端ip
}

type ListRtpServerReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
}

type OpenRtpServerReq struct {
	Secret   string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Port     uint   `json:"port"`             // 绑定的端口，0时为随机端口
	TcpMode  uint   `json:"tcp_mode"`         // tcp模式，0时为不启用tcp监听，1时为启用tcp监听，2时为tcp主动连接模式
	StreamId string `json:"stream_id"`        // 该端口绑定的流id

	ReUsePort uint `json:"re_use_port,omitempty"` // 是否重用端口，默认为0，非必选参数
	Ssrc      uint `json:"ssrc,omitempty"`        // 是否指定收流的rtp ssrc, 十进制数字，不指定或指定0时则不过滤rtp，非必选参数
	OnlyAudio uint `json:"only_audio,omitempty"`  // 是否为单音频track，用于语音对讲
}

type RestartServerReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
}

// SetupTranscode pro版本
type SetupTranscodeReq struct {
	Secret          string  `json:"secret,omitempty"`           // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Vhost           string  `json:"vhost"`                      // 虚拟主机，例如__defaultVhost__
	App             string  `json:"app"`                        // 应用名，例如 live
	Stream          string  `json:"stream"`                     // 流id，例如 test
	Name            string  `json:"name"`                       // 转码名(后缀)，功能类似配置⽂件transcode_suffix
	Add             int     `json:"add"`                        // 1：添加转码; 0: 删除转码
	VideoCodec      string  `json:"video_codec,omitempty"`      // 视频转码的codec,⽀持H264/H265/JPEG
	VideoBitrate    int     `json:"video_bitrate,omitempty"`    // 转码后视频的⽐特率
	VideoScale      float32 `json:"video_scale,omitempty"`      // 转码视频宽⾼拉伸⽐例，取值范围0.1~10，原始比例1080，720:0.6666667，640:0.5925926，360:0.33333334
	AudioCodec      string  `json:"audio_codec,omitempty"`      // ⾳频转码codec，⽀持mpeg4-generic，PCMA，PCMU，opus
	AudioBitrate    int     `json:"audio_bitrate,omitempty"`    // 转码后⾳频⽐特率
	AudioSamplerate int     `json:"audio_samplerate,omitempty"` // 转码后⾳频采样率率

	Force         *uint  `json:"force,omitempty"`           // 是否强制转码 0不启用 1启用 不传默认
	Filter        string `json:"filter"`                    // avfilter滤镜参数,⽤法与ffmpeg -vf 参数⼀致 scale=640x360 或者水印
	EnableHls     *uint  `json:"enable_hls,omitempty"`      // 是否转换成hls协议 0不启用 1启用 不传默认
	EnableHlsFmp4 *uint  `json:"enable_hls_fmp4,omitempty"` // 转码后是否转换成hls-fmp4协议 0不启用 1启用 不传默认
	EnableMp4     *uint  `json:"enable_mp4,omitempty"`      // 是否允许mp4录制 0不启用 1启用 不传默认
	EnableRtsp    *uint  `json:"enable_rtsp,omitempty"`     // 是否转rtsp协议 0不启用 1启用 不传默认
	EnableRtmp    *uint  `json:"enable_rtmp,omitempty"`     // 是否转rtmp/flv协议 0不启用 1启用 不传默认
	EnableTs      *uint  `json:"enable_ts,omitempty"`       // 是否转http-ts/ws-ts协议 0不启用 1启用 不传默认
	EnableFmp4    *uint  `json:"enable_fmp4,omitempty"`     // 是否转http-fmp4/ws-fmp4协议 0不启用 1启用 不传默认
	HlsDemand     *uint  `json:"hls_demand,omitempty"`      // 转码后该协议是否有⼈观看才⽣成 0不启用 1启用 不传默认
	RtspDemand    *uint  `json:"rtsp_demand,omitempty"`     // 转码后该协议是否有⼈观看才⽣成 0不启用 1启用 不传默认
	RtmpDemand    *uint  `json:"rtmp_demand,omitempty"`     // 转码后该协议是否有⼈观看才⽣成 0不启用 1启用 不传默认
	TsDemand      *uint  `json:"ts_demand,omitempty"`       // 转码后该协议是否有⼈观看才⽣成 0不启用 1启用 不传默认
	Fmp4Demand    *uint  `json:"fmp4_demand,omitempty"`     // 转码后该协议是否有⼈观看才⽣成 0不启用 1启用 不传默认
	EnableAudio   *uint  `json:"enable_audio,omitempty"`    // 转协议时是否开启音频 0不启用 1启用 不传默认
	AddMuteAudio  *uint  `json:"add_mute_audio,omitempty"`  // 转协议时，无音频是否添加静音aac音频 0不启用 1启用 不传默认
	Mp4SavePath   string `json:"mp4_save_path,omitempty"`   // mp4录制文件保存根目录，置空使用默认
	Mp4MaxSecond  int    `json:"mp4_max_second,omitempty"`  // mp4录制切片大小，单位秒
	Mp4AsPlayer   *uint  `json:"mp4_as_player,omitempty"`   // MP4录制是否当作观看者参与播放人数计数 0不启用 1启用 不传默认
	HlsSavePath   string `json:"hls_save_path,omitempty"`   // hls文件保存保存根目录，置空使用默认
	ModifyStamp   *uint  `json:"modify_stamp,omitempty"`    // 该流是否开启时间戳覆盖 0不启用 1启用 不传默认
	AutoClose     *uint  `json:"auto_close,omitempty"`      // 无人观看是否自动关闭流(不触发无人观看hook) 0不启用 1启用 不传默认

	HwDecoder      *uint    `json:"hw_decoder,omitempty"`      // 是否开启硬件解码 0不启用 1启用 不传默认
	HwEncoder      *uint    `json:"hw_encoder,omitempty"`      // 是否开启硬件编码 0不启用 1启用 不传默认
	DecoderThreads uint     `json:"decoder_threads,omitempty"` // 解码器线程数
	EncoderThreads uint     `json:"encoder_threads,omitempty"` // 编码器线程数
	DecoderList    []string `json:"decoder_list,omitempty"`    // 解码器名称数组，优先级以前面为大
	EncoderList    []string `json:"encoder_list,omitempty"`    // 编码器名称数组，优先级以前面为大
}

type StartRecordReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Type   uint   `json:"type"`             // 0为hls，1为mp4
	Vhost  string `json:"vhost"`            // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`              // 应用名，例如 live
	Stream string `json:"stream"`           // 流id，例如 obs

	CustomizedPath string `json:"customized_path,omitempty"` // 录像文件保存自定义根目录，为空则采用配置文件设置
	MaxSecond      uint   `json:"max_second,omitempty"`      // MP4录制的切片时间大小，单位秒
}

type StartSendRtpReq struct {
	Secret  string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Vhost   string `json:"vhost"`            // 虚拟主机，例如__defaultVhost__
	App     string `json:"app"`              // 应用名，例如 live
	Stream  string `json:"stream"`           // 流id，例如 obs
	Ssrc    uint   `json:"ssrc"`             // rtp推流的ssrc，ssrc不同时，可以推流到多个上级服务器
	DstUrl  string `json:"dst_url"`          // rtp打包采用ps还是es模式，默认采用ps模式，该参数非必选参数
	DstPort uint   `json:"dst_port"`         // 目标端口
	IsUdp   uint   `json:"is_udp"`           // 是否为udp模式,否则为tcp模式

	SrcPort        uint   `json:"src_port,omitempty"`         // 指定tcp/udp客户端使用的本地端口，0时为随机端口，该参数非必选参数，不传时为随机端口。
	FromMp4        uint   `json:"from_mp4,omitempty"`         // 是否推送本地MP4录像，该参数非必选参数
	UsePs          uint   `json:"use_ps,omitempty"`           // rtp打包采用ps还是es模式，默认采用ps模式，该参数非必选参数
	PT             uint   `json:"pt,omitempty"`               // rtp payload type，默认96，该参数非必选参数
	OnlyAudio      uint   `json:"only_audio,omitempty"`       // rtp es方式打包时，是否只打包音频，该参数非必选参数
	UdpRtcpTimeout uint   `json:"udp_rtcp_timeout,omitempty"` // udp方式推流时，是否开启rtcp发送和rtcp接收超时判断，开启后(默认关闭)，如果接收rr rtcp超时，将导致主动停止rtp发送
	RecvStreamId   string `json:"recv_stream_id,omitempty"`   // 发送rtp同时接收，一般用于双向语言对讲, 如果不为空，说明开启接收，值为接收流的id
}

type StopRecordReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Type   uint   `json:"type"`             // 0为hls，1为mp4
	Vhost  string `json:"vhost"`            // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`              // 应用名，例如 live
	Stream string `json:"stream"`           // 流id，例如 obs
}

type StopSendRtpReq struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
	Vhost  string `json:"vhost"`            // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`              // 应用名，例如 live
	Stream string `json:"stream"`           // 流id，例如 obs

	Ssrc uint `json:"ssrc,omitempty"` // 根据ssrc关停某路rtp推流，不传时关闭所有推流
}
