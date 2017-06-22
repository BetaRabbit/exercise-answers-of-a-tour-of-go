package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for n := range ch1 {
		if n != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("test Walk")
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for n := range ch {
		fmt.Printf("%v ", n)
	}

	fmt.Println("\n\n")

	fmt.Println("test Same")
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
