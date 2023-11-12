package singlelist

// 单链表的节点
type node[T comparable] struct {
	Data T
	Next *node[T]
}

// 新建一个节点，构造函数
func NewNode[T comparable](data T, next *node[T]) *node[T] {
	n := new(node[T])
	n.Data = data
	n.Next = next
	return n
}
