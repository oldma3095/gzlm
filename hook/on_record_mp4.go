package hook

type OnRecordMp4Inp struct {
	App           string  `json:"app"`           //	录制的流应用名
	FileName      string  `json:"file_name"`     //	文件名
	FilePath      string  `json:"file_path"`     //	文件绝对路径
	FileSize      int     `json:"file_size"`     //	文件大小，单位字节
	Folder        string  `json:"folder"`        //	文件所在目录路径
	StartTime     int     `json:"start_time"`    //	开始录制时间戳
	Stream        string  `json:"stream"`        //	录制的流ID
	TimeLen       float64 `json:"time_len"`      //	录制时长，单位秒
	Url           string  `json:"url"`           //	http/rtsp/rtmp点播相对url路径
	Vhost         string  `json:"vhost"`         //	流虚拟主机
	MediaServerId string  `json:"mediaServerId"` //	服务器id,通过配置文件设置
}

type OnRecordMp4Oup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
