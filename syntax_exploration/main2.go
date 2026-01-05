package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main2() {
	person := Person{Name: "John", Age: 20}
	fmt.Printf("This is out person %v\n",person)

	//anonymous structs
	employee := struct{
		name string
		id int
	}{
		name: "alice",
		id: 133,
	}
	fmt.Println(employee)


	fmt.Printf("The name of person is %s\n",person.Name)
	// modifyName(&person)
	person.modifyPersonalName("Abhishek")
	fmt.Printf("The name of person after modification is %s\n",person.Name)

	x := 20
	ptr := &x
	fmt.Printf("Value of X: %d and pointer is %p\n",x,ptr)
	*ptr = 30
	fmt.Printf("Value of X: %d and pointer is %p\n",x,ptr)
}

// since its name starts in lower case, its not exported to other files
func modifyName(person *Person){ //this fn is for whole file
	person.Name = "Abhishek"
	fmt.Printf("The modified name is %s\n",person.Name)
}

// since its name starts in lower case, its not exported to other files
func (person *Person) modifyPersonalName(name string){ //this fn is only for struct Person
	person.Name = name
	fmt.Printf("The modified name is %s\n",person.Name)
}