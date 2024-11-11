package load

import (
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-11-11 18:39
 */

var (
	Redis *redis.Redis
)

func UploadExamUsersToRedis(key string, users []ExamUser) error {

	for _, user := range users {
		data, err := json.Marshal(user)
		if err != nil {
			return fmt.Errorf("failed to marshal ExamUser: %w", err)
		}

		if _, err := Redis.Lpush(key, string(data)); err != nil {
			return fmt.Errorf("failed to push data to Redis: %w", err)
		}
	}

	return nil
}
