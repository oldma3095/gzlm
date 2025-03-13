package response

type GetWorkThreadsLoad struct {
	Base
	Data []LoadInfo `json:"data"`
}
