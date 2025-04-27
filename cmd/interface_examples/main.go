package main

import (
	"bytes"
	"fmt"
)

// Speaker interface 기본 예시
// Speak() 메서드를 구현하는 타입이면 자동으로 인터페이스를 만족
type Speaker interface {
	Speak() string
}

// Dog, Cat 타입이 Speaker 인터페이스 구현
// 명시적 implements 키워드 없이 자동으로 구현됨

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

// Person 타입: 포인터 리시버로 Speak() 구현
// 포인터 타입으로만 인터페이스 구현 가능 주의

type Person struct {
	Name string
}

func (p *Person) Speak() string { return "Hello, " + p.Name }

// Reader, Writer, ReadWriter 인터페이스 조합 예시
// Reader + Writer를 모두 만족해야 ReadWriter 구현

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}

// Interfaces 함수: 인터페이스 관련 예제 모음
func Interfaces() {
	// 기본 인터페이스 사용
	fmt.Println("\n// 기본 인터페이스 사용")
	var s Speaker = Dog{}
	fmt.Println(s.Speak())

	s = Cat{}
	fmt.Println(s.Speak())

	s = &Person{Name: "Alice"}
	fmt.Println(s.Speak())

	// 빈 인터페이스 사용
	fmt.Println("\n// 빈 인터페이스 사용")
	var i interface{}
	i = 123
	fmt.Println(i)
	i = "hello"
	fmt.Println(i)

	// 타입 단언 사용
	fmt.Println("\n// 타입 단언 사용")
	str, ok := i.(string)
	if ok {
		fmt.Println("string:", str)
	}
	num, ok := i.(int)
	if ok {
		fmt.Println("int:", num)
	} else {
		fmt.Println("int 단언 실패")
	}

	// 타입 스위치 사용
	fmt.Println("\n// 타입 스위치 사용")
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown")
	}

	// 인터페이스 조합 사용 (ReaderWriter)
	fmt.Println("\n// 인터페이스 조합 사용 (ReaderWriter)")
	var rw ReadWriter = bytes.NewBufferString("hello")
	buf := make([]byte, 5)
	rw.Read(buf)
	fmt.Println(string(buf))
	rw.Write([]byte(" world"))
	fmt.Println(rw)

	// 포인터 리시버 주의
	fmt.Println("\n// 포인터 리시버 주의")
	// var sp Speaker = Person{Name: "Bob"} // 컴파일 에러
	var sp Speaker = &Person{Name: "Bob"}
	fmt.Println(sp.Speak())
}

func main() {
	Interfaces()
}
