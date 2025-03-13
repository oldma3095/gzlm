package gzlm

type Key struct {
	Key string `json:"key"`
}

type Flag struct {
	Flag bool `json:"flag"` // 成功与否
}

type AddFFMpegSourceRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Key    `json:"data"`
}

type AddStreamProxyRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Key    `json:"data"`
}

type AddStreamPusherProxyRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Key    `json:"data"`
}

type CloseRtpServerRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Hit  uint   `json:"hit"` // 是否找到记录并关闭
}

type CloseStreamsRes struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	CountHit    int    `json:"count_hit"`    // 筛选命中的流个数
	CountClosed int    `json:"count_closed"` // 被关闭的流个数，可能小于count_hit
}

type DelFFMpegSourceRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Flag   `json:"data"`
}

type DelStreamProxyRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Flag   `json:"data"`
}

type DelStreamPusherProxyRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Flag   `json:"data"`
}

type GetAllSessionRes struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data []Session `json:"data"`
}

type Session struct {
	Id        string `json:"id"`         // 该tcp链接唯一id
	LocalIp   string `json:"local_ip"`   // 本机网卡ip
	LocalPort int    `json:"local_port"` // 本机端口号(这是个rtmp播放器或推流器)
	PeerIp    string `json:"peer_ip"`    // 客户端ip
	PeerPort  int    `json:"peer_port"`  // 客户端端口号
	Typeid    string `json:"typeid"`     // 客户端TCPSession typeid
}

type GetApiListRes struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

type GetMediaListRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data []MediaInfo `json:"data"`
}

type GetMp4RecordFileRes struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data RecordInfo `json:"data"`
}

type RecordInfo struct {
	Paths    []string `json:"paths"`
	RootPath string   `json:"rootPath"`
}

// 搜索文件夹列表(按照前缀匹配规则)：period = 2020-01
//{
//	"code" : 0,
//	"data" : {
//	"paths" : [ "2020-01-25", "2020-01-24" ],
//	"rootPath" : "/www/record/live/ss/"
//	}
//}

// 搜索mp4文件列表：period = 2020-01-24
//{
//	"code" : 0,
//	"data" : {
//	"paths" : [
//		"22-20-30.mp4",
//		"22-13-12.mp4",
//		"21-57-07.mp4",
//		"21-19-18.mp4",
//		"21-24-21.mp4",
//		"21-15-10.mp4",
//		"22-14-14.mp4"
//	],
//	"rootPath" : "/www/live/ss/2020-01-24/"
//	}
//}

type GetProxyInfoRes struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data ProxyInfoData `json:"data"`
}

type ProxyInfoData struct {
	LiveSecs    int `json:"liveSecs"`
	RePullCount int `json:"rePullCount"`
	Status      int `json:"status"`
}

type GetProxyPusherInfoRes struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data ProxyPusherInfoData `json:"data"`
}

type ProxyPusherInfoData struct {
	LiveSecs       int `json:"liveSecs"`
	RePublishCount int `json:"rePublishCount"`
	Status         int `json:"status"`
}

type GetRtpInfoRes struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Exist     bool   `json:"exist"`      // 是否存在
	PeerIp    string `json:"peer_ip"`    // 推流客户端ip
	PeerPort  int    `json:"peer_port"`  // 客户端端口号
	LocalIp   string `json:"local_ip"`   // 本地监听的网卡ip
	LocalPort int    `json:"local_port"` // 本地监听的端口号
}

type GetServerConfigRes struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data []map[string]string `json:"data"`
}

type GetSnapRes struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Status bool   `json:"status"` // false:未录制,true:正在录制
}

type GetStatisticRes struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data StatisticInfo `json:"data"`
}

type StatisticInfo struct {
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

type GetThreadsLoadRes struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data []LoadInfo `json:"data"`
}

type LoadInfo struct {
	Delay int `json:"delay"` // 该线程延时
	Load  int `json:"load"`  // 该线程负载，0 ~ 100
}

type GetWorkThreadsLoadRes struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data []LoadInfo `json:"data"`
}

type IsMediaOnlineRes struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Online bool   `json:"online"` // 是否在线
}

type KickSessionsRes struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	CountHit int    `json:"count_hit"` // 筛选命中客户端个数
}

type ListRtpServerRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data []RtpServer `json:"data"`
}

type RtpServer struct {
	Port     int    `json:"port"`      // 绑定的端口号
	StreamId string `json:"stream_id"` // 绑定的流ID
}

type OpenRtpServerRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Port int    `json:"port"` // 接收端口，方便获取随机端口号
}

type RestartServerRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SetServerConfigRes struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Changed int    `json:"changed"` // 配置项变更个数
}

type SetupTranscodeRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type StartRecordRes struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result bool   `json:"result"` // 成功与否
}

type StartSendRtpRes struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	LocalPort int    `json:"local_port"` // 使用的本地端口号
}

type StopRecordRes struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result bool   `json:"result"` // 成功与否
}

type StopSendRtpRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
