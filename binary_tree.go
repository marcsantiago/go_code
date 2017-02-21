package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	Left  *node
	Right *node
	Data  int
}

type tree struct {
	Root *node
}

func (t *tree) findRoot(data int) bool {
	newNode := node{
		Data: data,
	}
	if t.Root != nil {
		if t.findRootLogic(t.Root, newNode) != nil {
			return true
		}
	}
	return false
}

func (t *tree) findRootLogic(search *node, target node) *node {
	var returnNode *node
	if search == nil {
		return returnNode
	}
	if search.Data == target.Data {
		return search
	}

	returnNode = t.findRootLogic(search.Left, target)
	if returnNode == nil {
		returnNode = t.findRootLogic(search.Right, target)
	}
	return returnNode

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
	// BASED ON http://rosalind.info/problems/tree/
	// NOT CORRECT YET
	t := new(tree)
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	s := string(b)
	parts := strings.Split(s, "\n")
	n := parts[0]

	startNode, _ := strconv.Atoi(n)
	t.add(startNode)

	for {
		counter := 1
		oldSize := len(parts) - 1
		newSize := oldSize
		for _, part := range parts[1:] {
			nums := strings.Split(part, " ")
			if len(nums) == 2 {
				toAdd, _ := strconv.Atoi(nums[0])
				root, _ := strconv.Atoi(nums[1])
				if t.findRoot(root) {
					t.add(toAdd)
					parts = parts[:counter+copy(parts[counter:], parts[counter+1:])]
				} else if t.findRoot(toAdd) {
					t.add(root)
					parts = parts[:counter+copy(parts[counter:], parts[counter+1:])]
				}

			}
			counter++
		}
		oldSize = len(parts) - 1
		if newSize == oldSize {
			break
		}

	}

	b, _ = json.MarshalIndent(t, "", " ")
	fmt.Println(string(b))
	fmt.Println(t.countEdges())

}
