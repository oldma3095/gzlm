package hook

type OnPublishInp struct {
	App           string `json:"app"`           // 流应用名
	Id            string `json:"id"`            // TCP链接唯一ID
	Ip            string `json:"ip"`            // 播放器ip
	Params        string `json:"params"`        // 播放url参数
	Port          uint32 `json:"port"`          // 播放器端口号
	Schema        string `json:"schema"`        // 播放的媒体源类型，可能是rtsp、rtmp、fmp4、ts、hls
	Protocol      string `json:"protocol"`      // 播放的传输协议，可能是rtsp/rtmp/rtsps/rtmps/rtc/srt/rtp/tcp/udp
	Stream        string `json:"stream"`        // 流ID
	Vhost         string `json:"vhost"`         // 流虚拟主机
	MediaServerId string `json:"mediaServerId"` // 服务器id,通过配置文件设置
}

type OnPublishOup struct {
	Code int    `json:"code"` // 错误代码，0代表允许推流
	Msg  string `json:"msg"`  // 不允许推流时的错误提示

	EnableHls      bool   `json:"enable_hls"`       //	非必须，是否转换成hls-mpegts协议
	EnableHlsFmp4  bool   `json:"enable_hls_fmp4"`  //	非必须，是否转换成hls-fmp4协议
	EnableMp4      bool   `json:"enable_mp4"`       //	非必须，是否允许mp4录制
	EnableRtsp     bool   `json:"enable_rtsp"`      //	非必须，是否转rtsp协议
	EnableRtmp     bool   `json:"enable_rtmp"`      //	非必须，是否转rtmp/flv协议
	EnableTs       bool   `json:"enable_ts"`        //	非必须，是否转http-ts/ws-ts协议
	EnableFmp4     bool   `json:"enable_fmp4"`      //	非必须，是否转http-fmp4/ws-fmp4协议
	HlsDemand      bool   `json:"hls_demand"`       //	非必须，该协议是否有人观看才生成
	RtspDemand     bool   `json:"rtsp_demand"`      //	非必须，该协议是否有人观看才生成
	RtmpDemand     bool   `json:"rtmp_demand"`      //	非必须，该协议是否有人观看才生成
	TsDemand       bool   `json:"ts_demand"`        //	非必须，该协议是否有人观看才生成
	Fmp4Demand     bool   `json:"fmp4_demand"`      //	非必须，该协议是否有人观看才生成
	EnableAudio    bool   `json:"enable_audio"`     //	非必须，转协议时是否开启音频
	AddMuteAudio   bool   `json:"add_mute_audio"`   //	非必须，转协议时，无音频是否添加静音aac音频
	Mp4SavePath    string `json:"mp4_save_path"`    //	非必须，mp4录制文件保存根目录，置空使用默认
	Mp4MaxSecond   int    `json:"mp4_max_second"`   //	非必须，mp4录制切片大小，单位秒
	Mp4AsPlayer    bool   `json:"mp4_as_player"`    //	非必须，MP4录制是否当作观看者参与播放人数计数
	HlsSavePath    string `json:"hls_save_path"`    //	非必须，hls文件保存保存根目录，置空使用默认
	ModifyStamp    int    `json:"modify_stamp"`     //	非必须，该流是否开启时间戳覆盖(0:绝对时间戳/1:系统时间戳/2:相对时间戳)
	ContinuePushMs uint32 `json:"continue_push_ms"` //	非必须，断连续推延时，单位毫秒，置空使用配置文件默认值
	AutoClose      bool   `json:"auto_close"`       //	非必须，无人观看是否自动关闭流(不触发无人观看hook)
	StreamReplace  string `json:"stream_replace"`   //	非必须，是否修改流id, 通过此参数可以自定义流id(譬如替换ssrc)

	AudioTranscode bool `json:"audio_transcode,omitempty"` // pro版本，自动opus->aac
}
