package structures

import "fmt"

type LinkedListNode struct {
	data int
	prev *LinkedListNode
	next *LinkedListNode
}

func NewLinkedList() *LinkedListNode {
	return &LinkedListNode{}
}

func LinkedListInsert(node *LinkedListNode, data int) {
	newNode := &LinkedListNode{
		data: data,
		prev: nil,
		next: nil,
	}

	if node == nil {
		node = newNode
		return
	}

	current := node
	for current.next != nil {
		current = current.next
	}

	newNode.prev = current
	current.next = newNode

}

func InsertAfter(node *LinkedListNode, data int, newData int) {

	if node == nil {
		fmt.Println("Linked list is empty")
		return
	}

	newNode := &LinkedListNode{
		data: data,
		prev: nil,
		next: nil,
	}

	current := node

	for current.next != nil {
		if current.data == data {
			newNode.next = current.next
			if current.next != nil {
				current.next.prev = newNode
				newNode.prev = current
				return
			}
			current.next = newNode
			return
		}
		current = current.next
	}

}

func InsertBefore(node *LinkedListNode, data int, newData int) {

	if node == nil {
		fmt.Println("Linked list is empty")
		return
	}

	newNode := &LinkedListNode{
		data: data,
		prev: nil,
		next: nil,
	}

	current := node

	for current.prev != nil {
		if current.data == data {
			newNode.next = current
			if current.prev != nil {
				newNode.prev = current.prev
				current.prev.next = newNode
				return
			}
			current.prev = newNode
			return
		}
		current = current.prev
	}

}

func Remove(node *LinkedListNode, data int) {

	if node == nil {
		fmt.Println("Linked list is empty")
		return
	}

	current := node

	for current.next != nil {
		if current.data == data {
			if current.next != nil {
				if current.prev != nil {
					current.next.prev = current.prev
				}
				current.prev.next = current.next
			}
			current = nil
			return
		}
		current = current.next
	}
}

func Show(node *LinkedListNode) {
	current := node

	if current == nil {
		fmt.Println("Linked list is empty")
	}

	fmt.Println("Linked list:")
	for current.next != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}
