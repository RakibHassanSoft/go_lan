package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	wellcome := "Welcome to user input"
	fmt.Println(wellcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	//comma ok || err ok (to store reader data)
	// input, err or __ 
	// _ , err :=reader.ReadString('\n')  // do not care about input
	// input , err :=reader.ReadString('\n')
	input , _ :=reader.ReadString('\n') // do not care about error
	fmt.Println("Thannks for reading ",input)

	fmt.Printf("Type of this reating %T ",input)
}