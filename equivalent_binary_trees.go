package main

import "golang.org/x/tour/tree"
import "fmt"
import "sort"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	go Walk(t1, ch1)
	ch2 := make(chan int, 10)
	go Walk(t2, ch2)

	s1 := make(sort.IntSlice, 10)
	for i := 0; i < 10; i++ {
		s1[i] = <-ch1
	}
	close(ch1)

	s2 := make(sort.IntSlice, 10)
	for i := 0; i < 10; i++ {
		s2[i] = <-ch2
	}
	close(ch2)

	s1.Sort()
	s2.Sort()

	for i := 0; i < 10; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func main() {
	res := Same(tree.New(1), tree.New(1))
	fmt.Println(res)
}
