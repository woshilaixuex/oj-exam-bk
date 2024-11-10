package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/logic"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/types"
)

// 开始考试
func ExamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StartRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewExamLogic(r.Context(), svcCtx)
		resp, err := l.Exam(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
