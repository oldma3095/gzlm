package hook

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
