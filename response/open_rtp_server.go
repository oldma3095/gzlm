package response

type OpenRtpServer struct {
	Base
	Port int `json:"port"` // 接收端口，方便获取随机端口号
}
