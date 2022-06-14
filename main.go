package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	s "strings"
)

type Serve struct {
	Name  string
	Class string
}

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func (u *User) /*=>receiver*/ GetName() string { //func (t Type) MethodName(params)
	return u.name
}
func (u *User) getAge() int64 {
	return u.age
}
func (u *User) getGender() bool {
	return u.gender
}
func (u *User) getAddress() string {
	return u.address
}
func main() {
	fmt.Println("-----------------Bài 2-------------------") //đọc file Json

	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile("./serve.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var infor []Serve
	err = json.Unmarshal(content, &infor)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Let's print the unmarshalled data!
	log.Printf("name: %v\n", infor)
	for _, item := range infor {
		log.Printf(item.Name)
		log.Printf(item.Class)
	}
	fmt.Println("---------------Bài 3---------------")
	for _, check := range infor {
		lower := check.Class
		lower = s.ToLower(lower)
		if strings.Contains(lower, "admin") {
			fmt.Println(lower)
		}

	}
	// fmt.Println(strings.Contains("GeeksforGeeks", "for"))
	// fmt.Println(strings.Contains("A computer science portal", "science"))
	fmt.Println("---------------Bài 5---------------")
	insert := Serve{
		"fileCustome",
		"org.cofax.cds.FileServlet.Custome",
	}
	infor = append(infor, insert)
	fmt.Println(infor)

	fmt.Println("---------------Bài 6---------------")
	var slice = []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	a := slice[1:7]
	b := a[1:8]
	fmt.Println(slice)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Print("-----------------Bài 7---------------\n")
	getUser := User{}
	getUser.name = "Hà Hoàng Tú"
	getUser.age = 20
	getUser.gender = true
	getUser.address = "Hà Nội"

}
