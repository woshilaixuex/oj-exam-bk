package logic

import (
	"context"

	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 开始考试
func NewExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExamLogic {
	return &ExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExamLogic) Exam(req *types.StartRequest) (resp *types.StartResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
