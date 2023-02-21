package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

//type Tree struct {
//	Left  *Tree
//	Value int
//	Right *Tree
//}

func Path(t *tree.Tree, ch chan int) {
	fmt.Println("\n", t.String())
	fmt.Println("Left", t.Left.String())
	if t.Left != nil {
		Path(t.Left, ch)
	}
	fmt.Println("Value", t.Value)
	ch <- t.Value
	fmt.Println("Right", t.Right.String())
	if t.Right != nil {
		Path(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Path(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		x, ok1 := <-ch1
		y, ok2 := <-ch2
		fmt.Println(x, y)
		if x != y {
			return false
		}
		if ok1 == false && ok2 == false {
			return true
		}
	}
}

func main() {
	fmt.Println("Result", Same(tree.New(1), tree.New(1)))
	fmt.Println("Result", Same(tree.New(1), tree.New(2)))
}
