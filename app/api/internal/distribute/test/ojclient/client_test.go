package ojclient_test

import (
	"flag"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/config"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/ojclient"
	"github/lyr1cs/v0/oj-exam-backend/common/constm"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-11-18 20:05
 */
var configFile = flag.String("f", "../../../../etc/exam-api.yaml", "the config file")

func TestPushAccount(t *testing.T) {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	client := ojclient.NewOJClient(&c.OJClinetConfig)
	examUsers := []constm.ExamUser{
		{
			Account:  "csd00051",
			Password: "xHu6fL32",
			Email:    "csd00051@exam.com",
			Name:     "csd00051",
		},
		{
			Account:  "csd00052",
			Password: "yJt7gK91",
			Email:    "csd00052@exam.com",
			Name:     "csd00052",
		},
	}

	client.PushAccount(examUsers)
}
