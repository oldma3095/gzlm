package response

type SetServerConfig struct {
	Base
	Changed int `json:"changed"` // 配置项变更个数
}
