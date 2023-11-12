package singlelist

import (
	"fmt"
	"testing"
)

func TestSingleList(t *testing.T) {
	sl := NewSingleList[int]()
	sl.Preappend(66)
	sl.Preappend(33)
	sl.Preappend(99)
	fmt.Printf("size = %d\n", sl.Size())
	sl.Show()
}
