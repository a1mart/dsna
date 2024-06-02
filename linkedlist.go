package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

type SinglyLinkedList struct {
    Head *ListNode
}

func (list *SinglyLinkedList) Insert(val int) {
    newNode := &ListNode{Val: val}
    if list.Head == nil {
        list.Head = newNode
    } else {
        current := list.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
}

func (list *SinglyLinkedList) Display() {
    current := list.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Val)
        current = current.Next
    }
    fmt.Println("nil")
}

func main() {
    linkedList := SinglyLinkedList{}
    linkedList.Insert(1)
    linkedList.Insert(2)
    linkedList.Insert(3)
    linkedList.Display()
}
