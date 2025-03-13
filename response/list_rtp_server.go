package response

type ListRtpServer struct {
	Base
	Data []RtpServer `json:"data"`
}

type RtpServer struct {
	Port     int    `json:"port"`      // 绑定的端口号
	StreamId string `json:"stream_id"` // 绑定的流ID
}
