package basicStudy

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
	// 밑의 문법을 이용하면 구조체 이용 시 Key Field를 명시 해줘야 함
	// https://stackoverflow.com/questions/48381241/what-is-the-purpose-of-a-field-named-underscore-containing-an-empty-struct
	//_ struct{}
}

func Call() {
	//reserved()
	//types()
}

func reserved() {
	// PrintLn 은 fmt package 에서 export 된 함수이다
	// 이처럼 go 에서는 export 하기 위해선 함수 함수의 맨 앞자를 대문자로 정의 해야 export
	fmt.Println("Hello Golang")
	fmt.Println(Add(1, 2, 3))

	// double quotes
	const x = "ab\t" // const 는 := 사용 불가능

	// back quotes - 이스케이프 허용 안함
	y := `d\t` // var string y = "d"

	// single quotes - (byte(ascii) === uinit8) 또는 (rune(utf-8) === int32) 타입 선언 할 때
	var z byte
	z = 'a' // 97

	// defer 예약어, defer 예약어를 가진 함수가 종료 될 때까지 대기 하다 마지막에 호출됨, defer가 여러개일 경우 stack에 쌓임
	defer fmt.Println(Strings(x, y, z))
	fmt.Printf("x is of type %T \n", x)
	defer fmt.Println(Strings("a", "b", 'c')) // 스택으로 인해 먼저 호출

	//Gugudan(9)
	//fmt.Println(Sqrt(2))
}

func types() {
	var nullArr []string //nil slices
	// nullArr is nil
	// p is not nil
	p := &nullArr // &[]

	fmt.Println(IsNil(p))  // false
	fmt.Println(IsNil(*p)) // true

	// Array는 공간의 크기를 줄 수 있음
	// ex) var intArr [3]int
	// ex) intArr := [3]int {1,2,3}
	//		refIntArr := intArr[1:2] // Slices - Slices는 참조 전용이기 때문에 값을 저장할 수 없다.
	var intArr []int // Slices
	fmt.Println(append(intArr, 2))

	Slices()
}

func Method() {
	//v := Vertex{3, 4}
	//v.Scale(10)
	//fmt.Println(Abs(v))
	//
	//f := MyFloat(-math.Sqrt2)
	//fmt.Println(f.Abs())

	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8) // p 선언에 포인터로
	fmt.Println(v, p)

}

func Interface() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}
