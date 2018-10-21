package main

import (
	"fmt"
	"sync"
)

var x int

func incrementWithoutMutex(wg *sync.WaitGroup) {
	x++
	wg.Done()
}

// func main() {
// 	//race condition: xảy ra khi 2 goroutine cùng tác động đến 1 giá trị trong cùng thời điểm
// 	//khiến cho việc tương tác với giá trị đó ko được như ý
// 	//ví dụ:
// 	var wg sync.WaitGroup
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go incrementWithoutMutex(&wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("Expected x = 100")
// 	fmt.Println("Output x =", x)
// 	//Thử chạy với go run main.go, sẽ nhận được x không bằng 100 mà loanh quanh 91 92 ...
// 	//Chạy với go run --race main.go, sẽ in ra được x = 100, cùng với thông báo về race condition
// }

var mutex sync.Mutex //cung cấp cơ chế đảm bảo chỉ 1 goroutin thực hiện 1 đoạn code nhâts định tại 1 thời điểm

func incrementWithMutex(wg *sync.WaitGroup) {
	//đoạn code nằm giữa mutex.Lock() và mutex.Unlock()
	//sẽ chỉ được thực hiện bởi 1 goroutine tại 1 thời điểm
	//do vậy các goroutine sau phải chờ goroutine đó thực hiện xog mới được chạy
	//tránh được race condition
	mutex.Lock()
	x++
	mutex.Unlock()
	wg.Done()
}

func main() {
	//Dùng mutex để tránh race condition
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incrementWithMutex(&wg)
	}
	wg.Wait()
	fmt.Println("Expected x = 100")
	fmt.Println("Output x =", x)
}
