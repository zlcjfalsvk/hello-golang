package basicStudy

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// 포인트 리시버로 메서드 선언
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeForany(i any) { // any === interface {}
	fmt.Printf("(%v, %T)\n", i, i)
}
