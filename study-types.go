package main

import (
	"fmt"
	"math"
	"reflect"
)

func IsNil[T any](i T) bool {
	return reflect.ValueOf(i).IsNil()
}

func Slices() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// range arr => [index, value]
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow { // for i, _ := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}

func Map() {
	type Vertex struct {
		Lat, Long float64
		// 밑의 문법을 이용하면 구조체 이용 시 Key Field를 명시 해줘야 함
		// https://stackoverflow.com/questions/48381241/what-is-the-purpose-of-a-field-named-underscore-containing-an-empty-struct
		//_ struct{}
	}

	var m map[string]Vertex // nil

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	v, isExist := m["Bell Labs"]
	fmt.Println(v, isExist)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func Fn() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// Closure 반환
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
