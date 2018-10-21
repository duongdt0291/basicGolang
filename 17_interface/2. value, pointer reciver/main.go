package main

import "fmt"

//value receiver và pointer receiver
type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	city, country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("City %s, country %s\n", a.city, a.country)
}

func main() {
	//method dùng value receiver
	var d1 Describer
	p1 := Person{"mrhayxinhun", 27}
	d1 = p1
	d1.Describe() //mrhayxinhun is 27 years old

	p2 := Person{"long", 24}
	d1 = &p2
	d1.Describe() //long is 24 years old

	//method dùng pointer receiver
	var d2 Describer
	a1 := Address{"Ha Noi", "Viet Nam"}
	//cách khai báo dưới đây sẽ có lỗi như sau
	//cannot use a1 (type Address) as type Describer in assignment:
	//Address does not implement Describer (Describe method has pointer receiver)
	//d2 = a1
	//Tức là nếu method của Address dùng pointer receiver, thì phải truyền vào đúng pointer
	//còn như cách trên, dùng value receiver chúng ta thấy linh động hơn, có thể truyền vào cả value và pointer
	//It is legal to call a value method on anything which is a value or whose value can be dereferenced.
	//do vậy, ở trường hợp này, cần viết như sau
	//The reason is that it is legal to call a pointer-valued method on anything
	//that is already a pointer or whose address can be taken.
	//The concrete value stored in an interface is not addressable
	//and hence it is not possible for the compiler
	//to automatically take the address of a in line no. 45 and hence this code fails.
	d2 = &a1
	d2.Describe() //City Ha Noi, country Viet Nam
}
