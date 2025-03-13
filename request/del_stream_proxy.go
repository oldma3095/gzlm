package request

type DelStreamProxy struct {
	Base
	Key string `json:"key"` // addStreamProxy接口返回的key
}
