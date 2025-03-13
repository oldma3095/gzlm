package hook

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
