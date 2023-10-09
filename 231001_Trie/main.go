package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/maheis/CodingCircle/231001_/trie"
)

// 01.10.2023 - https://www.youtube.com/watch?v=wklh-1YTzmM

func main() {
	t := trie.New()

	fmt.Println("Load List from file")
	file, err := os.Open("wortliste.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		t.Add(word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	fmt.Println("List loaded")
	start := time.Now()
	list := t.ListPossibleWords("Text", 135)
	end := time.Now()

	for word := range list {
		fmt.Println(word)
	}

	fmt.Println("Searchtime needed:", end.Sub(start), ", Found Words:", len(list))
}
