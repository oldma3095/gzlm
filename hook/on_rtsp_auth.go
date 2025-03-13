package hook

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
