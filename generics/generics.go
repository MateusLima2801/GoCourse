package generics

import "fmt"

func Start() {
	result := add(1, 2)
	fmt.Println(result)
}

func add[T int | float64 | string](a, b T) T { //pretty cool
	return a + b
}
