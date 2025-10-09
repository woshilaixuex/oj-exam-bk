package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/config"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/handler"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/svc"
	_ "github/lyr1cs/v0/oj-exam-backend/common/thread"
)

var configFile = flag.String("f", "etc/exam-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	//wd, _ := os.Getwd()
	//logx.Infof("当前工作目录: %s", wd)
	//logx.Infof("配置文件路径: %s", *configFile) // 输出 etc/exam-api.yaml
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 开发阶段debug测试
	logx.SetLevel(logx.DebugLevel)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
