package main

/*
输入两个链表，找出它们的第一个公共节点。

双指针

每步操作需要同时更新指针 pA 和 pB。
如果指针 pA 不为空，则将指针 pA 移到下一个节点；如果指针 pB 不为空，则将指针 pB 移到下一个节点。
如果指针 pA 为空，则将指针 pA 移到链表 headB 的头节点；如果指针 pB 为空，则将指针 pB 移到链表 headA 的头节点。
当指针 pA 和 pB 指向同一个节点或者都为空时，返回它们指向的节点或者 null。

证明：
情况一：两个链表相交
假设链表 headA 的不相交部分有 aa 个节点，链表 headB 的不相交部分有 bb 个节点，两个链表相交的部分有 cc 个节点，则有 a+c=m，b+c=n。
如果 a=b，则两个指针会同时到达两个链表的第一个公共节点，此时返回两个链表的第一个公共节点；
如果 a！=b，则指针 pA 会遍历完链表 headA，指针 pB 会遍历完链表 headB，两个指针不会同时到达链表的尾节点，然后指针 pA 移到链表 headB 的头节点，
	指针 pB 移到链表 headA 的头节点，然后两个指针继续移动，在指针 pA 移动了 a+c+ba+c+b 次、指针 pB 移动了 b+c+ab+c+a 次之后，
	两个指针会同时到达两个链表的第一个公共节点，该节点也是两个指针第一次同时指向的节点，此时返回两个链表的第一个公共节点。
情况二：两个链表不相交
链表 headA 和 headB 的长度分别是 m 和 n。考虑当 m=n 和 m！=n 时，两个指针分别会如何移动：
如果 m=n，则两个指针会同时到达两个链表的尾节点，然后同时变成空值 null，此时返回 null；
如果 m！=n，则由于两个链表没有公共节点，两个指针也不会同时到达两个链表的尾节点，因此两个指针都会遍历完两个链表，
	在指针 pA 移动了 m+nm+n 次、指针 pB 移动了 n+mn+m 次之后，两个指针会同时变成空值 null，此时返回 null。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
