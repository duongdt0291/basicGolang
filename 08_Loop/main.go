package main

import "fmt"

func main() {
	//Go chỉ cung cấp duy nhất vòng lặp For, không có while hay do while như các ngôn ngữ khác
	//Cú pháp for thông thường
	//For init; cond; post {code body}
	fmt.Println("Example for basic For loop:")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//Không như các ngôn ngữ khác, các thành phần init, cond, post đều có thể bỏ qua
	//For cond {code} - Lúc này for sẽ như vòng lặp while trong các ngôn ngữ khác
	//Khi nào cond còn thỏa mãn sẽ vẫn chay;
	fmt.Println("Exp use For like While in another languages:")
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	//break - thoát hoàn toàn khỏi vòng lặp
	//đoạn code dưới sẽ chỉ in ra 0 ở màn hình
	//do sau khi in 0 ở lượt đầu tiên nó gặp câu lệnh break
	fmt.Println("Example for break:")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		break
	}

	//contiune - nhảy đến lượt kế tiếp, bỏ qua các câu lệnh còn lại của lượt đó
	//đoạn code dưới sẽ chỉ in ra những số lẻ 1 3
	//do câu lệnh if ở dưới kiểm tra nếu i là số chẵn thì sẽ nhảy đến lượt chạy tiếp theo
	fmt.Println("Example for continue:")
	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}

	//For {} - sẽ chạy vòng lặp vô tận, chỉ thoát khi gặp câu lệnh break;
	//Có thể hiểu For {} tương đương với For true {}
	fmt.Println("Exp with no init, cond, post statement:")
	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 3 {
			break
		}
	}
}
