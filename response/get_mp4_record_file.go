package response

type GetMp4RecordFile struct {
	Base
	Data RecordInfo `json:"data"`
}

type RecordInfo struct {
	Paths    []string `json:"paths"`
	RootPath string   `json:"rootPath"`
}

// 搜索文件夹列表(按照前缀匹配规则)：period = 2020-01
//{
//	"code" : 0,
//	"data" : {
//	"paths" : [ "2020-01-25", "2020-01-24" ],
//	"rootPath" : "/www/record/live/ss/"
//	}
//}

// 搜索mp4文件列表：period = 2020-01-24
//{
//	"code" : 0,
//	"data" : {
//	"paths" : [
//		"22-20-30.mp4",
//		"22-13-12.mp4",
//		"21-57-07.mp4",
//		"21-19-18.mp4",
//		"21-24-21.mp4",
//		"21-15-10.mp4",
//		"22-14-14.mp4"
//	],
//	"rootPath" : "/www/live/ss/2020-01-24/"
//	}
//}
