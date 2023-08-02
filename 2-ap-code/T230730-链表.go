package main

/*
	链表
	链表相关的核心点
		null/nil 异常处理
		dummy node 哑巴节点: 头节点不确定的时候使用
		快慢指针
		插入一个节点到排序链表
		从一个链表中移除一个节点
		翻转链表
		合并两个链表
		找到链表的中间节点

题目：
	83. 删除链表中的重复元素(保留1个）
	82. 删除排序链表中的重复元素
	206. 反转链表
	92. 翻转链表2（反转中间部分）
	21. 合并两个有序链表
	86. 分隔链表
	148. 链表排序
	143. 重排链表
	141. 环形链表
	142. 环形链表2

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
			head = head.Next // 注意，要写在else里面，想想下一次循环
		}
	}
	return dummy.Next
}

// 206.反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

// 92.翻转链表2
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 头节点可能被删，所以需要dummy节点
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head} // 头节点之前的哑节点
	head = dummy
	pre := &ListNode{} // 要反转链表的前一个节点
	i := 0
	for i < left {
		pre = head
		head = head.Next
		i++
	}

	j := i
	next := &ListNode{} // 用来反转链表时使用
	tail := head        // 要反转链表的最后一个节点
	for head != nil && j <= right {
		tmp := head.Next
		head.Next = next
		next = head
		head = tmp
		j++
	}
	pre.Next = next
	tail.Next = head
	return dummy.Next
}

// 21. 合并两个有序链表
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil {
			head.Next = l2
			return dummy.Next
		}
		if l2 == nil {
			head.Next = l1
			return dummy.Next
		}
		if l1.Val <= l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	return dummy.Next
}

// 86. 分隔链表
func partitionList(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	headDummy := &ListNode{} // 前半部分的头节点
	tailDummy := &ListNode{} // 后半部分的头节点
	tail := tailDummy
	headDummy.Next = head
	head = headDummy

	for head.Next != nil {
		if head.Next.Val < x { // 小于x就下一位
			head = head.Next
		} else {
			tmp := head.Next
			head.Next = head.Next.Next
			tail.Next = tmp
			tail = tail.Next
		}
	}
	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}

// 148. 链表排序--归并排序
func sortList(head *ListNode) *ListNode {
	return mergeSortList(head)
}
func mergeSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 1. 返回条件
		return head
	}
	mid := findMiddle(head) // 2. 分段处理
	tail := mid.Next
	mid.Next = nil
	left := mergeSortList(head)
	right := mergeSortList(tail)

	result := mergeTwoLists(left, right) //3.合并结果
	return result
}
func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	} // 快慢指针找中点
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 143. 重排链表---=快慢指针找中点+翻转链表+合并链表
func reorderList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	mid := findMiddle(head)
	tail := reverseList(mid.Next)
	mid.Next = nil
	result := mergeTwoLists2(head, tail)
	return result
}
func mergeTwoLists2(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	flag := true
	for l1 != nil || l2 != nil {
		if l1 == nil {
			head.Next = l2
			return dummy.Next
		}
		if l2 == nil {
			head.Next = l1
			return dummy.Next
		}
		if flag {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
		flag = !flag
	}
	return dummy.Next
}

// 141. 环形链表-快慢指针，有环肯定会相遇
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// 142. 环形链表2--返回入口。快慢指针，第一次相遇后，慢指针重置到头节点，第二次相遇就是结果
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			fast = fast.Next
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}
