package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/load"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/types"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type CheckOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type ContestRankResponse struct {
	Error interface{} `json:"error"`
	Data  struct {
		Results []struct {
			User struct {
				Username string `json:"username"`
			} `json:"user"`
			TotalScore int `json:"total_score"`
		} `json:"results"`
	} `json:"data"`
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
	webhookURL := "https://open.feishu.cn/open-apis/bot/v2/hook/f0eaf3ec-c617-4c27-863a-a9448a729ad8"
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
	url := "https://oj.dsstudio.tech/api/contest_rank?offset=0&limit=" + strconv.Itoa(200) + "&contest_id=23&force_refresh=0"
	//url := "https://oj.dsstudio.tech/api/profile?username=" + req.Account
	// 创建请求
	rq, _ := http.NewRequest("GET", url, nil)

	// 添加 Cookie 到请求头
	rq.Header.Add("Cookie", "_ga=GA1.2.1938032973.1728895971; csrftoken=V2XYFHfhKH7IrDBM6fEtG5VR3IVR7Zup89Hqd5V1m2JLiyjO8gULhq0jaaiheS9T; sessionid=hsglloqgnxugjxxppmpcljeuunpd15nv; _gid=GA1.2.693926491.1756951783; _gat=1; _ga_59QEB25NR7=GS2.2.s1756951784$o26$g1$t1756951878$j26$l0$h0")

	// 发送请求
	client := &http.Client{}
	rp, _ := client.Do(rq)

	// 读取响应
	body, _ := ioutil.ReadAll(rp.Body)
	//l.Infof("OJ原始响应：%s", string(body))
	var response ContestRankResponse
	if err = json.Unmarshal(body, &response); err != nil {
		fmt.Printf("解析JSON失败: %v\n", err)
		return
	}
	for _, result := range response.Data.Results {
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
				Score:     response.Data.Results[0].TotalScore,
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
