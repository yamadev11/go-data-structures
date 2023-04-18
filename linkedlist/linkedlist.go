package linkedlist

import "fmt"

// create node struct
type node struct {
	val  interface{}
	next *node
}

type LinkedList struct {
	head *node
}

func NewLinkedList(values ...interface{}) *LinkedList {
	var head, curr *node
	for i, value := range values {
		if i == 0 {
			head = &node{val: value, next: nil}
			curr = head
			continue
		}
		curr.next = &node{val: value, next: nil}
		curr = curr.next
	}

	return &LinkedList{head: head}
}

func (list *LinkedList) AddNode(val interface{}) {
	if list.head == nil {
		list.head = &node{val: val, next: nil}
		return
	}

	curr := list.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &node{val: val, next: nil}
}

func (list *LinkedList) PrintList() {
	curr := list.head
	for curr != nil {
		fmt.Print(" ", curr.val)
		curr = curr.next
	}

	fmt.Println()
}

func (list *LinkedList) Reverse() {

	if list == nil || list.head == nil || list.head.next == nil {
		return
	}

	curr := list.head.next
	list.head.next = nil
	for curr != nil {
		next := curr.next
		curr.next = list.head
		list.head = curr
		curr = next
	}

	fmt.Println("Reversed Linked List Successfully")
}

func (list *LinkedList) DeleteNode(val interface{}) {

	if list == nil || list.head == nil {
		return
	}

	if list.head.val == val {
		fmt.Println("Deleting Node having value", val)
		list.head = list.head.next
	}

	curr := list.head
	for curr.next != nil {
		if curr.next.val == val {
			fmt.Println("Deleting Node having value", val)
			curr.next = curr.next.next
		}

		curr = curr.next
	}

}
