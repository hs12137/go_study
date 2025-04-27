package main

import "fmt"

// Pointer (포인터)
// 변수의 메모리 주소를 저장하는 타입
// Go에서는 포인터 연산(++, --, +, -)은 지원하지 않고 안전하게 주소 참조/역참조만 제공한다.
func Pointer() {
	// 기본 포인터 사용
	fmt.Println("\n// 기본 포인터 사용")
	num := 10
	p := &num       // num의 메모리 주소 저장
	fmt.Println(*p) // 10 (역참조해서 값 가져오기)

	// 함수에 포인터 전달
	fmt.Println("\n// 함수에 포인터 전달")
	a := 5
	fmt.Println("before:", a)
	addOne(&a) // 주소를 넘겨서 함수 안에서 수정
	fmt.Println("after:", a)

	// nil 포인터
	fmt.Println("\n// nil 포인터")
	var ptr *int
	fmt.Println(ptr) // <nil>
	if ptr == nil {
		fmt.Println("포인터가 nil입니다.")
	}

	// new 키워드로 포인터 생성
	fmt.Println("\n// new 키워드 사용")
	p2 := new(int) // int 타입 0으로 초기화된 공간 할당
	*p2 = 100
	fmt.Println(*p2) // 100

	// 포인터 배열
	fmt.Println("\n// 포인터 배열")
	x, y := 1, 2
	arr := [2]*int{&x, &y}
	fmt.Println(*arr[0], *arr[1]) // 1 2

	// 배열 포인터
	fmt.Println("\n// 배열 포인터")
	arr2 := [3]int{10, 20, 30}
	pArr := &arr2
	fmt.Println((*pArr)[0]) // 10
}

// addOne 함수: 포인터로 값을 수정
func addOne(n *int) {
	*n = *n + 1
}

func main() {
	Pointer()
}
