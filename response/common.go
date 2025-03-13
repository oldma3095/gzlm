package response

type Base struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Key struct {
	Key string `json:"key"`
}

type Flag struct {
	Flag bool `json:"flag"` // 成功与否
}
