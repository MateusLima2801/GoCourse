package lists

import "fmt"

func Practice() {
	hobbies := []string{"reading", "watching movies", "doing necklaces"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
}
func Start2() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)

	// prices[2] = 9.99 //error
	prices = append(prices, 5.99)
	fmt.Println(prices)

	discountPrices := []float64{101.99, 99.99, 25.99}
	prices = append(prices, discountPrices...) // deconstructing array
	fmt.Println(prices)
}

func Start() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{1.0, 2.0, 3.0, 4.0}
	fmt.Println(prices)
	productNames[2] = "A Carpet"
	fmt.Println(productNames)
	fmt.Println(prices[0])

	featuredPrices := prices[1:3] //first included, last excluded
	//slice
	fmt.Println(featuredPrices)

	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)

	featuredPrices[0] = 7.0
	fmt.Println(featuredPrices)
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices))
	fmt.Println(len(featuredPrices), cap(featuredPrices))
	highlightedPrices := featuredPrices[:1]
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
