package singlelist

// 单链表
type SingleList[T comparable] struct {
	Head *node[T]
	Rear *node[T]
	size uint
}

// 单链表构造函数
func NewSingleList[T comparable]() *SingleList[T] {
	sl := new(SingleList[T])
	sl.Head = nil
	sl.Rear = nil
	sl.size = 0
	return sl
}

// 获取链表的大小
func (sl SingleList[T]) Size() uint {
	return sl.size
}

// 判断链表是否为空链表
func (sl SingleList[T]) Empty() bool {
	return sl.size == 0
}

// 头插入
func (sl SingleList[T]) Preappend(data T) *node[T] {

	return nil
}

// 尾插入
func (sl SingleList[T]) Append(data T) *node[T] {
	return nil
}
