package request

type StartRecord struct {
	Base
	Type   uint   `json:"type"`   // 0为hls，1为mp4
	Vhost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	APP    string `json:"app"`    // 应用名，例如 live
	Stream string `json:"stream"` // 流id，例如 obs

	CustomizedPath string `json:"customized_path,omitempty"` // 录像文件保存自定义根目录，为空则采用配置文件设置
	MaxSecond      uint   `json:"max_second,omitempty"`      // MP4录制的切片时间大小，单位秒
}
