package main

import "fmt"

// ControlFlow (제어문)
// 조건문과 반복문을 통한 흐름 제어
func ControlFlow() {
	// if-else 문
	fmt.Println("\n// if-else 문")
	x := 15
	if x > 10 {
		fmt.Println("x는 10보다 크다")
	} else if x == 10 {
		fmt.Println("x는 10이다")
	} else {
		fmt.Println("x는 10보다 작다")
	}

	// if 문에서 변수 선언하기
	fmt.Println("\n// if 문에서 변수 선언")
	if y := 20; y > 10 {
		fmt.Println("y는 10보다 크다")
	}

	// switch 문
	fmt.Println("\n// switch 문")
	switch x {
	case 5:
		fmt.Println("x는 5")
	case 10, 15:
		fmt.Println("x는 10 또는 15")
	default:
		fmt.Println("x는 5, 10, 15가 아니다")
	}

	// for 문 (일반 반복)
	fmt.Println("\n// for 문 (일반 반복)")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for 문 (조건만)
	fmt.Println("\n// for 문 (조건만)")
	n := 3
	for n > 0 {
		fmt.Println(n)
		n--
	}

	// for 문 (무한 루프)
	fmt.Println("\n// for 문 (무한 루프 예시)")
	count := 0
	for {
		fmt.Println("count:", count)
		count++
		if count == 3 {
			break
		}
	}

	// for-range 문 (슬라이스 순회)
	fmt.Println("\n// for-range 문 (슬라이스 순회)")
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	// for-range 문 (맵 순회)
	fmt.Println("\n// for-range 문 (맵 순회)")
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}

	// for-range 문 (문자열 순회)
	fmt.Println("\n// for-range 문 (문자열 순회)")
	str := "GoLang"
	for i, r := range str {
		fmt.Printf("index: %d, rune: %c\n", i, r)
	}

	// 슬라이스 수정시 주의점 (값 복사본 수정은 원본에 영향 없음)
	fmt.Println("\n// 슬라이스 수정시 주의점")
	nums2 := []int{1, 2, 3}
	for _, v := range nums2 {
		v = v * 2
	}
	fmt.Println(nums2) // [1 2 3] - 수정되지 않음

	// 인덱스로 직접 수정해야 함
	for i := range nums2 {
		nums2[i] = nums2[i] * 2
	}
	fmt.Println(nums2) // [2 4 6]

	// copy 함수 사용 예시
	fmt.Println("\n// copy 함수 사용")
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)

	dst[0] = 100
	fmt.Println("수정 후 src:", src)
	fmt.Println("수정 후 dst:", dst)
}

func main() {
	ControlFlow()
}
