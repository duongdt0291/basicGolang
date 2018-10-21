package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//exported fields
type Person struct {
	FirstName, LastName string
	Age                 int
}

//unexported fields
type Workder struct {
	firstName, lastName string
	age                 int
}

func main() {
	//Exported fields
	p1 := Person{
		FirstName: "Duy",
		LastName:  "Long",
		Age:       25,
	}
	fmt.Println("p1 =", p1) //{Duy Long 25}

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("bs =", bs)
	fmt.Printf("Type of bs: %T\n", bs)      //[]uint8
	fmt.Println("string(bs) =", string(bs)) //{"FirstName":"Duy","LastName":"Long","Age":25}

	//unexported fields
	w1 := Workder{
		firstName: "Duy",
		lastName:  "Long",
		age:       25,
	}
	fmt.Println("w1 =", w1) //{Duy Long 25}
	bs, err = json.Marshal(w1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("bs =", bs)
	fmt.Printf("Type of bs: %T\n", bs)      //[]uint8
	fmt.Println("string(bs) =", string(bs)) //{}
	//Từ 2 vd trên, thấy rằng nếu là json ko encode các unexported fields

	//Encode
	json.NewEncoder(os.Stdout).Encode(&p1)
}
