package main

import "fmt"

func main()  {
	fmt.Println("Welcome to class in pointers")
	/*
--------------------------------------------------------------
	the operation will be perform on the actual value 
	----------------------------------------------------
	*/


	// var ptr *int
	// fmt.Println("The value of the pointer is",ptr)

	
	//refferece
	myNumber := 23

	var ptr = &myNumber
    fmt.Println("Value of actual pointer is ",ptr)
    fmt.Println("Value of actual pointer is ",*ptr)

	*ptr = *ptr * 2 //23 * 2
	fmt.Println("New value is : ",myNumber)
    
}