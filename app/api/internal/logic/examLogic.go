package logic

import (
	"context"

	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/types"
	"github/lyr1cs/v0/oj-exam-backend/common/constm"
	"github/lyr1cs/v0/oj-exam-backend/common/ctxinfo"

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
	userId := ctxinfo.GetIdFromCtx(l.ctx)
	model, err := l.CheckStrategy()
	if err != nil {

	}
	l.DoStrategy(userId, model)
	resp = &types.StartResponse{
		StudentID: userId,
	}
	return
}

func (l *ExamLogic) CheckStrategy() (constm.StrategyModel, error) {
	return l.svcCtx.Redis.GetStrategy(constm.STRATEGY_KEY)
}
func (l *ExamLogic) DoStrategy(userId string, model constm.StrategyModel) {
	if model == constm.DEFULT_RULE {
		l.svcCtx.Rule.DoDefultRule(userId)
	}
	if model == constm.PRODUCT_RULE {
		l.svcCtx.Rule.DoProductRule(userId)
	}
}
