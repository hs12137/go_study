package main

import "fmt"

// 제어문

// if-else
func if_else() {
	x := 15
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is 10 or less")
	} // x is greater than 10
}

// Switch
func Switch() {
	x := 5
	switch x {
	case 1:
		fmt.Println("One")
	case 2, 3:
		fmt.Println("Two or Three")
	default:
		fmt.Println("Other")
	} // Other
}

// For
func For() {
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	} // 01234
	fmt.Println()

	// 배열이나 슬라이스 순회
	// i: 인덱스
	// v: 해당 인덱스의 값
	nums := []int{10, 20, 30}

	for i, v := range nums {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
	// index: 0, value: 10
	// index: 1, value: 20
	// index: 2, value: 30

	// 값만 사용하고 싶을 때
	for _, v := range nums {
		fmt.Print(v, " ")
	} // 10 20 30
	fmt.Println()

	// 인덱스만 필요할 때
	for i := range nums {
		fmt.Print(i, " ")
	} // 0 1 2
	fmt.Println()

	// map 순회
	m := map[string]int{"사과": 100, "바나나": 200}
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	// key: 사과, value: 100
	// key: 바나나, value: 200

	//문자열 순회
	str := "안녕Go!"
	for i, r := range str {
		fmt.Printf("index: %d, rune: %c\n", i, r)
	}
	// index: 0, rune: 안
	// index: 3, rune: 녕
	// index: 6, rune: G
	// index: 7, rune: o
	// index: 8, rune: !

	// 채널 순회
	// range로 채널을 닫힐 때까지 읽을 수 있음
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch) // 반드시 닫아야 range가 종료됨
	}()

	for v := range ch {
		fmt.Println("Received:", v)
	}
	// Received: 1
	// Received: 2
	// Received: 3

	// 자료형 타입을 순회하면서 v는 복사본이기 때문에 수정해도 원본에 영향 x
	//
}

func controlStatement() {
	if_else()
	Switch()
	For()
}
