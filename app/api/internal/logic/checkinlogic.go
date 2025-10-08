package logic

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/config"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/load"
	"net/http"
	"time"

	"github.com/skip2/go-qrcode"
	"github.com/zeromicro/go-zero/core/logx"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/types"
)

type CheckInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type FeishuMessage struct {
	MsgType string      `json:"msg_type"`
	Content TextContent `json:"content"`
}

type Content struct {
	StuNum  string
	StuName string
	StuTime string
}

type TextContent struct {
	Text string `json:"text"`
}

var c config.Config

func NewCheckInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckInLogic {
	return &CheckInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckInLogic) CheckIn(req *types.CheckInRequest) (resp *types.CheckInResponse, err error) {
	// todo: add your logic here and delete this line
	resp = &types.CheckInResponse{}
	c = l.svcCtx.Config
	webhookURL := l.svcCtx.Config.WebHookUrl
	info, err := l.svcCtx.EnrollTableModel.FindOneByStudentId(l.ctx, req.StudentID)
	hour, _, _ := time.Now().Clock()
	var time string
	if hour >= 13 {
		time = "下午"
	} else {
		time = "上午"
	}
	content := Content{
		StuNum:  info.StudentNumber,
		StuName: info.Name,
		StuTime: time,
	}
	if err := sendFeishuMessage(webhookURL, &content); err != nil {
		fmt.Println("发送消息失败:", err)
		return nil, err
	}
	qrBytes, err := generateAuthQR(info.StudentNumber)
	if err != nil {
		return nil, err
	}
	resp.QrBytes = qrBytes
	return resp, nil
}

func sendFeishuMessage(webhookURL string, userInfo *Content) error {
	message := FeishuMessage{
		MsgType: "text",
		Content: TextContent{
			Text: fmt.Sprintf("学号: %s\n姓名: %s\n考试时间: %s",
				userInfo.StuNum, userInfo.StuName, userInfo.StuTime),
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

func generateAuthQR(userID string) ([]byte, error) {
	// 1. 生成随机Token
	token := generateToken()

	// 2. 存储Token与用户关联(使用Redis或数据库)
	err := load.InitRedisService.StoreToken(userID, token)
	if err != nil {
		return nil, err
	}
	// 3. 生成带Token的URL
	authURL := fmt.Sprintf("http://"+c.OjDomain.Domain+"/exam?token=%s", token)

	// 4. 生成二维码
	return qrcode.Encode(authURL, qrcode.Medium, 256)
}

// 生成随机Token
func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
