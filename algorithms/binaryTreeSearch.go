package algorithms

import "main/structures"

func BinaryTreeSearch(node *structures.BinaryNode, value *int) *structures.BinaryNode {
	if node == nil {
		return nil
	}

	if *node.Data == *value {
		return node
	}

	if *value < *node.Data {
		return BinaryTreeSearch(node.Left, value)
	}

	return BinaryTreeSearch(node.Right, value)
}
