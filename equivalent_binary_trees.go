package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		close(ch)
		return
	}
	queue := make([]*tree.Tree, 0)
	queue = append(queue, t)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		ch <- node.Value
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	close(ch)

}

func equal(s1 []int, s2 []int) bool {
	if s1 == nil && s2 == nil {
		return true
	}
	if s1 == nil || s2 == nil {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
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

	s1 := make([]int, 10)
	s2 := make([]int, 10)
	for item1 := range ch1 {
		s1 = append(s1, item1)
	}
	for item1 := range ch2 {
		s2 = append(s2, item1)
	}

	sort.Ints(s1)
	sort.Ints(s2)
	return equal(s1, s2)
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(nil, tree.New(2)))
	fmt.Println(Same(nil, nil))
}
