package hook

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
