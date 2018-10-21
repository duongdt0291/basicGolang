package main

import (
	"fmt"
	"sync"
)

//Có thể coi channels là 1 đường ống giúp kết nối các goroutine
//dữ liệu có thể truyền từ goroutine này sang goroutine khác

//chan T là channel có kiểu dữ liệu T
//channel có zero value là nil. nil channel ko có bâts kì tác dụng gì
//var ch chan int //ch là channel nil, vô tác dụng
//do vậy để khởi tạo 1 channel, ta dùng func make, tương tự như với slice và map

// func main() {
// 	c := make(chan int)

// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c <- i
// 		}
// 		wg.Done()
// 	}()

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c)
// 		}
// 		wg.Done()
// 	}()

// 	wg.Wait()

// 	// time.Sleep(time.Second)
// }

//Cùng nhìn lại ví dụ dùng Mutex ở phần trước,
//thử thay thế mutex bằng channel
var x int

func incrementWithChannel(wg *sync.WaitGroup, ch chan interface{}) {
	x++
	ch <- 1
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan interface{})
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incrementWithChannel(&wg, ch)
		<-ch
	}
	wg.Wait()
	fmt.Println("Expected x = 100")
	fmt.Println("Output x =", x) //100
	//Như vậy kết quả vẫn ra như khi ta sử dụng Mutex
	//Vậy khi nào nên sử dụng Mutex, khi nào nên dùng Channel
	//Có 1 qui tắc khá phổ biến như sau:
	//Mutex: caches, state
	//Channel: passing ownership of data, distributing units of work, communicating async results
}
