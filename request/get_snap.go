package request

type GetSnap struct {
	Base
	URL        string `json:"url"`         // 需要截图的url，可以是本机的，也可以是远程主机的
	TimeoutSec uint   `json:"timeout_Sec"` // 截图失败超时时间，防止FFmpeg一直等待截图
	ExpireSec  uint   `json:"expire_Sec"`  // 截图的过期时间，该时间内产生的截图都会作为缓存返回
}
