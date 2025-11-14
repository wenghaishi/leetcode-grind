package main


type ListNode struct {
    Val int
    Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
    dummy := &ListNode{} 
    newListEnd := dummy
    current := head.Next
    sum := 0
    for current != nil {
        if current.Val == 0 {
            newNode := &ListNode{}
            newNode.Val = sum
            newListEnd.Next = newNode
            newListEnd = newNode
            sum = 0
        } else {
            sum += current.Val
        }
        current = current.Next
    }

    return dummy.Next
}