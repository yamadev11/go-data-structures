package linkedlist

import "fmt"

// create node struct
type node struct {
	val  interface{}
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func NewLinkedList(values ...interface{}) *LinkedList {
	var head, curr *node
	var length int
	for i, value := range values {
		if i == 0 {
			head = &node{val: value, next: nil}
			curr = head
			length = 1
			continue
		}
		curr.next = &node{val: value, next: nil}
		curr = curr.next
		length++
	}

	return &LinkedList{head: head, length: length}
}

func (list *LinkedList) AddNode(val interface{}) {
	if list.head == nil {
		list.head = &node{val: val, next: nil}
		list.length = 1
		return
	}

	curr := list.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &node{val: val, next: nil}
	list.length += 1
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
		return
	}

	curr := list.head
	for curr.next != nil {
		if curr.next.val == val {
			fmt.Println("Deleting Node having value", val)
			curr.next = curr.next.next
			return
		}

		curr = curr.next
	}

}

func (list *LinkedList) SwapNode(index1, index2 int) {
	if index1 > list.length || index2 > list.length || index1 < 0 || index2 < 0 {
		// log error
		return
	}

	var node1Prev, node2Prev, node1Next, node2Next *node
	node1 := list.head
	for i := 1; i <= index1; i++ {
		node1Prev = node1
		node1 = node1.next
		node1Next = node1.next
	}

	node2 := list.head
	for i := 1; i <= index2; i++ {
		node2Prev = node2
		node2 = node2.next
		node2Next = node2.next
	}

	node2Prev.next = node1
	node1.next = node2Next
	node1Prev.next = node2
	node2.next = node1Next
}

func (list *LinkedList) Length() int {
	if list != nil {
		return list.length
	}

	return 0
}