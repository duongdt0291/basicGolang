package main

import "fmt"

func main() {
	//Ở những phần trước, dùng sync.WaitGroup để chờ các goroutine thực hiện hết
	//thay vì thế, có thể dùng toàn bộ channel để thực hiện quá trình này
	ch := make(chan int)
	chDone := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		chDone <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		chDone <- true
	}()
	go func() {
		//Phải biết được số lượng chDone chính xác
		//vì ko thể dùng for range và func close(chDone) trong trường hợp này
		<-chDone
		<-chDone

		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}
