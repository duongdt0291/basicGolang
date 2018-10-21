package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

//error trong GO là 1 type,
//do vậy nó cũng có thể được gán giá trị cho 1 biến, hoặc return trong func,...
var errorNegNum = errors.New("Lỗi, function không nhận số nhỏ hơn 0")

func main() {
	fmt.Printf("Type of errorNegNum is %T\n", errorNegNum) //*errors.errorString
	_, err := Sqrt(-10)
	//so sánh erro với nil, nếu bằng nil thì là ko có lỗi, nếu khác nil là có lỗi xảy ra
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if (f < 0) {
		//func New(text string) error
		//return 0, errors.New("Loi, so nhap vao nho hon 0.")
		//hoặc
		//return 0, errorNegNum
		//hoặc có thể dùng các func builtin của các package để tạo error
		return 0, fmt.Errorf("Lỗi, số nhập vào nhỏ hơn 0: %v", f)
	}

	return math.Sqrt(f), nil
}
