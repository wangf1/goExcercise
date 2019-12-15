package add2nums

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{5, nil}
	tmp := l1
	tmp.Next = &ListNode{6, nil}
	tmp = tmp.Next
	tmp.Next = &ListNode{7, nil}
	tmp = tmp.Next

	l2 := &ListNode{6, nil}
	tmp = l2
	tmp.Next = &ListNode{1, nil}
	tmp = tmp.Next
	tmp.Next = &ListNode{8, nil}
	tmp = tmp.Next

	r := addTwoNumbers(l1, l2)
	if r.Val != 1 {
		t.Fatalf("First digit is not 1: %v", r.Val)
	}
	r = r.Next
	if r.Val != 8 {
		t.Fatalf("Second digit is not 8: %v", r.Val)
	}
	r = r.Next
	if r.Val != 5 {
		t.Fatalf("Third digit is not 5: %v", r.Val)
	}
	r = r.Next
	if r.Val != 1 {
		t.Fatalf("Forth digit is not 1: %v", r.Val)
	}
}
