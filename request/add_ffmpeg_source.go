package request

type AddFFMpegSource struct {
	Base
	SrcUrl    string `json:"src_url"`    // FFmpeg拉流地址,支持任意协议或格式(只要FFmpeg支持即可)
	DstUrl    string `json:"dst_url"`    // FFmpeg rtmp推流地址，一般都是推给自己，例如rtmp://127.0.0.1/live/stream_form_ffmpeg
	TimeoutMs uint   `json:"timeout_ms"` // FFmpeg推流成功超时时间,单位毫秒
	EnableHls uint   `json:"enable_hls"` // 是否开启hls录制
	EnableMp4 uint   `json:"enable_mp4"` // 是否开启mp4录制

	FfmpegCmdKey string `json:"ffmpeg_cmd_key,omitempty"` // FFmpeg命名参数模板，置空则采用配置项:ffmpeg.cmd
}
