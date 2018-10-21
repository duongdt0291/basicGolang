package main

import "fmt"

func main() {
	//sender có thể  đóng channel bằng func close(c Chan T) để thông báo cho receiver
	//rằng ko còn dữ liệu nào gửi đến cho channel nữa
	//receiver có thêm biến phụ khi nhận dữ liệu để check xem channel đã đóng hay chưa
	//cú pháp: v, ok := <- c
	//ok là true thì nhận dữ liệu thành công, ok false tức là channel đã đóng
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		close(c) //cần thiết để đóng khi chạy vòng lặp cho unbuffered channel
	}()

	for {
		v, ok := <-c
		//nếu ko kiểm tra biến ok, thì sẽ bị chạy vòng lặp vô tận
		if ok == false {
			break
		}
		fmt.Println(v)
	}

	//Go cung cấp for range để có thể duyệt giá trị trả về đến khi channel đóng
	//mà ko phải ktra biến ok như trên
	// for v := range c {
	// 	fmt.Println(v)
	// }
}
