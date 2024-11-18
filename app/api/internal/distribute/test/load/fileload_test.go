package load_test

import (
	"context"
	"flag"
	"fmt"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/config"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/load"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-11-10 16:45
 */
var configFile = flag.String("f", "../../../../etc/exam-api.yaml", "the config file")

func TestLoad(t *testing.T) {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svc.NewServiceContext(c)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	load.InitLoadServer(ctx, "./data/users_test.csv")
	fmt.Print("haha")
	go func() {
		var s string
		fmt.Scanln(&s)
		cancel()
	}()
	<-ctx.Done()
}

func TestPop(t *testing.T) {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	sctx := svc.NewServiceContext(c)
	data, m, err := sctx.Redis.PopSetAndDecrement("csd_test_l")
	if err != nil {
		t.Fatalf("err: %v", err)
		return
	}
	fmt.Println(m)
	fmt.Println(data)
}
