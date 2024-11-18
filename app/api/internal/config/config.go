package config

import (
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/ojclient"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis redis.RedisConf
	DB    struct {
		DataSource string
	}
	OJClinetConfig ojclient.OJClientConfig
}
