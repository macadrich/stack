package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push("aaa")
	s.Push("bbb")
	s.Push("ccc")
	
	for i := 0; i < s.Length()+1; i++ {
		s.Pop()
		if s.Pop() == nil {
			t.Errorf("unexpected value %v", s.Pop())
		}
	}
}