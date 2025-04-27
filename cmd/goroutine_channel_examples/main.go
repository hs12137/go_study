package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// GoroutinesAndChannels (고루틴과 채널)
// Go에서 고루틴은 경량 스레드이고, 채널은 고루틴 간 통신을 위한 도구이다.
// 고루틴은 매우 가볍고, 수천~수만 개를 띄우는 것도 가능하다.
func GoroutinesAndChannels() {
	// 고루틴 기본 사용
	fmt.Println("\n// 고루틴 기본 사용")
	// go 키워드를 사용해 함수나 메서드를 비동기로 실행할 수 있다.
	// 고루틴은 main 고루틴과 별도로 동작하지만, main 고루틴이 끝나면 함께 종료된다.
	go sayHello()
	// 메인 고루틴이 바로 끝나지 않게 잠깐 대기
	time.Sleep(1 * time.Second)

	// 채널 기본 사용
	fmt.Println("\n// 채널 기본 사용")
	// 채널은 고루틴 간 안전하게 데이터를 주고받을 수 있는 통로이다.
	// 기본 채널은 동기 채널이며, 송신자와 수신자가 모두 준비되어야 데이터 전송이 완료된다.
	ch := make(chan int)
	go func() {
		ch <- 42 // 채널에 값 전송
	}()
	value := <-ch // 채널로부터 값 수신
	fmt.Println("채널로 받은 값:", value)

	// 버퍼 채널 사용
	fmt.Println("\n// 버퍼 채널 사용")
	// 버퍼 채널은 송신자가 수신자 없이도 버퍼가 찰 때까지 데이터를 보낼 수 있다.
	bufCh := make(chan int, 2)
	bufCh <- 1
	bufCh <- 2
	fmt.Println(<-bufCh)
	fmt.Println(<-bufCh)

	// 채널 닫기와 range
	fmt.Println("\n// 채널 닫기와 range")
	// 채널을 닫으면 (close), 수신 쪽에서 range로 안전하게 데이터를 모두 받을 수 있다.
	rangeCh := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			rangeCh <- i
		}
		close(rangeCh)
	}()
	for v := range rangeCh {
		fmt.Println(v)
	}

	// WaitGroup 사용하여 고루틴 기다리기
	fmt.Println("\n// WaitGroup 사용")
	// WaitGroup은 여러 고루틴의 완료를 기다릴 때 사용한다.
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done() // 고루틴이 끝나면 Done 호출
			fmt.Printf("Worker %d 시작\n", id)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Worker %d 완료\n", id)
		}(i)
	}
	wg.Wait() // 모든 고루틴이 끝날 때까지 대기

	// select 문 사용 (채널 다중 대기)
	fmt.Println("\n// select 문 사용")
	// select 문은 여러 채널을 동시에 감시하고, 준비된 채널을 처리한다.
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "채널 1 완료"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "채널 2 완료"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 발생")
	}

	// Mutex 사용하여 공유 데이터 보호
	fmt.Println("\n// Mutex 사용")
	var mu sync.Mutex
	counter := 0
	var wg2 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Println("최종 counter 값:", counter)

	// Context 사용하여 고루틴 제어
	fmt.Println("\n// Context 사용")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("고루틴 종료:", ctx.Err())
				return
			default:
				fmt.Println("작업 중...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
}

// sayHello 함수: 고루틴 예시용
func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	GoroutinesAndChannels()
}
