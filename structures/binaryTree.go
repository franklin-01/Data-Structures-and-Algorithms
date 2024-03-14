package structures

import "fmt"

type BinaryNode struct {
	Data  *int
	Left  *BinaryNode
	Right *BinaryNode
}

func NewBinaryTree() *BinaryNode {
	return &BinaryNode{
		Data:  nil,
		Left:  nil,
		Right: nil,
	}
}

func Insert(node *BinaryNode, data int) {

	if node == nil {
		node = &BinaryNode{
			Data:  &data,
			Left:  nil,
			Right: nil,
		}
	}

	node.insert(data)

}

func (node *BinaryNode) insert(data int) {
	newNode := &BinaryNode{
		Data:  &data,
		Left:  nil,
		Right: nil,
	}

	if data < *node.Data {
		if node.Left == nil {
			node.Left = newNode
			return
		}
		node.Left.insert(data)
		return
	}

	if node.Right == nil {
		node.Right = newNode
		return
	}

	node.Right.insert(data)
}

func printPreOrder(node *BinaryNode) {
	if node == nil {
		return
	} else {
		fmt.Printf("%d ", node.Data)
		printPreOrder(node.Left)
		printPreOrder(node.Right)
	}
}

func printPostOrder(node *BinaryNode) {
	if node == nil {
		return
	} else {
		printPreOrder(node.Left)
		printPreOrder(node.Right)
		fmt.Printf("%d ", node.Data)
	}
}

func printInOrder(node *BinaryNode) {
	if node == nil {
		return
	} else {
		printPreOrder(node.Left)
		fmt.Printf("%d ", node.Data)
		printPreOrder(node.Right)
	}
}
