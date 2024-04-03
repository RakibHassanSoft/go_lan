package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Wellcome to time study of go lan")

	// presentTime := time.Now().Nanosecond()
	presentTime := time.Now()

	fmt.Println(presentTime)

	// fmt.Println(presentTime.Format("01-02-2006 ")) 
	fmt.Println(presentTime.Format("01-02-2006 Monday")) 
	// fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) 
	//01-02-2006 standered from formating 


	//creating date
	createdData := time.Date(2024,time.April,10,23,23,0,0,time.UTC)

	// fmt.Println(createdData)
	fmt.Println(createdData.Format("01-02-2006 Monday"))
}