package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	if t != nil {
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(nil, tree.New(2)))
	fmt.Println(Same(nil, nil))
}
