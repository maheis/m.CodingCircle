package trie

import "fmt"

type Trie struct {
	root *Node
}

func New() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func (t *Trie) Add(word string) {
	node := t.root

	//Itteriert Buchstabenweise durch das Wort. i, der Index ist dabei nicht wichtig und wird mit _ ignoriert.
	for _, r := range word {
		//Prüft ob es den Buchstaben schon gibt, wenn nicht wird er hinzugefügt.
		if _, ok := node.children[r]; !ok {
			node.children[r] = NewNode()
		}
		node = node.children[r]
	}
	node.isWord = true
}

func (t *Trie) ListPossibleWords(prefix string, length int) map[string]int {
	node := t.root
	words := make(map[string]int, length)

	for _, r := range prefix {
		if _, ok := node.children[r]; !ok {
			return words
		}
		node = node.children[r]
	}

	t.listPossibleWords(node, prefix, words, length)

	return words
}

func (t *Trie) listPossibleWords(node *Node, prefix string, words map[string]int, length int) {
	if len(words) >= length {
		return
	}

	if node.isWord {
		words[prefix] = 0
	}

	for char, node := range node.children {
		t.listPossibleWords(node, prefix+string(char), words, length)
	}
}

func (t *Trie) Print() {
	fmt.Println("Trie:")
	t.print(t.root, 0)
}

func (t *Trie) print(node *Node, deep int) {
	for char, node := range node.children {
		for i := 0; i < deep; i++ {
			fmt.Print("  ")
		}

		fmt.Println(string(char))
		t.print(node, deep+1)
	}
}
