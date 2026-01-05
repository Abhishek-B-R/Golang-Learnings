package main

import (
	"fmt"
)

func main1() {
	var name string = "Abhi"
	fmt.Printf("This is my name %s\n",name)

	age := 27
	fmt.Printf("this is my age %d\n",age)

	var city string
	city = "Bangalore"
	fmt.Printf("this is my city : %s\n",city)

	var country,continent string= "India","Asia"
	fmt.Printf("This is my country %s and this is my continent %s\n",country,continent)

	var (
		isEmployed bool = true
		salary int = 50000
		position string = "developer"
	)

	fmt.Printf("isemployed : %t this is my salary: %d and this is my position: %s\n",isEmployed,salary,position)

	//zero values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Println(defaultInt,defaultFloat,defaultString,defaultBool)

	//constants
	const pi = 3.14 //if we use it or not, doesn't matter

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
		Jul
		Aug
	)

	fmt.Println(Jan,Feb,Mar,Apr,May,Jun,Jul,Aug)
	fmt.Println(add(2,3))

	day := "Tuesday"
	switch day{
		case "Monday":
			fmt.Println("start of the week")
		case "Tuesday","Wednesday","Thursday":
			fmt.Println("week in progress")
		case "Friday":
			fmt.Println("end of the week")
		default:
			fmt.Println("Weekend")
	}

	//loops
	for i := 0; i < 5; i++{
		fmt.Println("this is i",i)
	}

	// no while loop exists
	counter := 0
	for counter < 3 {
		fmt.Println("This is counter", counter)
		counter++
	}

	//arrays
	numbers := [5]int{1,2,3,4,5}
	fmt.Println(numbers)
	//  numbersAtInitialization := [...]int{1,2,3,4,5})

	matrix := [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Println(matrix)
	
	// slices
	allNumbers := numbers[:]
	firstThree := numbers[0:3]
	fmt.Println(allNumbers,firstThree)

	fruits := []string{"apple","banana","strawberry"}
	fmt.Printf("There are my fruits %v\n",fruits)

	fruits = append(fruits, "kiwi", "mango")
	fmt.Printf("these are my fruits with more fruits %v\n",fruits)

	moreFruits := []string{"blueberries","pineapple"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("these are my fruits with more fruits %v\n",fruits)

	for index, value := range fruits{
		fmt.Printf("%s fruit at index: %d\n",value,index)
	}


	//maps
	capitalCities := map[string]string{
		"USA":"Washington D.C.",
		"India":"New Delhi",
		"UK":"London",
	}
	fmt.Println(capitalCities["USA"])
	// fmt.Println(capitalCities["Germany"]) //no output

	capital, exists := capitalCities["Germany"]
	if exists{
		fmt.Println(capital)
	}else{
		fmt.Println("Does not exist for Germany")
	}

	delete(capitalCities,"UK")
	fmt.Println(capitalCities)
}

func add(a, b int) int {
	return a + b
}