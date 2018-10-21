package main

import (
	"basic-golang/16_struct/defination/test"
	"fmt"
)

//Cấu trúc là một kiểu dữ liệu tập hợp bao gồm nhóm nhiều đối tượng có kiểu dữ liệu khác nhau.
//Mỗi đối tượng gọi là 1 trường (field).
//Bản chất dòng khai báo dưới là tạo 1 kiểu dữ liệu Person, còn kiểu nền là struct
//struct Person này được gọi là named struct, vì nó tạo ra 1 type mới
//mà từ đó có thể tạo các structure của type Person
//Lưu ý: nếu muốn package khác sử dụng được thì tên struct và các fields của nó đều phải
//viết in hoa chữ cái đầu, các struct và các fields ko viết in hoa chữ đầu sẽ coi là unexported
type Person struct {
	firstName string
	lastName  string
	age       int
}

//struct với anonymous fields
type Men struct {
	string
	int
}

func main() {
	//Tạo 1 structure từ type Person mà ko sử dụng tên trường(field name)
	//kiểu này cần khai báo các trường đúng theo thứ tự khi khai báo struct
	p1 := Person{"Duy", "Ram", 26}
	fmt.Println(p1) //{Duy Ram 26}
	fmt.Printf("kiểu của p1: %T\n", p1)
	//truy cập đến 1 trường trong struct
	fmt.Println(p1.firstName) //Duy

	//Tạo 1 structure variable từ type Person, sử dụng field name,
	//kiểu này ko quan trọng thứ tự các trường
	p2 := Person{
		firstName: "Long",
		lastName:  "Yen",
		age:       22,
	}
	fmt.Println(p2) //{Long Yen 22}
	fmt.Printf("kiểu của p2: %T\n", p2)

	//Tạo 1 structure variable từ type Person, chưa khai báo giá trị cho các field name
	//lúc này các fields sẽ tự động nhận zero value của type làm giá trị
	var p3 Person
	fmt.Println("p3:", p3) //p3: {   0}

	//dùng anonymous struct
	p4 := struct {
		name string
		age  int
	}{
		name: "Duy Long",
		age:  54,
	}
	fmt.Println("p4:", p4)              //{Duy Long 54}
	fmt.Printf("kiểu của p4: %T\n", p4) //struct { name string; age int }{Balotelli forward}

	//con trỏ đến struct
	newPerson := &Person{"Long", "Yen", 23}
	fmt.Println("Firstname of newPerson: ", (*newPerson).firstName)
	fmt.Println("Lastname of newPerson: ", (*newPerson).lastName)
	//Go cho phép không cần viết dấu * , có thể viết ngắn lại là newPerson.firstName
	fmt.Println("Firstname of newPerson: ", newPerson.firstName)

	//anonymous fields (embedded field)
	p8 := Men{"mrhayxinhun", 27}
	fmt.Println("p8:", p8) //p8: {mrhayxinhun 27}
	//Go sẽ tự coi tên anonymous field là tên type
	//Lưu ý là trong 1 struct ko thể có 2 anonymous field trùng nhau
	fmt.Println("p8.string =", p8.string)

	//Nested struct
	p9 := test.Worker{
		Name:     "Duy",
		Position: "Developer",
		Address: test.Address{
			Street: "Nguyen Dinh Chieu",
			City:   "Ha Noi",
		},
	}
	fmt.Println("p9:", p9)

	//nested struct
	//nếu 1 struct A có anonymous field là 1 struct khác
	//thì các trường trong anonymous struct đó sẽ coi như là nằm trực tiếp trong struct A
	//các trường này được gọi là promoted fields
	var p10 test.Worker2
	p10.Name = "Long"
	p10.Position = "DevOps"
	p10.Wife = "Yen"          //Wife là 1 promoted field
	p10.Daughter = "Long Yen" //Daughter là 1 promoted field
	fmt.Println("p10 :", p10)
	//tuy nhiên khi khai báo kiểu composit struct literals thì vẫn phải khai báo đầy đủ như trên p9
	//không thể viết như đoạn code dưới đây
	//terminal sẽ báo lỗi như sau
	//cannot use promoted field FamilyMember.Wife in struct literal of type test.Worker2
	//cannot use promoted field FamilyMember.Daughter in struct literal of type test.Worker2
	// p10 := test.Worker2{
	// 	Name:     "Long",
	// 	Position: "Developer",
	// 	Wife:     "Yen",
	// 	Daughter: "AAA",
	// }
	//Phải viết như sau
	// p10 := test.Worker2{
	// 	Name:     "Long",
	// 	Position: "Developer",
	// 	FamilyMember: test.FamilyMember{
	// 		Wife:     "Yen",
	// 		Daughter: "AAA",
	// 	},
	// }

	//Exported, unexported struct, fields từ các package khác
	p5 := test.Worker{
		Name:     "Balotelli",
		Position: "forward",
	}
	fmt.Println(p5)
	//WorkerUnexportedField là 1 exported struct của package test
	//tuy nhiên các trường của nó lại ko public
	// p6 := test.WorkerUnexportedField{
	// 	name: "Long",
	// 	age:  22,
	// }
	//unknown field 'name' in struct literal of type test.WorkerUnexportedField
	//unknown field 'age' in struct literal of type test.WorkerUnexportedField

	//struct là 1 value type, do vậy 2 struct sẽ so sánh được với nhau
	//nếu tất cả các trường của nó đều là type so sánh được
	//và bằng nhau nếu tất cả các trường giống nhau
	p6 := Person{"Duy", "Ram", 25}
	p7 := Person{"Duy", "Ram", 25}
	fmt.Println("p6 = p7 is", p6 == p7)
}
