package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/yeliixhh/linkapi/internal/consts"
)

// http返回结果
func HttpResponse(c *gin.Context, data any) {
	c.JSON(consts.HTTP_CODE, data)
}
