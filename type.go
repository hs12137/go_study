package main

import (
	"fmt"
)

// DeclarationOfVariable (변수 선언 방식)
// var 키워드로 타입 명시 가능 (var 변수명 타입 = 값)
// :=는 타입 추론을 사용하는 단축 선언
func DeclarationOfVariable() {
	fmt.Println("변수 선언 방식")
	var name string = "정희석"
	age := 28

	fmt.Println(name, age) // 정희석 28
}

// 기본 타입
// 정수: int, int8, int16, int32, int64
// 실수: float32, float64
// 문자: rune (유니코드 문자)
// 문자열: string
// 불리언: bool

// 특수 타입

// Array (배열)
// 고정된 크기
func Array() {
	fmt.Println("\n배열")
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr[0]) // 1
}

// Slice (슬라이스)
// 동적 배열처럼 크기 자동 조절 가능
func Slice() {
	fmt.Println("\n슬라이스")
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println(slice) // [1 2 3 4]

	// 슬라이싱
	arr := [5]int{10, 20, 30, 40, 50}
	s := arr[1:4]
	fmt.Println(s) // 20, 30, 40
	s = append(s, 80)
	fmt.Println(s) // 20, 30, 40, 80
}

// Map (맵)
// 키-값 쌍 저장. 해시맵과 유사함
func Map() {
	map1 := map[string]int{
		"사과":  100,
		"바나나": 200,
	}
	fmt.Println(map1["사과"]) // 100

	map1["딸기"] = 300
	fmt.Println(map1["딸기"]) // 300
}

//// Struct (구조체)
//// 필드를 가지는 사용자 정의 타입
//func Struct() {
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	person1 := Person{Name: "정희석", Age: 28}
//	fmt.Println(person1.Name) // 정희석
//	fmt.Println(person1.Age)  // 28
//}
//
//// Pointer (포인터)
//// 변수의 메모리 주소를 저장하는 타입
//func Pointer() {
//	num1 := 10
//	var pointer1 *int = &num1
//	fmt.Println(*pointer1) // 10 (역참조)
//}
//
//// Channel (채널)
//// 채널은 그 채널을 통하여 데이타를 주고 받는 통로라 볼 수 있는데, 채널은 make() 함수를 통해 미리 생성되어야 하며, 채널 연산자 <- 을 통해 데이타를 보내고 받는다
//// 채널은 흔히 goroutine들 사이 데이타를 주고 받는데 사용되는데, 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock을 걸지 않고 데이타를 동기화하는데 사용된다
//func Channel() {
//	channel1 := make(chan int)
//	go func() {
//		channel1 <- 42
//	}()
//	value := <-channel1
//	fmt.Println(value) // 42
//}
//
//// Interface (인터페이스)
//// 구조체(struct)가 필드들의 집합체라면, interface는 메서드들의 집합체이다
//// interface는 타입(type)이 구현해야 하는 메서드 원형(prototype)들을 정의한다
//// 하나의 사용자 정의 타입이 interface를 구현하기 위해서는 단순히 그 인터페이스가 갖는 모든 메서드들을 구현하면 된다
//func Interface() {
//	r := Rect{10., 20.}
//	c := Circle{10}
//
//	fmt.Println(r.area())      // 200
//	fmt.Println(r.perimeter()) // 60
//	fmt.Println(c.area())      // 314.1592653589793
//	fmt.Println(c.perimeter()) // 62.83185307179586
//}
//
//type Shape interface {
//	area() float64
//	perimeter() float64
//}
//
//// Rect 정의
//type Rect struct {
//	width, height float64
//}
//
//// Circle 정의
//type Circle struct {
//	radius float64
//}
//
//// Rect 타입에 대한 Shape 인터페이스 구현
//func (r Rect) area() float64 { return r.width * r.height }
//func (r Rect) perimeter() float64 {
//	return 2 * (r.width + r.height)
//}
//
//// Circle 타입에 대한 Shape 인터페이스 구현
//func (c Circle) area() float64 {
//	return math.Pi * c.radius * c.radius
//}
//func (c Circle) perimeter() float64 {
//	return 2 * math.Pi * c.radius
//}
//
//// Copy
//// copy(dst, src)의 형태로 사용
//// src의 내용을 dst에 복사 (최대 복사 가능한 길이만큼)
//// 두 슬라이스는 미리 선언되어 있어야 함
//func Copy() {
//	a := []int{1, 2, 3}
//	b := make([]int, len(a)) // b: [0 0 0]
//	copy(b, a)
//	fmt.Println(b) // [1 2 3]
//
//	// 슬라이스는 내부적으로 포인터 기반이라, a와 b가 같은 데이터를 가리킨다
//	// 단순 대입은 얕은 복사이므로 같은 위치를 공유한다
//	c := []int{1, 2, 3}
//	d := c
//	d[0] = 100
//
//	fmt.Println(c) // [100 2 3] — 원본도 바뀐다
//
//	// copy는 깊은 복사로 다른 주소를 카르킨다
//	e := []int{1, 2, 3}
//	f := make([]int, len(a))
//	copy(f, e)
//	f[0] = 100
//
//	fmt.Println(e) // [1 2 3] — 원본은 그대로
//	fmt.Println(f) // [100 2 3]
//
//}

func go_type() {
	DeclarationOfVariable()
	Array()
	Slice()
	Map()
	//Struct()
	//Pointer()
	//Channel()
	//Interface()
	//Copy()
}
