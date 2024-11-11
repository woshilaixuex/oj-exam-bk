package thread

import "github.com/zeromicro/go-zero/core/threading"

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: 协程池
 * @Date: 2024-11-11 23:27
 */

var (
	ThreadNum int = 20
	Pool      *threading.TaskRunner
)

func init() {
	Pool = threading.NewTaskRunner(ThreadNum)
}
