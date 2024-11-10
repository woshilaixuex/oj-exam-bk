package tree

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-11-07 13:15
 */
type NodeId uint32

const (
	DO_FINISH NodeId = iota
	DO_PUSH_RESOUCE_NODE
	DO_DISTRIBUTE_NODE
	DO_CHECK_NODE
)
