package main

import "fmt"

// func main(){
// 	var firstName string = "bayu"

// 	var lastName string
// 	lastName = "andhika"

// 	fmt.Printf("Hello %s %s!\n", firstName,lastName)
// }

// ==================================
// Deklarasi Variable Tanpa tipe data
// ==================================


// func main(){
// 	name := "bayu"
// 	fmt.Printf("Hello %s \n",name);

// 	name = "andhika"
// 	fmt.Println(name)

// 	last := "resky"
// 	fmt.Println(last)
// }


// ===========================================
// deklarasi variabel mmenggunakan keyword new
// ===========================================


func main(){
	name := new(string)

	fmt.Println(name)
	fmt.Println(*name)
}