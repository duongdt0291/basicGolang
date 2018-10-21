package main

import "fmt"

func main() {
	//map là 1 built-in type của Go, chứa các cặp key-value không theo thứ tự
	//Cách tạo map bằng func make(map[type of key]type of value)
	//type of key ko được là reference type
	salaryPeople := make(map[string]int)

	//Tạo cặp key-value
	salaryPeople["Duy"] = 1000
	salaryPeople["Long"] = 1111
	fmt.Println(salaryPeople) //map[Duy:1000 Long:1111]

	//Lấy value của 1 key
	v1 := salaryPeople["Duy"]
	fmt.Println(v1) // 1000

	//Giá trị phụ trả về để biết có tồn tại key đó trong map không
	_, ok := salaryPeople["Duong"]
	fmt.Println(ok) //false
	_, ok = salaryPeople["Long"]
	fmt.Println(ok) // true

	//Xóa 1 cặp key, value trong map bằng func delete(map, key)
	delete(salaryPeople, "Long")
	_, ok = salaryPeople["Long"]
	fmt.Println(ok) //false

	//Khai báo thông thường, ko dùng func make
	m := map[string]int{"Long": 1111, "Duy": 1000}
	fmt.Println(m) //map[Long:1111 Duy:1000]

	m1 := map[string]int{}
	m1["Long"] = 1111
	fmt.Println(m1) //map[Long:1111]

	//Cách khai báo như sau sẽ khiến cho map có giá trị nil
	//việc thêm phần tử vào map có giá trị nil sẽ bị báo lỗi
	var m2 map[string]int
	if m2 == nil {
		fmt.Println("M2 is nil")
	}
	// m2["Duy"] = 1000 //panic: assignment to entry in nil map

	//For range for map
	var m3 = map[int]string{
		1: "learn golang",
		2: "learn go-pg",
		3: "more with postgres",
	}

	for key, value := range m3 {
		//Để ý là do map là 1 tập hợp ko theo thứ tự nào,
		//nên mỗi lần in ra, thứ tự các cặp key-value sẽ có vị trí khác nhau
		fmt.Printf("%d: %s\n", key, value)
	}

	//Lưu ý: với các biến reference type, chúng ta không thể so sánh chúng
	//với các kiểu dữ liệu khác, ngoài trừ nil
}
