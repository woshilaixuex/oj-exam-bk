package load

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github/lyr1cs/v0/oj-exam-backend/app/api/internal/distribute/rule"
	"github/lyr1cs/v0/oj-exam-backend/common/constm"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description:  redis操作
 * @Date: 2024-11-11 18:39
 */

var (
	//go:embed popscript.lua
	atomPopAndMinusLuaScript string
	//go:embed optchangescript.lua
	strategyChangeLuaScript string
)

type LoadRedisService interface {
	InitUploadExamUsersToRedis(key string, users []ExamUser) error
	LPopAndDecrement(key string) (string, int64, error)
	AddToUniqueSet(key, randomSuffix string) (int, error)
	GetStrategy(key string) (rule.StrategyModel, error)
	ChangeStrategy(model rule.StrategyModel) error
}

func NewLoadRedisService(rediscli *redis.Redis) LoadRedisService {
	return &RedisService{
		Redis: rediscli,
	}
}

// 向redis提交list数据
func (rs *RedisService) InitUploadExamUsersToRedis(key string, users []ExamUser) error {

	for _, user := range users {
		data, err := json.Marshal(user)
		if err != nil {
			return fmt.Errorf("failed to marshal ExamUser: %w", err)
		}

		if _, err := rs.Redis.Sadd(key, string(data)); err != nil {
			return fmt.Errorf("failed to push data to Redis: %w", err)
		}
	}
	return nil
}

// 获取账号以及剩余账号数量
func (rs *RedisService) LPopAndDecrement(key string) (string, int64, error) {
	result, err := rs.Redis.Eval(atomPopAndMinusLuaScript, []string{key})
	if err != nil {
		return "", 0, fmt.Errorf("failed to execute Lua script: %w", err)
	}
	// 数据解析
	values := result.([]interface{})
	poppedValue := values[0]
	newLength := values[1].(int64)
	var value string
	if poppedValue != nil {
		value = poppedValue.(string)
	}
	return value, newLength, nil
}

// 获取账号ID
func (rs *RedisService) AddToUniqueSet(key, randomSuffix string) (int, error) {
	result, err := rs.Redis.Sadd(key, randomSuffix)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// 获取策略信息
func (rs *RedisService) GetStrategy(key string) (rule.StrategyModel, error) {
	strmodel, err := rs.Redis.Get(key)
	if err != nil {
		// 如果未查询到键值对
		if err == redis.Nil {

		} else {
			logx.Errorf("Get redis strategy err: %w", err)
			return rule.RETURN_ERROR, err
		}
	}
	intmodel, err := strconv.ParseUint(strmodel, 10, 64)
	if err != nil {
		logx.Errorf("Get redis strategy model trans err: %w", err)
		return rule.RETURN_ERROR, err
	}
	return rule.StrategyModel(intmodel), nil
}

// 修改策略信息
func (rs *RedisService) ChangeStrategy(model rule.StrategyModel) error {
	strModel := model.String()
	_, err := rs.Redis.Eval(strategyChangeLuaScript, []string{constm.STRATEGY_KEY}, strModel)
	if err != nil {
		logx.Errorf("Set redis strategy model err: %w", err)
		return err
	}
	return nil
}
