package main

import "fmt"

// GenericsExamples (제네릭)
// Go 1.18부터 지원하는 기능으로, 타입을 매개변수화하여 다양한 타입에 대해 재사용 가능한 코드를 작성할 수 있다.
func GenericsExamples() {
	// 제네릭 함수 사용
	fmt.Println("\n// 제네릭 함수 사용")
	fmt.Println(Echo(10))      // int 타입
	fmt.Println(Echo("hello")) // string 타입

	// 제네릭 구조체 사용
	fmt.Println("\n// 제네릭 구조체 사용")
	intBox := Box[int]{value: 123}
	strBox := Box[string]{value: "hello"}
	fmt.Println(intBox.value)
	fmt.Println(strBox.value)

	// 타입 제약 사용
	fmt.Println("\n// 타입 제약 사용")
	fmt.Println(Max(10, 20))     // int 타입
	fmt.Println(Max(3.14, 2.71)) // float64 타입

	// 여러 타입 매개변수 사용
	fmt.Println("\n// 여러 타입 매개변수 사용")
	a, b := Pair(1, "one")
	fmt.Println(a, b)
}

// Echo 함수: 입력받은 값을 그대로 반환하는 제네릭 함수
// any는 모든 타입을 허용하는 특별한 타입 제약이다.
func Echo[T any](value T) T {
	return value
}

// Box 구조체: 하나의 값(value)을 저장하는 제네릭 구조체
type Box[T any] struct {
	value T
}

// Number 타입 제약: 비교 연산이 가능한 타입만 허용
// ~ 기호는 기본 타입을 포함하거나 확장하는 의미를 가진다.
type Number interface {
	~int | ~float64
}

// Max 함수: 두 값을 비교하여 큰 값을 반환하는 제네릭 함수
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Pair 함수: 두 타입 매개변수를 받아 튜플 형태로 반환하는 예제
func Pair[A, B any](a A, b B) (A, B) {
	return a, b
}

func main() {
	GenericsExamples()
}
