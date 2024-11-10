package httpm

import (
	"context"
	"net/http"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: http上下文替换
 * @Date: 2024-11-05 10:28
 */

func Succeedhandler(ctx context.Context, body any) any {
	return &HttpResponse{
		Code: http.StatusOK,
		Msg:  SUCCEED_MSG,
		Data: body,
	}
}

func DTErrorHandlerCtx(context.Context, error) (int, any) {
	return 0, &HttpResponse{}
}
func PPrrorHandlerCtx(context.Context, error) (int, any) {
	return 0, &HttpResponse{}
}
