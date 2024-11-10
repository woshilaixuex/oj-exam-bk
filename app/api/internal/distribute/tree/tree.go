package tree

import (
	"context"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-11-07 12:52
 */
type DoEngine func()
type RuleNode interface {
	Dologic(logger logx.Logger, ctx context.Context, svcCtx *svc.ServiceContext) (NodeId, error)
}

type TreeEngine interface {
	StartEngine()
	DoEngine()
}
type Node interface {
	DoNextNode()
}
type Engine struct {
	fistnode node
	svcCtx   *svc.ServiceContext
}
type node struct {
	nextnode map[NodeId]node
	RuleNode
}

func NewTreeEngine(svcCtx *svc.ServiceContext) TreeEngine {

	return &Engine{
		svcCtx: svcCtx,
	}
}
func (e Engine) StartEngine() {

}
func (e Engine) DoEngine() {

}
