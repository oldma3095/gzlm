package request

type CloseRtpServer struct {
	Base
	StreamId string `json:"stream_id"` // 该端口绑定的流id
}
