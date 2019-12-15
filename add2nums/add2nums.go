package add2nums

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{(l1.Val + l2.Val) % 10, nil}
	currentPoint := r
	carry := (l1.Val + l2.Val) / 10
	l1 = l1.Next
	l2 = l2.Next

	for l1 != nil || l2 != nil || carry > 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		currentPoint.Next = &ListNode{(v1 + v2 + carry) % 10, nil}
		currentPoint = currentPoint.Next
		carry = (v1 + v2 + carry) / 10
	}
	return r
}
