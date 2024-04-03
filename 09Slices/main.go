package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to slices")

	var fruitList = []string  {"Apple","Greap"}

	fmt.Printf("The type of fruitList %T\n",fruitList)

	fruitList = append(fruitList, "Mango","banana")
   fmt.Println(fruitList)

   fruitList = append(fruitList[1:])
   fruitList = append(fruitList[:2])
   fruitList = append(fruitList[1:2])
   fmt.Println(fruitList)



   highScore := make([]int,4)
   highScore[0]= 234
   highScore[1]= 945
   highScore[2]= 123
   highScore[3]= 111
//    highScore[4]= 777

  // realocation value and lengh is 7
    highScore = append(highScore, 55,666,000)
    
   fmt.Println(highScore)
  //sorting the slices
   sort.Ints(highScore)

   fmt.Println(highScore)
   fmt.Println(sort.IntsAreSorted((highScore))) // true because it has shorted



  // how to reamove a value from a slices based on index

  var courses = []string {"ReactsJs", "javaScript","Sweft","Python"}
  fmt.Println(courses)
  var index int = 2
  courses = append(courses[:index],courses[index+1:]...)
  fmt.Println(courses)

}