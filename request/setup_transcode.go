package request

type SetupTranscode struct {
	Base
	Vhost           string  `json:"vhost"`                      // 虚拟主机，例如__defaultVhost__
	APP             string  `json:"app"`                        // 应用名，例如 live
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
