package response

type GetSnap struct {
	Base
	Status bool `json:"status"` // false:未录制,true:正在录制
}
