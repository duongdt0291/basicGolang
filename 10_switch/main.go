package main

import "fmt"

func main() {
	//Switch trong Go có cách hoạt động gần tương tự như trong các ngôn ngữ khác
	//Tuy nhiên, trong Go, switch case không cần có break, và có thêm câu lệnh fallthrough
	//Cú pháp
	//switch expression_switch {
	//	case exp_c1 :
	//		code
	//	case exp_c2 :
	//		code
	//	default :
	//		code
	//}
	name := "Duong"
	switch name {
	case "Duy":
		fmt.Println("Hello Duy")
	case "Long":
		fmt.Println("Hello Long")
	case "Duong":
		fmt.Println("Hello Duong")
	default:
		fmt.Println("User not found.")
	}

	//Dùng fallthrough thì khi gặp trường hợp đúng, sẽ vẫn chạy lệnh ở case phía dưới
	num := 3
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
		fallthrough
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Number not found")
	}

	//Khai báo nhiều expression 1 lúc ở case
	switch num {
	case 1, 2:
		fmt.Println("This is One or Two")
	case 3, 4:
		fmt.Println("This is Three or Four")
	default:
		fmt.Println("Number not found")
	}

	//Không khai báo expression_switch, thành giá trị mặc định là true
	switch {
	case name == "Long":
		fmt.Println("Hello Long")
	case name == "Duong":
		fmt.Println("Hello Duong")
	default:
		fmt.Println("User not found")
	}

	//Còn switch type, sẽ tìm hiều ở những bài sau.
}
