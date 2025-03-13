package hook

import "gzlm/mediainfo"

// OnStreamChangedInp
// 流注册参数: mediaServerId app regist schema stream vhost
type OnStreamChangedInp struct {
	Regist        bool   `json:"regist"`        //	流注册或注销
	MediaServerId string `json:"mediaServerId"` //	服务器id,通过配置文件设置
	mediainfo.MediaInfo
}

type OnStreamChangedOup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
