package rule

import "strconv"

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-11-14 21:04
 */
type StrategyModel uint64

const (
	// 默认直接从redis中获取key
	DEFULT_RULE StrategyModel = 0 + iota
	PRODUCT_RULE
	RETURN_ERROR
)

func (modle *StrategyModel) String() string {
	return strconv.FormatUint(uint64(*modle), 10)
}
