package hook

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
