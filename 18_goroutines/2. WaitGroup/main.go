package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup) {
	fmt.Println("Hello")
	//func (wg *WaitGroup) Done() trừ WaitGroup counter đi 1
	wg.Done()
}

func world() {
	fmt.Println("World")
}

func main() {
	//WaitGroup cung cấp các method để giúp chờ các goroutine thực hiện xog tiến trình
	//trước khi main goroutine (func main) chạy xong
	var wg sync.WaitGroup
	//func (wg *WaitGroup) Add(delta int) khai báo số lượng goroutine cho bộ đếm của WaitGroup
	//khi bộ đếm về 0 thì tất cả goroutine bị block bởi WG sẽ được giải phóng.
	wg.Add(1)
	//mỗi gorountine khi chạy xog sẽ phải gọi func Done
	go hello(&wg)
	// time.Sleep(1 * time.Second)
	world()
	//func (wg *WaitGroup) Wait() sẽ block cho tới khi bộ đếm của Wg về 0
	wg.Wait()
}

// Ví dụ 2
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(2)
// 	go foo()
// 	go bar()
// 	wg.Wait()
// }

// func foo() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("Foo:", i)
// 		//func Sleep(d Duration) sẽ tạm dừng goroutine hiện tại trong 1 khoảng thời gian d
// 		//nếu giá trị d <=0 thì sẽ ko có ảnh hưởng gì.
// 		time.Sleep(time.Duration(1 * time.Second))
// 		//time.Duration(1 * time.Second) là khai báo 1 giá trị của kiểu Duration(alias của int64)
// 		//có thể viết ngắn gọn hơn là time.Sleep(1 * time.Second)
// 	}
// 	wg.Done()
// }

// func bar() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("Bar:", i)
// 		time.Sleep(time.Duration(1 * time.Second))
// 	}
// 	wg.Done()
// }
