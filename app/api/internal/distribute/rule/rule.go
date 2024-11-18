package rule

import (
	"encoding/json"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/load"
	"github/lyr1cs/v0/oj-exam-backend/common/constm"
	"github/lyr1cs/v0/oj-exam-backend/common/thread"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: 生产用户的业务编排
 * @Date: 2024-11-07 12:52
 */
type RuleService interface {
	DoDefultRule(userId string) (*constm.ExamUser, error)
	DoProductRule(userId string) (*constm.ExamUser, error)
}

type RuleDeploy struct {
	LoadService load.LoadService
	Redis       *redis.Redis
}

func NewRuleService(LoadService load.LoadService, Redis *redis.Redis) RuleService {
	return &RuleDeploy{
		LoadService: LoadService,
		Redis:       Redis,
	}
}
func (deploy *RuleDeploy) DoDefultRule(userId string) (*constm.ExamUser, error) {
	lockKey := constm.RULE_USER_LOCK + userId
	// 创建 Redis 锁(gozero这个锁本来就防误删的)
	lock := redis.NewRedisLock(deploy.Redis, lockKey)
	lock.SetExpire(constm.DEFULT_RULE_USER_EXP)
	acquired, err := lock.Acquire()
	if err != nil {
		logx.Errorf("Failed to acquire lock for user %s: %v", userId, err)
		return nil, err
	}
	if !acquired {
		logx.Infof("Lock not acquired for user %s, operation is idempotent", userId)
		return nil, err
	}
	defer func() {
		released, err := lock.Release()
		if err != nil || !released {
			logx.Errorf("Failed to release lock for user %s: %v", userId, err)
		}
	}()
	var account constm.ExamUser
	data, num, err := deploy.LoadService.PopSetAndDecrement("csd_test_l")
	if err != nil {
		logx.Errorf("pop account for user %s: %v", userId, err)
		return nil, err
	}
	switch {
	case num <= 0:
		{
			key := "csd_test_l"
			thread.Pool.Schedule(func() {
				deploy.ProductAccountToRedis(key)
			})
			return deploy.DoProductRule(userId)
		}
	case num <= 10 && num > 0:
		{
			deploy.LoadService.ChangeStrategy(constm.PRODUCT_RULE)
		}
	}
	json.Unmarshal([]byte(data), &account)
	return &account, nil
}

func (deploy *RuleDeploy) DoProductRule(userId string) (*constm.ExamUser, error) {
	var account constm.ExamUser
	// 生成一个账号
	data := load.ProductData(1)
	json.Unmarshal([]byte(data), &account)
	return &account, nil
}

func (deploy *RuleDeploy) ProductAccountToRedis(key string) error {
	examUsers, err := load.ParseCSVData(load.ProductData(constm.DEFULT_RULE_PRO_NUM))
	if err != nil {
		return err
	}
	err = deploy.LoadService.UploadExamUsersToRedis(key, examUsers)
	if err != nil {
		return err
	}
	deploy.LoadService.ChangeStrategy(constm.DEFULT_RULE)
	return nil
}
