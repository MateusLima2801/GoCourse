package pointers

import "fmt"

func StartPointers() {
	age := 32
	fmt.Println("Age: ", age)
	getAdultAge(&age)
	fmt.Println("Adult age: ", age)
}

// nil is the default value for an address

func PrintAges()           {}
func getAdultAge(age *int) { *age -= 18 }
