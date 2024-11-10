package logic

import (
	"context"
	"errors"

	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 检查学生信息
func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req *types.CheckRequest) (resp *types.CheckResponse, err error) {
	// todo: add your logic here and delete this line
	if req.StudentID == "" {
		return nil, errors.New("student id is null")
	}
	logx.Debug("test sql")
	info, err := l.svcCtx.EnrollTableModel.FindOneByStudentId(l.ctx, req.StudentID)
	if err != nil {
		return nil, err
	}
	resp = &types.CheckResponse{
		StudentID: info.StudentNumber,
		Name:      info.Name,
	}
	return
}
