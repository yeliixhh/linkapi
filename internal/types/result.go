package types

var (
	SUCCESS_CODE = 200
	FAILT_CODE   = 500
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// 成功
func SuccessResult(data any) *Result {
	return &Result{Code: SUCCESS_CODE, Message: "success", Data: data}
}

// 失败返回体
func FailResult(message string) *Result {
	return &Result{Code: FAILT_CODE, Message: message}
}

func FailResultWithCode(code int, message string) *Result {
	return &Result{Code: code, Message: message}
}

// 失败返回体
func FailResultError(e error) *Result {
	if e == nil {
		return &Result{Code: FAILT_CODE, Message: "fail"}
	}
	return &Result{Code: FAILT_CODE, Message: e.Error()}
}

// 失败返回体
func FailResultWithData(message string, data any) *Result {
	return &Result{Code: FAILT_CODE, Message: message, Data: data}
}
