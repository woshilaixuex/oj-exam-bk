package constm

import "strconv"

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-11-14 21:04
 */
type StrategyModel uint64

func (modle *StrategyModel) String() string {
	return strconv.FormatUint(uint64(*modle), 10)
}

const (
	// 默认直接从redis中获取key
	DEFULT_RULE StrategyModel = 0 + iota
	PRODUCT_RULE
	RETURN_ERROR
)
const (
	DEFULT_RULE_USER_EXP = 5 // second 秒
	DEFULT_RULE_PRO_NUM  = 50
)
const (
	PUSH_ACCOUNT = "api/admin/user"
)

type ExamUser struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type RuleService interface {
}

type RuleDeploy struct {
}
