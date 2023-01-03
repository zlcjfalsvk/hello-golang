package basicStudy

import (
	"fmt"
	"golang.org/x/tour/tree"
	"time"
)

func run() {
	// goroutine
	go say("world")

	s := []int{7, 2, 8, -9, 4, 0}

	// channel 생성
	// make(chan int, 10) -> 10 size 의 buffer
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	cc := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-cc)
		}
		quit <- 0
	}()
	goFibonacci(cc, quit)
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Channel '<-'
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c

	// channel 을 range 나 send 할 값이 없다는 것을 알아야 할 때만 주로 사용
	// !!! 수신자가 아닌 전송자만 channel을 닫아야 함
	//close(c)

}

func goFibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// goroutine이 다중 커뮤니케이션 연산에서 대기할 수 있게 함
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

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
