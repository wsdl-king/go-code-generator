package result

const (
	unkonw     = iota
	success    = 10000
	error50000 = 50000
	error50001 = 50001
)

var Msg = map[int]string{
	success:    "成功",
	error50000: "默认失败",
	error50001: "数据返回失败",
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetMsg(code int) string {
	s, ok := Msg[code]
	if ok {
		return s
	}
	return Msg[error50000]
}

func SuccessResult(data interface{}) *Result {
	return &Result{Code: success, Msg: GetMsg(success), Data: data}
}

func ErrorResult(data interface{}, code int) *Result {
	return &Result{Code: code, Msg: GetMsg(code), Data: data}
}
