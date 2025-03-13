package hook

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
