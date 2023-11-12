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

func TestRemoveHead(t *testing.T) {
	sl := NewSingleList[int]()
	sl.Append(3)
	sl.Append(1)
	sl.Append(7)
	sl.Append(4)
	sl.Append(0)
	sl.Show()
	sl.RemoveHead()
	sl.RemoveHead()
	sl.Show()
}

func TestRemoveRear(t *testing.T) {
	sl := NewSingleList[int]()
	sl.Append(3)
	sl.Append(1)
	sl.Append(7)
	sl.Append(4)
	sl.Append(0)
	sl.Show()
	sl.RemoveRear()
	sl.RemoveRear()
	sl.Show()
	fmt.Printf("sl.size: %v\n", sl.size)
}

func TestAddress(t *testing.T) {
	sl := NewSingleList[int]()
	sl.Append(3)
	sl.Append(1)
	sl.Append(7)
	sl.Append(4)
	sl.Append(0)
	sl.Show()
	res := sl.Address(sl.head, sl.rear, []int{1, 4})
	fmt.Printf("address: %v\n", res)
	for _, v := range res {
		fmt.Printf("v = %v", v)
	}
}

func TestRemoveUseAddress(t *testing.T) {
	sl := NewSingleList[int]()
	sl.Append(3)
	sl.Append(1)
	sl.Append(5)
	sl.Append(7)
	sl.Append(8)
	sl.Append(4)
	sl.Append(0)
	sl.Show()
	res := sl.Address(sl.head, sl.rear, []int{1, 7, 4})
	n := len(res)
	for i := 0; i < n; i++ {
		sl.Remove(res[i])
	}
	sl.Show()
}
