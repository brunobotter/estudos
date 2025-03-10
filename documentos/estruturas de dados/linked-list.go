package main

import "fmt"

type Node struct {
	Data string
	Next *Node
}

func main() {
	no := Node{Data: "teste", Next: &Node{Data: "novo", Next: &Node{Data: "faf", Next: nil}}}
	no.printList()
}

func (n *Node) printList() {
	current := n
	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}
	fmt.Print("null")
}
