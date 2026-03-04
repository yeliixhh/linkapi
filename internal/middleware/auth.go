package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yeliixhh/linkapi/internal/consts"
	"github.com/yeliixhh/linkapi/internal/types"
	"github.com/yeliixhh/linkapi/internal/types/interfaces"
)

func Auth(userServer interfaces.UserService) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")

		user, err := userServer.ValidateToken(ctx, token)

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.JSON(http.StatusOK, types.FailResultWithCode(consts.TOKEN_EXPIRES_CODE, "token过期"))
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusOK, types.FailResultWithCode(consts.TGOKEN_ERROR_CODE, "token异常"))
			ctx.Abort()
			return
		}

		ctx.Set(consts.CONTEXT_USER_KEY, user)
	}

}
