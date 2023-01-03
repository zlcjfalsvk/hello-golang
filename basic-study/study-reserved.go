package basicStudy

import "fmt"

func Add(x, y, z int) int { // type이 같을 경우 param 마다의 type 선언을 생략할 수 있음 => func add(x int, y int, z int)
	return x + y + z
}

func Strings(x, y string, z byte) (string, string, float64) {
	return x, y, float64(z)
}

// Split
// naked return return value가 없지만 return type 에 변수가 정의 됨
// 짧은 함수에서 사용되어야 가독성이 좋음
func plit(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Gugudan(dan int) {
	// for 처럼 ; 이용 하여 여러가지 구문 가능
	if dan == 0 {
		dan = 9
	}

	//for i := 1; i <= dan; i++ {
	//	for j := 1; j <= 9; j++ {
	//		fmt.Printf("%d X %d = %d \n", i, j, i*j)
	//	}
	//}

	i := 1
	j := 1
	for ; i <= dan; i++ {
		// go 는 다음과 같이 ;를 생략할 수 있기 때문에 while 예약어 없음
		// for {} === while(true) {}
		for j <= 9 {
			fmt.Printf("%d X %d = %d \n", i, j, i*j)
			j++
		}
		j = 1
	}
}

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Sqrt approximation of %v attempt %v = %v\n", x, i, z)
	}
	return z
}
