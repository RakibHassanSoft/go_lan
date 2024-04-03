package main

import "fmt"
func main()  {
	fmt.Println("Maps in go lan")


	languages := make(map[string]string)

	languages["js"]= "JavaScripts"
	languages["Rb"]= "Ruby"
	languages["py"]= "Python"
    
	fmt.Println(languages)

	fmt.Println("js sorts for : ",languages["js"])

	delete(languages,"Rb")
	fmt.Println(languages)




	// loop are interesting in golan

	for key,value:=range languages{
		fmt.Printf("For key %v , value is %v\n",key,value)
	}
	for _,value:=range languages{
		fmt.Printf("For key nothing, value is %v\n",value)
	}
}