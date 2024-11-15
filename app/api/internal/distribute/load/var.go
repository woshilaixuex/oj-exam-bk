package load

import (
	"encoding/csv"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/rule"
	"io"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: 数据绑定与实体模型
 * @Date: 2024-11-09 21:50
 */

type ExamUser struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func parseCSVData(data []byte) ([]ExamUser, error) {
	reader := csv.NewReader(strings.NewReader(string(data)))

	var users []ExamUser
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logx.Errorf("failed to read CSV record: %v", err)
			return nil, err
		}
		// CSV 的列顺序为：account, password, email, name
		if len(record) >= 4 {
			user := ExamUser{
				Account:  record[0],
				Password: record[1],
				Email:    record[2],
				Name:     record[3],
			}
			users = append(users, user)
		} else {
			logx.Statf("skipping invalid CSV record: %v", record)
			continue
		}
	}
	return users, nil
}

type RedisService struct {
	Redis *redis.Redis
}
type LockModel struct {
	Model   rule.StrategyModel
	Version string
}
