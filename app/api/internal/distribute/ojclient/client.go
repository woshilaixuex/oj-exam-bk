package ojclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/lyr1cs/v0/oj-exam-backend/common/constm"
	"io"
	"net/http"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: OJ客户端（Oj端操作封装）
 * @Date: 2024-11-17 17:58
 */
type OJClient interface {
	PushAccount(examusers []constm.ExamUser)
}
type OJClientConfig struct {
	Url       string
	Origin    string
	Sessionid string
	CsrfToken string
}
type OJClientDeploy struct {
	url       string
	origin    string
	sessionid string
	csrfToken string
}

func NewOJClient(config *OJClientConfig) OJClient {
	return &OJClientDeploy{
		url:       config.Url,
		origin:    config.Origin,
		sessionid: config.Sessionid,
		csrfToken: config.CsrfToken,
	}
}

// oj的鉴权这些东西都不能少
func (deploy *OJClientDeploy) setHead(req *http.Request) {
	req.Header.Add("X-Csrftoken", deploy.csrfToken)
	req.Header.Add("Origin", deploy.origin)
	req.Header.Add("Referer", deploy.url)
	req.Header.Add("Cookie", fmt.Sprintf("csrftoken=%s;sessionid=%s", deploy.csrfToken, deploy.sessionid))
	req.Header.Add("Content-Type", "application/json")
}
func (deploy *OJClientDeploy) PushAccount(examusers []constm.ExamUser) {
	var users [][]string
	for _, user := range examusers {
		users = append(users, []string{
			user.Account,
			user.Password,
			user.Email,
			user.Name,
		})
	}

	payloadData := map[string][][]string{
		"users": users,
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	payload := bytes.NewReader(payloadBytes)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, deploy.origin+constm.PUSH_ACCOUNT, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	deploy.setHead(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
