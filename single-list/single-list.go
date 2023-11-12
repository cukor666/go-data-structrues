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

// 头插入
func (sl *SingleList[T]) Preappend(data T) *node[T] {
	n := NewNode[T](data, nil)
	if sl.size == 0 {
		sl.head = n
		sl.rear = n
	} else {
		n.Next = sl.head
		sl.head = n
	}
	sl.size++
	return n
}

// 尾插入
func (sl *SingleList[T]) Append(data T) *node[T] {
	n := NewNode[T](data, nil)
	if sl.size == 0 {
		sl.head = n
		sl.rear = n
	} else {
		sl.rear.Next = n
		sl.rear = n
	}
	sl.size++
	return n
}
