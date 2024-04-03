package main

import "fmt"

func main()  {
	fmt.Println("Wlcome to array in go lan")

	var fruits [4]string 
	fruits[0] = "Apply"
	//[0] position is gap
	fruits[2] = "Greap"
	fruits[3] = "Banana"

	fmt.Println("Fruit list is : ",fruits)
	fmt.Println("Fruit list is : ",len(fruits))

	var vegList = [3]string{"Potato", " beans", "Mashroom"}

	fmt.Println("The vagy list : ",vegList)
	fmt.Println("The vagy list : ",len(vegList))

}