package request

type DelStreamPusherProxy struct {
	Base
	Key string `json:"key"` // addStreamPusherProxy接口返回的key
}
