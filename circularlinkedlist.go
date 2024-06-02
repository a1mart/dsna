package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

type CircularLinkedList struct {
    Head *ListNode
}

func (list *CircularLinkedList) Insert(val int) {
    newNode := &ListNode{Val: val}
    if list.Head == nil {
        list.Head = newNode
        newNode.Next = list.Head
    } else {
        current := list.Head
        for current.Next != list.Head {
            current = current.Next
        }
        current.Next = newNode
        newNode.Next = list.Head
    }
}

func (list *CircularLinkedList) Display() {
    current := list.Head
    if current != nil {
        for {
            fmt.Printf("%d -> ", current.Val)
            current = current.Next
            if current == list.Head {
                break
            }
        }
    }
    fmt.Println("nil")
}

func main() {
    linkedList := CircularLinkedList{}
    linkedList.Insert(1)
    linkedList.Insert(2)
    linkedList.Insert(3)
    linkedList.Display()
}
