package trie

type Node struct {
	children map[rune]*Node
	isWord   bool
}

func NewNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
		isWord:   false,
	}
}
