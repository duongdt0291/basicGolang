package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		//ko thể dùng wg.Add(1) trong goroutine mà phải chạy luôn wg.Add(2) ở trên,
		//vì nếu viết trong goroutine thì bộ điều khiển còn chưa chạy vào thì func main đã kết thúc
		// wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		// wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		//ở đây do chưa c chưa close, nên goroutine này mới có thể chạy được
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
