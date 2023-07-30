package main

/*
	链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 83. 删除链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}

// 82. 删除排序链表中的重复元素
func deleteDuplicates2(head *ListNode) *ListNode {
	// 头节点可能被删除，需要dummy节点
	dummy := &ListNode{Next: head}
	head = dummy
	var tmp int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			tmp = head.Next.Val // 记录要删除的节点值
			for head.Next != nil && head.Next.Val == tmp {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next	// 注意，要写在else里面，想想下一次循环
		}
	}
	return dummy.Next
}
