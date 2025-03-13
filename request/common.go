package request

type Base struct {
	Secret string `json:"secret,omitempty"` // api操作密钥(配置文件配置)，如果操作ip是127.0.0.1，则不需要此参数
}
