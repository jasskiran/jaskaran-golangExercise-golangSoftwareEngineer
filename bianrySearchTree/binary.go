package BST

import "fmt"

type Tree struct {
	Root *Node
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func BinarySearchTree(arr []int) {
	var t Tree

	fmt.Printf("Inserting data into binary tree: ")
	for _, v := range arr {
		t.Insert(v)
	}

	fmt.Printf("In Order: ")
	PrintInOrder(t.Root)
	fmt.Println()

	fmt.Printf("Pre Order: ")
	PrintPreOrder(t.Root)
	fmt.Println()

	fmt.Printf("Post Order: ")
	PrintPostOrder(t.Root)
	fmt.Println()
}

func (t *Tree) Insert(data int) {
	if t.Root == nil {
		t.Root = &Node{Key: data}
	} else {
		t.Root.Insert(data)
	}
}

func (n *Node) Insert(data int) {
	if data <= n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: data}
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Key: data}
		} else {
			n.Right.Insert(data)
		}
	}
}

func PrintInOrder(n *Node) {
	if n == nil {
		return
	} else {
		PrintInOrder(n.Left)
		fmt.Printf("%d ", n.Key)
		PrintInOrder(n.Right)
	}
}

func PrintPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.Key)
		PrintPreOrder(n.Left)
		PrintPreOrder(n.Right)
	}
}

func PrintPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		PrintPostOrder(n.Left)
		PrintPostOrder(n.Right)
		fmt.Printf("%d ", n.Key)
	}
}
