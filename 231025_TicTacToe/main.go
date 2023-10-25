package main

import "fmt"

type Cord struct {
	i, j int
}

func main() {
	n := 3
	c := []Cord{{0, 0}, {0, 2}, {1, 1}, {2, 2}}

	printBoard(c, n)
}

func printBoard(c []Cord, n int) {
	//Header
	fmt.Print("  ")
	for i := 0; i < n; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//Ausgabe der Koordinaten
	for i := 0; i < n; i++ {
		fmt.Print(i, " ")
		for j := 0; j < n; j++ {
			if contains(c, Cord{i, j}) {
				fmt.Print("X ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}

func contains(c []Cord, e Cord) bool {
	//Vergleicht die aktuelle Koordinate e mit den Koordinaten in der Liste
	for _, a := range c {
		if a == e {
			return true
		}
	}
	return false
}
