package response

type GetThreadsLoad struct {
	Base
	Data []LoadInfo `json:"data"`
}

type LoadInfo struct {
	Delay int `json:"delay"` // 该线程延时
	Load  int `json:"load"`  // 该线程负载，0 ~ 100
}
