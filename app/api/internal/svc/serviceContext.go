package svc

import (
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/config"
	"github/lyr1cs/v0/oj-exam-backend/app/model/enroll"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	Redis            *redis.Redis
	EnrollTableModel enroll.EnrollTableModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:           c,
		Redis:            redis.MustNewRedis(c.Redis),
		EnrollTableModel: enroll.NewEnrollTableModel(conn),
	}
}
