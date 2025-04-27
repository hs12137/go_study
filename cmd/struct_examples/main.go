package main

import (
	"encoding/json"
	"fmt"
)

// Person 정의 (패키지 레벨로 이동)
type Person struct {
	Name string
	Age  int
}

// User 정의 (메서드 예시용)
type User struct {
	Name string
	Age  int
}

// Employee 정의 (임베딩 예시용)
type Employee struct {
	Person
	Company string
}

// UserJSON 정의 (JSON 태그 예시용)
type UserJSON struct {
	Name string `json:"name,omitempty"` // JSON 직렬화 시 name 필드로 출력, 빈 값이면 제외
}

// Struct 함수: 구조체 관련 예제 모음
func Struct() {
	// 구조체 기본 사용
	fmt.Println("\n구조체 기본 사용")
	person1 := Person{Name: "정희석", Age: 28}
	fmt.Println(person1.Name) // 정희석
	fmt.Println(person1.Age)  // 28

	// 구조체 포인터 사용
	fmt.Println("\n구조체 포인터 사용")
	personPtr := &Person{Name: "희석", Age: 29}
	fmt.Println(personPtr.Name) // 포인터인데도 바로 필드 접근 가능 (자동 역참조)

	// 구조체 메서드 사용
	fmt.Println("\n구조체 메서드 사용")
	u := User{Name: "정희석", Age: 28}
	u.SayHi()
	u.Birthday()
	fmt.Println(u.Age) // 생일 이후 나이 29

	// 구조체 비교
	fmt.Println("\n구조체 비교")
	p1 := Person{Name: "정희석", Age: 28}
	p2 := Person{Name: "정희석", Age: 28}
	fmt.Println(p1 == p2) // true (필드 값이 모두 같으면 true)

	// 구조체 슬라이스 사용
	fmt.Println("\n구조체 슬라이스 사용")
	users := []Person{
		{Name: "A", Age: 20},
		{Name: "B", Age: 21},
	}
	for _, u := range users {
		fmt.Println(u.Name, u.Age)
	}

	// 익명 구조체 사용
	fmt.Println("\n익명 구조체 사용")
	anon := struct {
		Name string
		Age  int
	}{
		Name: "익명",
		Age:  99,
	}
	fmt.Println(anon)

	// 구조체 임베딩 (구조체 안에 구조체 포함)
	fmt.Println("\n구조체 임베딩")
	e := Employee{
		Person:  Person{Name: "김개발", Age: 30},
		Company: "GoCompany",
	}
	fmt.Println(e.Name)    // 임베딩으로 바로 필드 접근 가능
	fmt.Println(e.Company) // GoCompany

	// 구조체 JSON 태그 사용
	fmt.Println("\n구조체 JSON 태그 사용")
	userJson := UserJSON{Name: "희석"}
	jsonBytes, _ := json.Marshal(userJson)
	fmt.Println(string(jsonBytes)) // {"name":"희석"}

	// 구조체 JSON omitempty 사용
	fmt.Println("\n구조체 JSON omitempty 사용")
	userJsonEmpty := UserJSON{Name: ""}
	jsonBytes2, _ := json.Marshal(userJsonEmpty)
	fmt.Println(string(jsonBytes2)) // {} (Name이 빈 문자열이면 제외)
}

// User 타입 메서드들
func (u User) SayHi() {
	fmt.Printf("Hi, I'm %s\n", u.Name)
}

func (u *User) Birthday() {
	u.Age++
}

func main() {
	Struct()
}
