package main

import (
	"fmt"

	t "golang.org/x/tour/tree"
)

func Walk(t *t.Tree, ch chan int) {
	recWalk(t, ch)
	close(ch)
}

func recWalk(t *t.Tree, ch chan int) {
	if t != nil {
		recWalk(t.Left, ch)
		ch <- t.Value
		recWalk(t.Right, ch)
	}
}

func Same(t1, t2 *t.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2

		switch {
		// not same length
		case ok1 != ok2:
			return false
			// not
		case x1 != x2:
			return false
		case !ok1:
			return true
		default:
			// do nothing
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(t.New(1), ch)
	fmt.Println("Walking the tree:")
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(Same(t.New(1), t.New(1)))
	fmt.Println(Same(t.New(1), t.New(2)))
}
