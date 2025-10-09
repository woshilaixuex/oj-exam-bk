package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/woshilaixuex/third/onlinejudge"
	"github.com/zeromicro/go-zero/core/logx"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/load"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/types"
	"net/http"
	"strconv"
	"time"
)

type CheckOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type CheckoutFeiShuMsg struct {
	Userinfo *Content
	Score    string `json:"score"`
}

func NewCheckOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckOutLogic {
	return &CheckOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckOutLogic) CheckOut(req *types.CheckOutRequest) (resp *types.CheckOutResponse, err error) {
	// todo: add your logic here and delete this line
	webhookURL := l.svcCtx.Config.WebHookUrl
	studentId := load.FindStudenIdtByAccountAndPassword(l.svcCtx.Config.DB.DataSource, req.Account)
	if studentId == "" {
		return &types.CheckOutResponse{
			StudentID: "账号或密码错误",
			Score:     0,
		}, nil
	}

	if studentId == "1" {
		return &types.CheckOutResponse{
			StudentID: "该账号还没签到",
			Score:     0,
		}, nil
	}
	tools := onlinejudge.NewOjTools()
	params := onlinejudge.GetRankParams{
		Offset:       "0",
		Limit:        "200",
		ContestId:    l.svcCtx.Config.ContestId,
		ForceRefresh: "0",
	}
	response, err := tools.GetExamRank(params)
	if err != nil {
		return nil, err
	}
	for _, result := range response.Results {
		if result.User.Username == req.Account {
			score := result.TotalScore
			load.SaveScore(l.svcCtx.Config.DB.DataSource, studentId, score)
			user, _ := l.svcCtx.EnrollTableModel.FindOneByStudentId(context.Background(), studentId)
			hour, _, _ := time.Now().Clock()
			var when string
			if hour >= 13 {
				when = "下午"
			} else {
				when = "上午"
			}
			userInfo := &CheckoutFeiShuMsg{
				Userinfo: &Content{
					StuNum:  studentId,
					StuName: user.Name,
					StuTime: when,
				},
				Score: strconv.Itoa(int(score)),
			}
			err := sendScoreToFeiShu(webhookURL, userInfo)
			if err != nil {
				logx.Errorf("%s", err.Error())
			}
			resp = &types.CheckOutResponse{
				StudentID: studentId,
				Score:     score,
			}
			return resp, nil
		}
	}
	return &types.CheckOutResponse{
		StudentID: "请确认此账号是否参加比赛",
		Score:     0,
	}, nil
}

func sendScoreToFeiShu(webhookURL string, userInfo *CheckoutFeiShuMsg) error {
	message := FeishuMessage{
		MsgType: "text",
		Content: TextContent{
			Text: fmt.Sprintf("学号: %s\n姓名: %s\n考试时间: %s\n成绩: %s",
				userInfo.Userinfo.StuNum, userInfo.Userinfo.StuName, userInfo.Userinfo.StuTime, userInfo.Score),
		},
	}

	jsonData, _ := json.Marshal(message)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}
	return nil
}
