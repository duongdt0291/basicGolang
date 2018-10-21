package main

import "fmt"

func main() {
	n := 10
	ch := make(chan int)
	chDone := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for v := range ch {
				fmt.Println(v)
			}
			chDone <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-chDone
	}
}
