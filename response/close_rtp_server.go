package response

type CloseRtpServer struct {
	Base
	Hit uint `json:"hit"` // 是否找到记录并关闭
}
