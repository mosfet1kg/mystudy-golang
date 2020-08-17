package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	//println( len(ch) )
	defer func() {
		fmt.Println("test")
	}()
	defer wg.Done()


	for i := range ch {
		fmt.Printf("전송된 값: %d ", i)
	}

}

// main 함수는 프로그램의 진입점이다. 
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)
	// 채널에 10개의 정수를 보낸다.
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c) // for 반복문 안에서 range 키워드를 사용하면 채널이 닫힐 때까지 반복하면서 값을 꺼냅니다.
	wg.Wait()
}
