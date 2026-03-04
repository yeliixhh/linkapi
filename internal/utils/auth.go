package utils

import (
	"context"
	"errors"

	"github.com/yeliixhh/linkapi/internal/consts"
	"github.com/yeliixhh/linkapi/internal/types"
)

func GetAuthUser(ctx context.Context) (*types.SysUser, error) {

	user, ok := ctx.Value(consts.CONTEXT_USER_KEY).(*types.SysUser)

	if !ok {
		return nil, errors.New("获取用户失败")
	}

	return user, nil
}
