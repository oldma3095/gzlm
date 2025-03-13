package response

import (
	"gzlm/mediainfo"
)

type GetMediaList struct {
	Base
	Data []mediainfo.MediaInfo `json:"data"`
}
