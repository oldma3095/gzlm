package request

type GetRtpInfo struct {
	Base
	StreamId string `json:"stream_id"` // 流id
}
