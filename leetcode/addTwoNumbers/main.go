package main

import "fmt"

//给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0开头。

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	return fmt.Sprintf("%d, %v", l.Val, l.Next)
}

//递归方案
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil

	} else if l2 == nil {
		return addTwoNumbers1(l1, &ListNode{Val: 0})
	} else if l1 == nil {
		return addTwoNumbers1(&ListNode{Val: 0}, l2)
	}
	v := l1.Val + l2.Val
	(*l1).Val = v % 10
	if n := v / 10; n == 1 {
		if l1.Next != nil {
			(*l1.Next).Val += n
		} else {
			(*l1).Next = &ListNode{Val: n}
		}
	}
	l1.Next = addTwoNumbers1(l1.Next, l2.Next)
	return l1
}

//循环方案
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail = &ListNode{}
	head := tail
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		n := n1 + n2 + tail.Val
		if m := n / 10; m > 0 {
			*tail = ListNode{n % 10, &ListNode{Val: m}}
		} else {
			var next *ListNode
			if l1 != nil || l2 != nil {
				next = &ListNode{}
			} else {
				next = nil
			}
			*tail = ListNode{n % 10, next}
		}
		tail = tail.Next
	}
	return head
}

//标准答案
func standardAnswer(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return

}

func main() {
	a := &ListNode{
		2, &ListNode{
			4, &ListNode{
				3, nil,
			},
		},
	}
	b := &ListNode{
		5, &ListNode{
			6, &ListNode{
				4, nil,
			},
		},
	}
	fmt.Println(*addTwoNumbers(a, b))
}
