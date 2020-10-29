/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    h := &ListNode{
        Next: head,
    }
    fast, slow := head, h
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    for ; fast != nil; fast = fast.Next {
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return h.Next
}