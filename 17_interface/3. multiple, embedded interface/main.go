package main

import "fmt"

//1 type có thể thực hiện 1 hoặc nhiều interface
type UserInfo interface {
	displayUserInfo()
}

type UserAddress interface {
	displayUserAddress()
}

type Person struct {
	name    string
	age     int
	city    string
	country string
}

func (p Person) displayUserInfo() {
	fmt.Printf("%s is %d years old.\n", p.name, p.age)
}

func (p Person) displayUserAddress() {
	fmt.Printf("City %s, country %s\n", p.city, p.country)
}

//Go cho phép chúng ta nhúng thẳng 1 interface này vào 1 interface khác, gọi là embedding interface
type UserManagement interface {
	UserAddress
	UserInfo
}

func showAllInfoUser(u UserManagement) {
	u.displayUserInfo()
	u.displayUserAddress()
}

func main() {
	p1 := Person{
		name:    "mrhayxinhun",
		age:     27,
		city:    "Ha Noi",
		country: "Viet Nam",
	}

	var i UserInfo = p1
	i.displayUserInfo()
	var a UserAddress = p1
	a.displayUserAddress()

	//Embedding interface
	fmt.Println("Example embedding interface")
	showAllInfoUser(p1)

	//Zero value của 1 interface là nil
	var u1 UserInfo
	//u1 interface chưa có giá trị nền nên zero value là nil
	//nó sẽ có cả underlying value và concrete type là nil
	fmt.Printf("Underlying value u1: %v, concrete type u1: %T", u1, u1)
}
