package main

import (
	"errors"
	"fmt"
)

// ErrorHandling (에러 처리)
// Go에서는 에러를 값으로 반환하고, 이를 직접 체크해서 처리한다.
func ErrorHandling() {
	// 기본 에러 처리
	fmt.Println("\n// 기본 에러 처리")
	// Go에서는 함수에서 에러를 리턴하고, 호출한 쪽에서 반드시 에러를 확인하고 처리해야 한다.
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error 발생:", err)
	} else {
		fmt.Println("결과:", result)
	}

	// errors.New와 fmt.Errorf 사용
	fmt.Println("\n// errors.New와 fmt.Errorf 사용")
	// errors.New는 단순한 에러 메시지를 생성할 때 사용
	// fmt.Errorf는 포맷팅된 에러 메시지 생성 및 기존 에러를 wrapping할 때 사용
	errSimple := errors.New("단순 에러")
	errFormatted := fmt.Errorf("상세 에러: %w", errSimple)
	fmt.Println(errFormatted)

	// error wrapping 확인 (errors.Is)
	fmt.Println("\n// error wrapping 확인")
	// fmt.Errorf로 감싼 에러라도, errors.Is를 사용하면 원본 에러를 비교할 수 있다.
	if errors.Is(errFormatted, errSimple) {
		fmt.Println("원본 에러와 매칭됨")
	}

	// panic 사용 예제
	fmt.Println("\n// panic 사용 예제")
	// panic은 복구하지 않으면 프로그램을 즉시 종료시킨다.
	// 심각한 오류가 발생했을 때 사용하는데, 일반 로직에서는 가급적 사용하지 않는다.
	//panic("치명적인 에러 발생")
	// 주석 처리: 강제 종료되므로 주의

	// panic + recover 조합 사용 예제
	fmt.Println("\n// panic + recover 조합 사용")
	// defer + recover 조합을 사용하면, panic 상황에서도 프로그램을 안전하게 복구할 수 있다.
	safeRun(func() {
		panic("고루틴 내 치명적 에러")
	})
	fmt.Println("panic 이후에도 프로그램이 계속 실행됨")
}

// divide 함수: 0으로 나누기를 체크하는 예제
// 에러를 반환하여 호출자가 직접 처리하도록 한다.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("0으로 나눌 수 없습니다")
	}
	return a / b, nil
}

// safeRun 함수: panic을 복구하는 패턴
// defer 안에서 recover()를 호출하여 panic을 잡고 프로그램이 중단되지 않게 한다.
type safeFunc func()

func safeRun(fn safeFunc) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("복구 성공:", r)
		}
	}()
	fn()
}

func main() {
	ErrorHandling()
}
