package svc

import (
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/config"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/load"
	"github/lyr1cs/v0/oj-exam-backend/app/model/enroll"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	Redis            load.LoadRedisService
	EnrollTableModel enroll.EnrollTableModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	Service := load.NewLoadRedisService(redis.MustNewRedis(c.Redis))
	load.InitRedisService = Service
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:           c,
		Redis:            Service,
		EnrollTableModel: enroll.NewEnrollTableModel(conn),
	}
}
