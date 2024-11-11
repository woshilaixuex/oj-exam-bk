package load_test

import (
	"context"
	"flag"
	"fmt"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/load"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-11-10 16:45
 */
var configFile = flag.String("f", "../../../../etc/exam-api.yaml", "the config file")

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
	Redis redis.RedisConf
	DB    struct {
		DataSource string
	}
}

func TestLoad(t *testing.T) {
	flag.Parse()
	var c Config
	conf.MustLoad(*configFile, &c)
	load.Redis = redis.MustNewRedis(c.Redis)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	load.InitLoadServer(ctx, "./data/user_test.csv")
	go func() {
		var s string
		fmt.Scanln(&s)
		cancel()
	}()
	<-ctx.Done()
}
