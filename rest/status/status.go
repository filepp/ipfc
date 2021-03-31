package status

const (
	// 成功
	StatusOK = 0

	// 一般错误
	StatusInvalidParam = 10001
	StatusNotFound     = 10002
	StatusNotImpl      = 10003

	// 严重错误
	StatusDbError            = 50001
	StatusFileOperationError = 50002
	StatusNetworkError       = 50003
)

type Language int

const (
	En = Language(0)
	Cn = Language(1)
)

var statusText map[Language]map[int]string

func init() {
	statusText = make(map[Language]map[int]string)
	statusText[En] = make(map[int]string)
	statusText[Cn] = make(map[int]string)
	initEnCode(statusText[En])
	initCnCode(statusText[Cn])
}

func initEnCode(statusText map[int]string) {
	statusText[StatusOK] = "success"
	statusText[StatusInvalidParam] = "invalid param"
	statusText[StatusNotFound] = "resource not found"
	statusText[StatusNotImpl] = "function not implemented"

	statusText[StatusDbError] = "db error"
	statusText[StatusFileOperationError] = "file operation error"
	statusText[StatusNetworkError] = "network error"
}

func initCnCode(statusText map[int]string) {
	statusText[StatusOK] = "成功"
	statusText[StatusInvalidParam] = "参数不合法，请检查参数"
	statusText[StatusNotFound] = "请求的资源不存在"
	statusText[StatusNotImpl] = "功能未实现"
	statusText[StatusDbError] = "数据库错误"
	statusText[StatusFileOperationError] = "文件操作错误"
	statusText[StatusNetworkError] = "网络错误"
}

type Code int

func (c Code) String() string {
	msg, _ := statusText[En][int(c)]
	return msg
}

func (c Code) Text(lang Language) string {
	msg, _ := statusText[lang][int(c)]
	return msg
}
