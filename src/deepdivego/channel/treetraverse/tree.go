package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

func (node *treeNode) setValue(value int) {
	// :ATTENTION: 使用指针才能改变内部值
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func (node *treeNode) traverseWithFunc(f func(*treeNode)) {
	if node == nil {
		return
	}
	node.left.traverseWithFunc(f)
	f(node)
	node.right.traverseWithFunc(f)
}

func (node *treeNode) traverseWithChannel() chan *treeNode {
	out := make(chan *treeNode)
	go func() {
		node.traverseWithFunc(func(node *treeNode) {
			out <- node
		})
		close(out)
	}()
	return out
}

func main() {
	root := treeNode{value: 3}
	root.left = &treeNode{value: 5}
	root.right = &treeNode{6, nil, nil}
	root.right.left = new(treeNode)
	root.right.right = &treeNode{value: 7}

	// root.print()
	// root.setValue(4)
	// fmt.Println()
	// root.traverse()
	c := root.traverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.value > maxNode {
			maxNode = node.value
		}
	}
	fmt.Println("Max node is: ", maxNode)
}
