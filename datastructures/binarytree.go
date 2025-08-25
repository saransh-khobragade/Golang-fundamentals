// Binary Tree and Binary Search Tree implementation in Go
// Includes both simple binary tree and BST with traversal methods

package datastructures

import "fmt"

// Simple Binary Tree Node
type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func NewBinaryTreeNode(data int) *BinaryTreeNode {
	return &BinaryTreeNode{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (n *BinaryTreeNode) InsertLeft(data int) {
	if n.left == nil {
		n.left = NewBinaryTreeNode(data)
	} else {
		temp := NewBinaryTreeNode(data)
		temp.left = n.left
		n.left = temp
	}
}

func (n *BinaryTreeNode) InsertRight(data int) {
	if n.right == nil {
		n.right = NewBinaryTreeNode(data)
	} else {
		temp := NewBinaryTreeNode(data)
		temp.right = n.right
		n.right = temp
	}
}

func (n *BinaryTreeNode) GetRight() *BinaryTreeNode {
	return n.right
}

func (n *BinaryTreeNode) GetLeft() *BinaryTreeNode {
	return n.left
}

func (n *BinaryTreeNode) GetData() int {
	return n.data
}

// Binary Search Tree Node
type BSTNode struct {
	data  int
	left  *BSTNode
	right *BSTNode
}

func NewBSTNode(data int) *BSTNode {
	return &BSTNode{
		data:  data,
		left:  nil,
		right: nil,
	}
}

// Binary Search Tree Class
type BinarySearchTree struct {
	root *BSTNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

func (bst *BinarySearchTree) Insert(data int) {
	newNode := NewBSTNode(data)
	if bst.root == nil {
		bst.root = newNode
		return
	}
	bst.insertNode(bst.root, newNode)
}

func (bst *BinarySearchTree) insertNode(node, newNode *BSTNode) {
	if newNode.data < node.data {
		if node.left == nil {
			node.left = newNode
		} else {
			bst.insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			bst.insertNode(node.right, newNode)
		}
	}
}

func (bst *BinarySearchTree) Search(data int) *BSTNode {
	return bst.searchNode(bst.root, data)
}

func (bst *BinarySearchTree) searchNode(node *BSTNode, data int) *BSTNode {
	if node == nil {
		return nil
	}
	if data < node.data {
		return bst.searchNode(node.left, data)
	} else if data > node.data {
		return bst.searchNode(node.right, data)
	} else {
		return node
	}
}

func (bst *BinarySearchTree) Remove(data int) {
	bst.root = bst.removeNode(bst.root, data)
}

func (bst *BinarySearchTree) removeNode(node *BSTNode, data int) *BSTNode {
	if node == nil {
		return nil
	}
	if data < node.data {
		node.left = bst.removeNode(node.left, data)
	} else if data > node.data {
		node.right = bst.removeNode(node.right, data)
	} else {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			minNode := bst.findMin(node.right)
			node.data = minNode.data
			node.right = bst.removeNode(node.right, minNode.data)
		}
	}
	return node
}

func (bst *BinarySearchTree) findMin(node *BSTNode) *BSTNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

// BST Traversals
func (bst *BinarySearchTree) InOrderTraversal() []int {
	result := []int{}
	bst.inOrderTraversal(bst.root, &result)
	return result
}

func (bst *BinarySearchTree) inOrderTraversal(node *BSTNode, result *[]int) {
	if node != nil {
		bst.inOrderTraversal(node.left, result)
		*result = append(*result, node.data)
		bst.inOrderTraversal(node.right, result)
	}
}

func (bst *BinarySearchTree) PreOrderTraversal() []int {
	result := []int{}
	bst.preOrderTraversal(bst.root, &result)
	return result
}

func (bst *BinarySearchTree) preOrderTraversal(node *BSTNode, result *[]int) {
	if node != nil {
		*result = append(*result, node.data)
		bst.preOrderTraversal(node.left, result)
		bst.preOrderTraversal(node.right, result)
	}
}

func (bst *BinarySearchTree) PostOrderTraversal() []int {
	result := []int{}
	bst.postOrderTraversal(bst.root, &result)
	return result
}

func (bst *BinarySearchTree) postOrderTraversal(node *BSTNode, result *[]int) {
	if node != nil {
		bst.postOrderTraversal(node.left, result)
		bst.postOrderTraversal(node.right, result)
		*result = append(*result, node.data)
	}
}

// Simple Binary Tree Traversal Functions
func Preorder(tree *BinaryTreeNode) {
	if tree != nil {
		fmt.Println(tree.GetData())
		Preorder(tree.GetLeft())
		Preorder(tree.GetRight())
	}
}

func Inorder(tree *BinaryTreeNode) {
	if tree != nil {
		Inorder(tree.GetLeft())
		fmt.Println(tree.GetData())
		Inorder(tree.GetRight())
	}
}

func Postorder(tree *BinaryTreeNode) {
	if tree != nil {
		Postorder(tree.GetLeft())
		Postorder(tree.GetRight())
		fmt.Println(tree.GetData())
	}
}

// Example usage function
func ExampleBinaryTree() {
	// Test Binary Search Tree
	bst := NewBinarySearchTree()
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)

	fmt.Println("In-order traversal:", bst.InOrderTraversal())
	fmt.Println("Pre-order traversal:", bst.PreOrderTraversal())
	fmt.Println("Post-order traversal:", bst.PostOrderTraversal())

	// Test simple binary tree
	root := NewBinaryTreeNode(1)
	root.InsertLeft(2)
	root.InsertRight(3)
	root.GetLeft().InsertLeft(4)
	root.GetLeft().InsertRight(5)

	fmt.Println("\nSimple Binary Tree Pre-order:")
	Preorder(root)
} 