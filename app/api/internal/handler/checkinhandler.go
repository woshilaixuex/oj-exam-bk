package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/logic"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/types"
)

func CheckInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckInRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCheckInLogic(r.Context(), svcCtx)
		resp, _ := l.CheckIn(&req)
		w.Header().Set("Content-Type", "image/png") // 如果是 PNG 格式
		w.Write(resp.QrBytes)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
