package singlelist

import (
	"fmt"
	"testing"
)

// 测试头部插入
func TestPreappend(t *testing.T) {
	sl := NewSingleList[int]()
	sl.Preappend(66)
	sl.Preappend(33)
	sl.Preappend(99)
	fmt.Printf("size = %d\n", sl.Size())
	sl.Show()
}

// 测试尾部插入
func TestAppend(t *testing.T) {
	sl := NewSingleList[int]()
	sl.Append(132)
	sl.Append(13)
	sl.Append(12)
	fmt.Printf("size = %d\n", sl.Size())
	sl.Show()
}

func TestInsert(t *testing.T) {
	sl := NewSingleList[int]()
	sl.Preappend(4)
	n := sl.Preappend(5)
	n = sl.Insert(88, n)
	n = sl.Insert(99, n)
	sl.Insert(101, n)
	sl.Insert(87, sl.rear)
	fmt.Printf("size = %d\n", sl.Size())
	sl.Show()
}

func TestInsert2(t *testing.T) {
	sl := NewSingleList[int]()
	sl.Insert(8, sl.rear)
	sl.Insert(9, sl.rear)
	sl.Insert(10, sl.rear)
	fmt.Printf("size = %d\n", sl.Size())
	sl.Show()
}
