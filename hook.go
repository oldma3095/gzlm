package gzlm

type OnFlowReportInp struct {
	App           string `json:"app"`           //	流应用名
	Duration      int    `json:"duration"`      //	tcp链接维持时间，单位秒
	Params        string `json:"params"`        //	推流或播放url参数
	Player        bool   `json:"player"`        //	true为播放器，false为推流器
	Schema        string `json:"schema"`        //	播放或推流的媒体源类型，可能是rtsp、rtmp、fmp4、ts、hls
	Protocol      string `json:"protocol"`      //	传输的协议，可能是/http/https/ws/wss/rtsp/rtmp/rtsps/rtmps/rtc/srt/rtp/tcp/udp
	Stream        string `json:"stream"`        //	流ID
	TotalBytes    int    `json:"totalBytes"`    //	耗费上下行流量总和，单位字节
	Vhost         string `json:"vhost"`         //	流虚拟主机
	Ip            string `json:"ip"`            //	客户端ip
	Port          int    `json:"port"`          //	客户端端口号
	Id            string `json:"id"`            //	TCP链接唯一ID
	MediaServerId string `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnFlowReportOup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OnHttpAccessInp struct {
	HeaderAccept                  string `json:"header.Accept"`
	HeaderAcceptEncoding          string `json:"header.Accept-Encoding"`
	HeaderAcceptLanguage          string `json:"header.Accept-Language"`
	HeaderCacheControl            string `json:"header.Cache-Control"`
	HeaderConnection              string `json:"header.Connection"`
	HeaderHost                    string `json:"header.Host"`
	HeaderUpgradeInsecureRequests string `json:"header.Upgrade-Insecure-Requests"`
	HeaderUserAgent               string `json:"header.User-Agent"`

	Id            string `json:"id"`            //	TCP链接唯一ID
	Ip            string `json:"ip"`            //	http客户端ip
	IsDir         bool   `json:"is_dir"`        //	http 访问路径是文件还是目录
	Params        string `json:"params"`        //	http url参数
	Path          string `json:"path"`          //	请求访问的文件或目录
	Port          uint32 `json:"port"`          // 	http客户端端口号
	MediaServerId string `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnHttpAccessOup struct {
	Code          int    `json:"code"`          //	请固定返回0
	Err           string `json:"err"`           //	不允许访问的错误提示，允许访问请置空
	Path          string `json:"path"`          //	该客户端能访问或被禁止的顶端目录，如果为空字符串，则表述为当前目录
	Second        int    `json:"second"`        //	本次授权结果的有效期，单位秒
	MediaServerId string `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnPlayInp struct {
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

type OnPlayOup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

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

	EnableHls      *bool   `json:"enable_hls,omitempty"`       //	非必须，是否转换成hls-mpegts协议
	EnableHlsFmp4  *bool   `json:"enable_hls_fmp4,omitempty"`  //	非必须，是否转换成hls-fmp4协议
	EnableMp4      *bool   `json:"enable_mp4,omitempty"`       //	非必须，是否允许mp4录制
	EnableRtsp     *bool   `json:"enable_rtsp,omitempty"`      //	非必须，是否转rtsp协议
	EnableRtmp     *bool   `json:"enable_rtmp,omitempty"`      //	非必须，是否转rtmp/flv协议
	EnableTs       *bool   `json:"enable_ts,omitempty"`        //	非必须，是否转http-ts/ws-ts协议
	EnableFmp4     *bool   `json:"enable_fmp4,omitempty"`      //	非必须，是否转http-fmp4/ws-fmp4协议
	HlsDemand      *bool   `json:"hls_demand,omitempty"`       //	非必须，该协议是否有人观看才生成
	RtspDemand     *bool   `json:"rtsp_demand,omitempty"`      //	非必须，该协议是否有人观看才生成
	RtmpDemand     *bool   `json:"rtmp_demand,omitempty"`      //	非必须，该协议是否有人观看才生成
	TsDemand       *bool   `json:"ts_demand,omitempty"`        //	非必须，该协议是否有人观看才生成
	Fmp4Demand     *bool   `json:"fmp4_demand,omitempty"`      //	非必须，该协议是否有人观看才生成
	EnableAudio    *bool   `json:"enable_audio,omitempty"`     //	非必须，转协议时是否开启音频
	AddMuteAudio   *bool   `json:"add_mute_audio,omitempty"`   //	非必须，转协议时，无音频是否添加静音aac音频
	Mp4SavePath    *string `json:"mp4_save_path,omitempty"`    //	非必须，mp4录制文件保存根目录，置空使用默认
	Mp4MaxSecond   *int    `json:"mp4_max_second,omitempty"`   //	非必须，mp4录制切片大小，单位秒
	Mp4AsPlayer    *bool   `json:"mp4_as_player,omitempty"`    //	非必须，MP4录制是否当作观看者参与播放人数计数
	HlsSavePath    *string `json:"hls_save_path,omitempty"`    //	非必须，hls文件保存保存根目录，置空使用默认
	ModifyStamp    *int    `json:"modify_stamp,omitempty"`     //	非必须，该流是否开启时间戳覆盖(0:绝对时间戳/1:系统时间戳/2:相对时间戳)
	ContinuePushMs *uint32 `json:"continue_push_ms,omitempty"` //	非必须，断连续推延时，单位毫秒，置空使用配置文件默认值
	AutoClose      *bool   `json:"auto_close,omitempty"`       //	非必须，无人观看是否自动关闭流(不触发无人观看hook)
	StreamReplace  *string `json:"stream_replace,omitempty"`   //	非必须，是否修改流id, 通过此参数可以自定义流id(譬如替换ssrc)

	AudioTranscode *bool `json:"audio_transcode,omitempty"` // pro版本，自动opus->aac
}

type OnRecordMp4Inp struct {
	App           string  `json:"app"`           //	录制的流应用名
	FileName      string  `json:"file_name"`     //	文件名
	FilePath      string  `json:"file_path"`     //	文件绝对路径
	FileSize      int     `json:"file_size"`     //	文件大小，单位字节
	Folder        string  `json:"folder"`        //	文件所在目录路径
	StartTime     int     `json:"start_time"`    //	开始录制时间戳
	Stream        string  `json:"stream"`        //	录制的流ID
	TimeLen       float64 `json:"time_len"`      //	录制时长，单位秒
	Url           string  `json:"url"`           //	http/rtsp/rtmp点播相对url路径
	Vhost         string  `json:"vhost"`         //	流虚拟主机
	MediaServerId string  `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnRecordMp4Oup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OnRtpServerTimeoutInp struct {
	LocalPort     int    `json:"local_port"`    //	openRtpServer 输入的参数
	ReUsePort     bool   `json:"re_use_port"`   //	openRtpServer 输入的参数
	Ssrc          uint32 `json:"ssrc"`          //	openRtpServer 输入的参数
	StreamId      string `json:"stream_id"`     //	openRtpServer 输入的参数
	TcpMode       int    `json:"tcp_mode"`      //	openRtpServer 输入的参数
	MediaServerId string `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnRtpServerTimeoutOup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OnRtspAuthInp struct {
	App           string `json:"app"`             //	流应用名
	Id            string `json:"id"`              //	TCP链接唯一ID
	Ip            string `json:"ip"`              //	rtsp播放器ip
	MustNoEncrypt bool   `json:"must_no_encrypt"` //	请求的密码是否必须为明文(base64鉴权需要明文密码)
	Params        string `json:"params"`          //	rtsp url参数
	Port          uint32 `json:"port"`            // 	rtsp播放器端口号
	Realm         string `json:"realm"`           //	rtsp播放鉴权加密realm
	Schema        string `json:"schema"`          //	rtsp
	Protocol      string `json:"protocol"`        //	rtsp或rtsps
	Stream        string `json:"stream"`          //	流ID
	UserName      string `json:"user_name"`       //	播放用户名
	Vhost         string `json:"vhost"`           //	流虚拟主机
	MediaServerId string `json:"mediaServerId"`   //	服务器id,通过配置文件设置
}

type OnRtspAuthOup struct {
	Code      int    `json:"code"`      //	错误代码，0代表允许播放
	Msg       string `json:"msg"`       //	播放鉴权失败时的错误提示
	Encrypted bool   `json:"encrypted"` //	用户密码是明文还是摘要
	Passwd    string `json:"passwd"`    //	用户密码明文或摘要(md5(username:realm:password))
}

type OnRtspRealmInp struct {
	App           string `json:"app"`           //	流应用名
	Id            string `json:"id"`            //	TCP链接唯一ID
	Ip            string `json:"ip"`            //	rtsp播放器ip
	Params        string `json:"params"`        //	播放rtsp url参数
	Port          uint32 `json:"port"`          // 	rtsp播放器端口号
	Schema        string `json:"schema"`        //	rtsp
	Protocol      string `json:"protocol"`      //	传输协议
	Stream        string `json:"stream"`        //	流ID
	Vhost         string `json:"vhost"`         //	流虚拟主机
	MediaServerId string `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnRtspRealmOup struct {
	Code  int    `json:"code"`  // 请固定返回0
	Realm string `json:"realm"` // 该rtsp流是否需要rtsp专有鉴权，空字符串代码不需要鉴权
}

type OnServerKeepaliveData struct {
	Buffer                int `json:"Buffer"`
	BufferLikeString      int `json:"BufferLikeString"`
	BufferList            int `json:"BufferList"`
	BufferRaw             int `json:"BufferRaw"`
	Frame                 int `json:"Frame"`
	FrameImp              int `json:"FrameImp"`
	MediaSource           int `json:"MediaSource"`
	MultiMediaSourceMuxer int `json:"MultiMediaSourceMuxer"`
	RtmpPacket            int `json:"RtmpPacket"`
	RtpPacket             int `json:"RtpPacket"`
	Socket                int `json:"Socket"`
	TcpClient             int `json:"TcpClient"`
	TcpServer             int `json:"TcpServer"`
	TcpSession            int `json:"TcpSession"`
	UdpServer             int `json:"UdpServer"`
	UdpSession            int `json:"UdpSession"`
}
type OnServerKeepaliveInp struct {
	Data          OnServerKeepaliveData `json:"data"`
	MediaServerId string                `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnServerKeepaliveOup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OnServerStartedInp struct {
	ApiApiDebug                    string `json:"api.apiDebug"`
	ApiSecret                      string `json:"api.secret"`
	FfmpegBin                      string `json:"ffmpeg.bin"`
	FfmpegCmd                      string `json:"ffmpeg.cmd"`
	FfmpegLog                      string `json:"ffmpeg.log"`
	GeneralMediaServerId           string `json:"general.mediaServerId"`
	GeneralAddMuteAudio            string `json:"general.addMuteAudio"`
	GeneralEnableVhost             string `json:"general.enableVhost"`
	GeneralFlowThreshold           string `json:"general.flowThreshold"`
	GeneralMaxStreamWaitMS         string `json:"general.maxStreamWaitMS"`
	GeneralPublishToHls            string `json:"general.publishToHls"`
	GeneralPublishToMP4            string `json:"general.publishToMP4"`
	GeneralPublishToRtxp           string `json:"general.publishToRtxp"`
	GeneralResetWhenRePlay         string `json:"general.resetWhenRePlay"`
	GeneralStreamNoneReaderDelayMS string `json:"general.streamNoneReaderDelayMS"`
	GeneralUltraLowDelay           string `json:"general.ultraLowDelay"`
	HlsFileBufSize                 string `json:"hls.fileBufSize"`
	HlsFilePath                    string `json:"hls.filePath"`
	HlsSegDur                      string `json:"hls.segDur"`
	HlsSegNum                      string `json:"hls.segNum"`
	HlsSegRetain                   string `json:"hls.segRetain"`
	HookAdminParams                string `json:"hook.admin_params"`
	HookEnable                     string `json:"hook.enable"`
	HookOnFlowReport               string `json:"hook.on_flow_report"`
	HookOnHttpAccess               string `json:"hook.on_http_access"`
	HookOnPlay                     string `json:"hook.on_play"`
	HookOnPublish                  string `json:"hook.on_publish"`
	HookOnRecordMp4                string `json:"hook.on_record_mp4"`
	HookOnRtspAuth                 string `json:"hook.on_rtsp_auth"`
	HookOnRtspRealm                string `json:"hook.on_rtsp_realm"`
	HookOnServerStarted            string `json:"hook.on_server_started"`
	HookOnShellLogin               string `json:"hook.on_shell_login"`
	HookOnStreamChanged            string `json:"hook.on_stream_changed"`
	HookOnStreamNoneReader         string `json:"hook.on_stream_none_reader"`
	HookOnStreamNotFound           string `json:"hook.on_stream_not_found"`
	HookTimeoutSec                 string `json:"hook.timeoutSec"`
	HttpCharSet                    string `json:"http.charSet"`
	HttpKeepAliveSecond            string `json:"http.keepAliveSecond"`
	HttpMaxReqCount                string `json:"http.maxReqCount"`
	HttpMaxReqSize                 string `json:"http.maxReqSize"`
	HttpNotFound                   string `json:"http.notFound"`
	HttpPort                       string `json:"http.port"`
	HttpRootPath                   string `json:"http.rootPath"`
	HttpSendBufSize                string `json:"http.sendBufSize"`
	HttpSslPort                    string `json:"http.sslport"`
	MulticastAddrMax               string `json:"multicast.addrMax"`
	MulticastAddrMin               string `json:"multicast.addrMin"`
	MulticastUdpTTL                string `json:"multicast.udpTTL"`
	RecordAppName                  string `json:"record.appName"`
	RecordFastStart                string `json:"record.fastStart"`
	RecordFileBufSize              string `json:"record.fileBufSize"`
	RecordFilePath                 string `json:"record.filePath"`
	RecordFileRepeat               string `json:"record.fileRepeat"`
	RecordFileSecond               string `json:"record.fileSecond"`
	RecordSampleMS                 string `json:"record.sampleMS"`
	RtmpHandshakeSecond            string `json:"rtmp.handshakeSecond"`
	RtmpKeepAliveSecond            string `json:"rtmp.keepAliveSecond"`
	RtmpModifyStamp                string `json:"rtmp.modifyStamp"`
	RtmpPort                       string `json:"rtmp.port"`
	RtpAudioMtuSize                string `json:"rtp.audioMtuSize"`
	RtpClearCount                  string `json:"rtp.clearCount"`
	RtpCycleMS                     string `json:"rtp.cycleMS"`
	RtpMaxRtpCount                 string `json:"rtp.maxRtpCount"`
	RtpVideoMtuSize                string `json:"rtp.videoMtuSize"`
	RtspAuthBasic                  string `json:"rtsp.authBasic"`
	RtspDirectProxy                string `json:"rtsp.directProxy"`
	RtspHandshakeSecond            string `json:"rtsp.handshakeSecond"`
	RtspKeepAliveSecond            string `json:"rtsp.keepAliveSecond"`
	RtspModifyStamp                string `json:"rtsp.modifyStamp"`
	RtspPort                       string `json:"rtsp.port"`
	RtspSslPort                    string `json:"rtsp.sslport"`
	ShellMaxReqSize                string `json:"shell.maxReqSize"`
	ShellPort                      string `json:"shell.port"`
}

type OnServerStartedOup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OnShellLoginInp struct {
	Id            string `json:"id"`            //	TCP链接唯一ID
	Ip            string `json:"ip"`            //	telnet 终端ip
	Passwd        bool   `json:"passwd"`        //	telnet 终端登录用户密码
	Port          uint32 `json:"port"`          //  telnet 终端端口号
	UserName      string `json:"user_name"`     //	telnet 终端登录用户名
	MediaServerId string `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnShellLoginOup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// OnStreamChangedInp
// 流注册参数: mediaServerId app regist schema stream vhost
type OnStreamChangedInp struct {
	Regist        bool   `json:"regist"`        //	流注册或注销
	MediaServerId string `json:"mediaServerId"` //	服务器id,通过配置文件设置
	MediaInfo
}

type OnStreamChangedOup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OnStreamNoneReaderInp struct {
	App           string `json:"app"`           // 流应用名
	Schema        string `json:"schema"`        // rtsp或rtmp
	Stream        string `json:"stream"`        // 流ID
	Vhost         string `json:"vhost"`         // 流虚拟主机
	MediaServerId string `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnStreamNoneReaderOup struct {
	Code  int  `json:"code"`  //	请固定返回0
	Close bool `json:"close"` //	是否关闭推流或拉流
}

type OnStreamNotFoundInp struct {
	App           string `json:"app"`           //	流应用名
	Id            string `json:"id"`            //	TCP链接唯一ID
	Ip            string `json:"ip"`            //	播放器ip
	Params        string `json:"params"`        //	播放url参数
	Port          uint32 `json:"port"`          // 播放器端口号
	Schema        string `json:"schema"`        //	播放媒体源类型，可能是rtsp、rtmp
	Protocol      string `json:"protocol"`      //	传输协议
	Stream        string `json:"stream"`        //	流ID
	Vhost         string `json:"vhost"`         //	流虚拟主机
	MediaServerId string `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnStreamNotFoundOup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
