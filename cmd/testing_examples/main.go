package main

import (
	"fmt"
	"testing"
)

// TestingExamples (테스트)
// Go의 testing 패키지를 사용하여 단위 테스트, 벤치마크 테스트, 테이블 기반 테스트를 작성할 수 있다.
// 테스트 파일은 *_test.go로 끝나야 하고, 테스트 함수는 Test로 시작해야 한다.

// add 함수: 두 정수를 더하는 간단한 함수 (테스트 대상)
func add(a, b int) int {
	return a + b
}

// TestAdd: add 함수를 테스트하는 기본 단위 테스트
func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("add(2,3) = %d; want %d", result, expected)
	}
}

// BenchmarkAdd: add 함수를 벤치마크하는 테스트
// 벤치마크 함수는 Benchmark로 시작하고, *testing.B를 매개변수로 받는다.
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(2, 3)
	}
}

// TestAddTableDriven: 테이블 기반 테스트
// 다양한 입력과 기대값을 테이블로 만들어 반복 테스트한다.
func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"zero", 0, 0, 0},
		{"negative numbers", -1, -1, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("add(%d,%d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// 메인 프로그램 (일반 실행 용도)
func main() {
	fmt.Println("테스트는 `go test` 명령어로 실행합니다.")
}
