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
	sl.head = new(node[T])
	sl.rear = new(node[T])
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
	for p := sl.Head(); p != nil; p = p.Next {
		fmt.Printf("%v ", p.Data)
	}
	fmt.Println()
}

// 设置头部节点
func (sl *SingleList[T]) setHead(n *node[T]) {
	sl.head.Next = n
}

// 获取头部
func (sl *SingleList[T]) Head() *node[T] {
	return sl.head.Next
}

// 设置尾部节点
func (sl *SingleList[T]) setRear(n *node[T]) {
	sl.rear.Next = n
	sl.rear = n
}

// 获取尾部
func (sl *SingleList[T]) Rear() *node[T] {
	return sl.rear.Next
}

// 头插入
func (sl *SingleList[T]) Preappend(data T) *node[T] {
	n := NewNode[T](data, nil)
	if sl.size == 0 {
		sl.setHead(n)
		sl.setRear(n)
	} else {
		n.Next = sl.Head()
		sl.setHead(n)
	}
	sl.size++
	return n
}

// 尾插入
func (sl *SingleList[T]) Append(data T) *node[T] {
	n := NewNode[T](data, nil)
	if sl.size == 0 {
		sl.setHead(n)
		sl.setRear(n)
	} else {
		sl.setRear(n)
	}
	sl.size++
	return n
}

// 指定位置插入，根据地址，将data插入到addr指向的节点后面
func (sl *SingleList[T]) Insert(data T, addr *node[T]) *node[T] {
	n := NewNode[T](data, addr.Next)
	addr.Next = n
	sl.size++
	return n
}
