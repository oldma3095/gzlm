package request

type GetMp4RecordFile struct {
	Base
	Vhost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	APP    string `json:"app"`    // 应用名，例如 live
	Stream string `json:"stream"` // 流id，例如 test

	CustomizedPath string `json:"customized_path,omitempty"` // 录像文件保存自定义根目录，为空则采用配置文件设置
	Period         string `json:"period,omitempty"`          // 流的录像日期，格式为2020-02-01,如果不是完整的日期，那么是搜索录像文件夹列表，否则搜索对应日期下的mp4文件列表
}
