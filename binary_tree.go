package main

import (
	"fmt"
)

type node struct {
	Left  *node
	Right *node
	Data  int
}

type tree struct {
	Root *node
}

func (t *tree) rosalindAdd(data int) {
	nodeToAdd := node{
		Data: data,
	}
	if t.Root == nil {
		t.Root = new(node)
	}
	if t.Root.Data == 0 {
		t.Root = &nodeToAdd
		return
	}
	t.rosalindTranverseAdd(t.Root, nodeToAdd)
	return
}

func (t *tree) rosalindTranverseAdd(oldnode *node, newNode node) {
	if newNode.Data < oldnode.Data {
		if oldnode.Left == nil {
			oldnode.Left = &newNode
		} else {
			t.tranverseAdd(oldnode.Left, newNode)
		}
	} else if newNode.Data > oldnode.Data {
		if oldnode.Right == nil {
			oldnode.Right = &newNode
		} else {

			t.rosalindTranverseAdd(oldnode.Right, newNode)
		}
	}
	return
}

func (t *tree) add(data int) {
	nodeToAdd := node{
		Data: data,
	}
	if t.Root == nil {
		t.Root = new(node)
	}
	if t.Root.Data == 0 {
		t.Root = &nodeToAdd
		return
	}
	t.tranverseAdd(t.Root, nodeToAdd)
	return
}

func (t *tree) tranverseAdd(oldnode *node, newNode node) {
	if newNode.Data < oldnode.Data {
		if oldnode.Left == nil {
			oldnode.Left = &newNode
		} else {
			t.tranverseAdd(oldnode.Left, newNode)
		}
	} else if newNode.Data > oldnode.Data {
		if oldnode.Right == nil {
			oldnode.Right = &newNode
		} else {

			t.tranverseAdd(oldnode.Right, newNode)
		}
	}
	return
}

func (t *tree) inOrderTraversal() {
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			fmt.Println(currentNode.Data)
		} else {
			t.inOrderTraversalLogic(currentNode)
		}
	}
	return
}

func (t *tree) inOrderTraversalLogic(n *node) {
	if n.Left != nil {
		t.inOrderTraversalLogic(n.Left)
	}
	fmt.Println(n.Data)
	if n.Right != nil {
		t.inOrderTraversalLogic(n.Right)
	}
	return
}

func (t *tree) countEdges() (edges int) {
	c := make(chan int, 10)
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			return 1
		}
		t.countEdgesLogic(currentNode, c)
	}

	for {
		n := <-c
		if n == 0 {
			close(c)
			break
		}
		edges++

	}
	return edges + 1 // connecting the root as 1 edge
}

func (t *tree) countEdgesLogic(n *node, counter chan int) {
	if n.Left != nil {
		go t.countEdgesLogic(n.Left, counter)
	}

	if n.Left == nil && n.Right == nil {
		counter <- 0
	} else {
		counter <- 1
	}

	if n.Right != nil {
		go t.countEdgesLogic(n.Right, counter)
	}
	return
}

func main() {
	t := new(tree)

	t.add(10)
	t.add(4)
	t.add(6)
	// t.add(9)
	// t.add(15)
	// t.add(20)
	// t.add(100)
	// t.add(8)
	// t.add(90)

	// b, _ := json.MarshalIndent(t, "", " ")
	// fmt.Println(string(b))
	n := t.countEdges()
	fmt.Println(n)

}
