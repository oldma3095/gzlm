package response

type GetStatistic struct {
	Base
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
