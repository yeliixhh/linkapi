package consts

import "net/http"

var (
	// http状态码
	HTTP_CODE = http.StatusOK

	// token异常
	TGOKEN_ERROR_CODE = 10001

	// token过期
	TOKEN_EXPIRES_CODE = 10002

	CONTEXT_USER_KEY = "user_info"
)
