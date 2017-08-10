package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    l1 := &ListNode {
        Val: 1,
        Next: &ListNode {
            Val: 9,
            Next: &ListNode {
                Val: 5,
                Next: nil,
            },
        },
    }
    l2 := &ListNode {
        Val: 1,
        Next: &ListNode {
            Val: 1,
            Next: nil,
        },
    }
    l3 := addTwoNumbers(l1, l2)
    fmt.Printf("%d %d %d\n", l3.Val, l3.Next.Val, l3.Next.Next.Val)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head, node *ListNode
    ln1 := l1
    ln2 := l2
    carry := false
    for ln1 != nil || ln2 != nil {
        tmp := new(ListNode)
        val := 0
        if ln1 != nil {
            val += ln1.Val
        }
        if ln2 != nil {
            val += ln2.Val
        }
        if carry {
            val += 1
        }
        if val >= 10 {
            val -= 10
            carry = true
        } else {
            carry = false
        }
        tmp.Val = val
        if head == nil {
            head = tmp
            node = tmp
        } else {
            node.Next = tmp
            node = node.Next
        }
        if ln1 != nil {
            ln1 = ln1.Next
        }
        if ln2 != nil {
            ln2 = ln2.Next
        }
    }

    if carry {
        tmp := new(ListNode)
        tmp.Val = 1
        node.Next = tmp
    }

    return head
}
