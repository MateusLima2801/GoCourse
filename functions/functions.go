package functions

import "fmt"

func Start() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, func(number int) int { //anonymous function
		return number * 2
	})
	fmt.Println(doubled)

	transformed := transformNumbers(&numbers, createTransformer(5))
	fmt.Println(transformed)
}

func Recursion() {
	fmt.Println(factorial(5))
}

func Variadic() {
	fmt.Println(sumup(1, 10, 15, 40, -5))
}

func sumup(start int, numbers ...int) int {
	sum := 0
	for _, val := range numbers[start:] {
		sum += val
	}
	return sum
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	result := []int{}
	for _, val := range *numbers {
		result = append(result, transform(val))
	}
	return result
}

func double(number int) int {
	return number * 2
}

func multiplier(number int, factor int) int {
	return number * factor
}

func getTransformedFunction() func(int) int {
	return double
}

func createTransformer(factor int) func(int) int { //closure
	return func(number int) int {
		return number * factor
	}
}
