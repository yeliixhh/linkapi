package interfaces

import (
	"context"

	"github.com/yeliixhh/linkapi/internal/types"
)

type SysTokenRepository interface {

	// 创建tgoken
	CreateSysToken(ctx context.Context, token *types.SysToken) error

	// 查询token通过token
	QuerySysTokenByToken(ctx context.Context, token string) (*types.SysToken, error)
}
