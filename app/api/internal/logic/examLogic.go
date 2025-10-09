package logic

import (
	"context"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/load"

	"github.com/zeromicro/go-zero/core/logx"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/types"
	"github/lyr1cs/v0/oj-exam-backend/common/constm"
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
	//userId := ctxinfo.GetIdFromCtx(l.ctx)
	userId := load.InitRedisService.FindUserByToken(req.Token)
	//model, err := l.CheckStrategy()
	//if err != nil {
	//
	//}
	examUser := load.FindExamUser(l.svcCtx.Config.DB.DataSource, userId)
	//examUser := l.DoStrategy(userId, model)
	resp = &types.StartResponse{
		StudentID: userId,
		Account:   examUser.Account,
		Password:  examUser.Password,
	}
	return
}

func (l *ExamLogic) CheckStrategy() (constm.StrategyModel, error) {
	return l.svcCtx.Redis.GetStrategy(constm.STRATEGY_KEY)
}
func (l *ExamLogic) DoStrategy(userId string, model constm.StrategyModel) *constm.ExamUser {
	var examUser *constm.ExamUser
	if model == constm.DEFULT_RULE {
		examUser, _ = l.svcCtx.Rule.DoDefultRule(userId)
	}
	if model == constm.PRODUCT_RULE {
		examUser, _ = l.svcCtx.Rule.DoProductRule(userId)
	}
	return examUser
}
