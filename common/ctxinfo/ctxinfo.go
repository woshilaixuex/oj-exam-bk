package ctxinfo

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-11-13 23:13
 */
var CtxJwtKey = "jwt_id"

func GetIdFromCtx(ctx context.Context) string {
	var id string
	if jsonId, ok := ctx.Value(CtxJwtKey).(string); ok {
		return jsonId
	}
	logx.WithContext(ctx).Error("Get id from ctx err")
	return id
}
