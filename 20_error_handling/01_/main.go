package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	//tạo file log.txt
	//dùng hàm log.SetOutput làm nơi ghi các câu lệnh log
	//nếu không có thì các câu lệnh log hiện ra terminal bình thường (os.Stdout)
	nf, err := os.Create("log.txt")
	if (err != nil) {
		fmt.Println("Not successfully creat log file")
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("Test.txt")
	if err != nil {
		//fmt.Println() in ra os.Stdout (màn hình terminal)
		//fmt.Println("error happened,", err)

		//log.Println() mặc định in ra os.Stdout, có thêm timestamp so với fmt.Println
		//nếu có khai báo log.SetOutput thì in ra writter đk khai báo
		//log.Println("error happened,", err)

		//Fatalln : Println + os.Exit(1) - thường os.Exit(0) là ctrinh chạy thành công, khác 0 là lỗi
		//log.Fatalln(err)

		//func panic
		panic(err)
	}
}
