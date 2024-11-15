package thread

import "github.com/zeromicro/go-zero/core/threading"

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: 协程池
 * @Date: 2024-11-11 23:27
 */

type gorpool struct {
	ThreadNum int
	Runners   *threading.TaskRunner
}

var Pool gorpool

func init() {
	Pool.ThreadNum = 20
	Pool.Runners = threading.NewTaskRunner(Pool.ThreadNum)
}
