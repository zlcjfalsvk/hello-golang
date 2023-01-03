package basicStudy

import "math"

type MyFloat float64

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 포인터 리시버로 메서드 선언
// 포인터 리시버를 사용하는 이유
//  1. 메서드가 리시버가 가리키는 값을 수정할 수 있음
//  2. 각각의 메서드 call에서의 value 복사 문제 방지
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
