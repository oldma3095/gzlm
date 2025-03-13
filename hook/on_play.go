package hook

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
