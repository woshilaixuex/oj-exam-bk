package load

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github/lyr1cs/v0/oj-exam-backend/common/constm"
	"strconv"

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

type LoadService interface {
	UploadExamUsersToRedis(key string, users []constm.ExamUser) error
	PopSetAndDecrement(key string) (string, int64, error)
	AddToUniqueSet(key, randomSuffix string) (int, error)
	GetStrategy(key string) (constm.StrategyModel, error)
	ChangeStrategy(model constm.StrategyModel) error
}
type LoadDeploy struct {
	Redis *redis.Redis
}

func NewLoadRedisService(rediscli *redis.Redis) LoadService {
	return &LoadDeploy{
		Redis: rediscli,
	}
}

// 初始化向redis提交list数据
func (rs *LoadDeploy) UploadExamUsersToRedis(key string, users []constm.ExamUser) error {
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
func (rs *LoadDeploy) PopSetAndDecrement(key string) (string, int64, error) {
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
func (rs *LoadDeploy) AddToUniqueSet(key, randomSuffix string) (int, error) {
	result, err := rs.Redis.Sadd(key, randomSuffix)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// 获取策略信息
func (rs *LoadDeploy) GetStrategy(key string) (constm.StrategyModel, error) {
	strmodel, err := rs.Redis.Get(key)
	if err != nil {
		// 如果未查询到键值对
		if err == redis.Nil {

		} else {
			return constm.RETURN_ERROR, fmt.Errorf("failed to get redis strategy err: %w", err)
		}
	}
	intmodel, err := strconv.ParseUint(strmodel, 10, 64)
	if err != nil {
		return constm.RETURN_ERROR, fmt.Errorf("failed to get redis strategy model trans err: %w", err)
	}
	return constm.StrategyModel(intmodel), nil
}

// 修改策略信息
func (rs *LoadDeploy) ChangeStrategy(model constm.StrategyModel) error {
	strModel := model.String()
	_, err := rs.Redis.Eval(strategyChangeLuaScript, []string{constm.STRATEGY_KEY}, strModel)
	if err != nil {
		return fmt.Errorf("failed to set redis strategy model err: %w", err)
	}
	return nil
}
