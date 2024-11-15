package logic

import (
	"context"
	"errors"
	"time"

	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/types"
	"github/lyr1cs/v0/oj-exam-backend/common/ctxinfo"

	"github.com/golang-jwt/jwt/v4"
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
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, info.StudentNumber)
	resp = &types.CheckResponse{
		StudentID: info.StudentNumber,
		Name:      info.Name,
		Token:     accessToken,
	}
	return
}

func getJwtToken(secretKey string, iat, seconds int64, Id string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxinfo.CtxJwtKey] = Id
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
