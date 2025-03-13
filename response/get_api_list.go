package response

type GetApiList struct {
	Base
	Data []string `json:"data"`
}
