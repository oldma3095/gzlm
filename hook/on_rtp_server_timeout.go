package hook

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
