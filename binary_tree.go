package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	Left  *node
	Right *node
	Data  int
}

// Tree ...
type Tree struct {
	Root *node
}

// FindNode ...
func (t *Tree) FindNode(data int) bool {
	newNode := node{
		Data: data,
	}
	if t.Root != nil {
		if t.findNode(t.Root, newNode) != nil {
			return true
		}
	}
	return false
}

func (t *Tree) findNode(search *node, target node) *node {
	var returnNode *node
	if search == nil {
		return returnNode
	}
	if search.Data == target.Data {
		return search
	}
	returnNode = t.findNode(search.Left, target)
	if returnNode == nil {
		returnNode = t.findNode(search.Right, target)
	}
	return returnNode
}

// Add ...
func (t *Tree) Add(data int) {
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
	t.add(t.Root, nodeToAdd)
	return
}

func (t *Tree) add(oldnode *node, newNode node) {
	if newNode.Data < oldnode.Data {
		if oldnode.Left == nil {
			oldnode.Left = &newNode
		} else {
			t.add(oldnode.Left, newNode)
		}
	} else if newNode.Data > oldnode.Data {
		if oldnode.Right == nil {
			oldnode.Right = &newNode
		} else {

			t.add(oldnode.Right, newNode)
		}
	}
	return
}

// InOrderTraversal ...
func (t *Tree) InOrderTraversal() {
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			fmt.Println(currentNode.Data)
		} else {
			t.inOrderTraversal(currentNode)
		}
	}
	return
}

func (t *Tree) inOrderTraversal(n *node) {
	if n.Left != nil {
		t.inOrderTraversal(n.Left)
	}
	fmt.Println(n.Data)
	if n.Right != nil {
		t.inOrderTraversal(n.Right)
	}
	return
}

// Traversal ...
func (t *Tree) Traversal() {
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			fmt.Println(currentNode.Data)
		} else {
			t.traversal(currentNode)
		}
	}
	return
}

func (t *Tree) traversal(n *node) {
	fmt.Println(n.Data)
	if n.Left != nil {
		t.traversal(n.Left)
	}

	if n.Right != nil {
		t.traversal(n.Right)
	}
	return
}

// CountEdges ...
func (t *Tree) CountEdges() (edges int) {
	c := make(chan int, 10)
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			return 1
		}
		t.countEdges(currentNode, c)
	}

	for {
		n := <-c
		if n == 0 {
			close(c)
			break
		}
		edges++

	}
	return edges
}

func (t *Tree) countEdges(n *node, counter chan int) {
	if n.Left != nil {
		go t.countEdges(n.Left, counter)
	}

	if n.Left == nil && n.Right == nil {
		counter <- 0
	} else {
		counter <- 1
	}

	if n.Right != nil {
		go t.countEdges(n.Right, counter)
	}
	return
}

// GenerateRandomTree ...
func (t *Tree) GenerateRandomTree() {
	u := time.Now()
	source := rand.NewSource(u.Unix())
	r := rand.New(source)
	arr := r.Perm(1000)
	for _, a := range arr {
		t.Add(a)
	}
	return
}

// PrintTree ...
func (t *Tree) PrintTree() {
	b, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func main() {
	t := Tree{}
	t.GenerateRandomTree()
	t.PrintTree()

}
