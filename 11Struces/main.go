package main

import "fmt"

func main() {
	
	fmt.Println("Starting learning struck")
	// There is no super or parent , no inheritance
	hitesh := User{"Hitesh", "hetesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are :%v\n", hitesh)
	fmt.Printf("Name is %v and Email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
