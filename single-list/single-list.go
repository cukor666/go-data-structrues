package singlelist

import "fmt"

// 单链表
type SingleList[T comparable] struct {
	head *node[T]
	rear *node[T]
	size uint
}

// 单链表构造函数
func NewSingleList[T comparable]() *SingleList[T] {
	sl := new(SingleList[T])
	sl.head = nil
	sl.rear = nil
	sl.size = 0
	return sl
}

// 获取链表的大小
func (sl *SingleList[T]) Size() uint {
	return sl.size
}

// 判断链表是否为空链表
func (sl *SingleList[T]) Empty() bool {
	return sl.size == 0
}

// 遍历单链表
func (sl *SingleList[T]) Show() {
	for p := sl.head; p != nil; p = p.Next {
		fmt.Printf("%v ", p.Data)
	}
	fmt.Println()
}

// 根据地址插入，将data插入到addr指向的节点后面，如果addr为空则调用头部插入
func (sl *SingleList[T]) Insert(data T, addr *node[T]) *node[T] {
	if addr == nil {
		return sl.Preappend(data)
	}
	n := NewNode[T](data, addr.Next)
	addr.Next = n
	sl.size++
	return n
}

// 头插入
func (sl *SingleList[T]) Preappend(data T) *node[T] {
	n := NewNode[T](data, nil)
	if sl.size == 0 {
		sl.rear = n
	} else {
		n.Next = sl.head
	}
	sl.head = n
	sl.size++
	return n
}

// 尾插入
func (sl *SingleList[T]) Append(data T) *node[T] {
	n := sl.Insert(data, sl.rear)
	sl.rear = n
	return n
}

// 头部删除
func (sl *SingleList[T]) RemoveHead() *node[T] {
	if sl.size == 0 {
		return nil
	}
	p := sl.head
	sl.head = sl.head.Next
	p.Next = nil
	sl.size--
	return p
}

// 尾部删除
func (sl *SingleList[T]) RemoveRear() *node[T] {
	if sl.size == 0 {
		return nil
	}
	p := sl.head
	for p != nil && p.Next != sl.rear {
		p = p.Next
	}
	if p != nil {
		p.Next = nil
		sl.rear = p
		sl.size--
		return p
	}
	return nil
}
