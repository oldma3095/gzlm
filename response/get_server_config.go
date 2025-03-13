package response

type GetServerConfig struct {
	Base
	Data []map[string]string `json:"data"`
}
