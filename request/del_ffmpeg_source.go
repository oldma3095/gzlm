package request

type DelFFMpegSource struct {
	Base
	Key string `json:"key"` // addFFmpegSource接口返回的key
}
