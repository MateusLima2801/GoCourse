package maps

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func Start() {
	sites := map[string]string{ // always dynamic
		"Google":              "google.com",
		"Amazon Web Services": "aws.com",
	}
	fmt.Println(sites)
	fmt.Println(sites["Amazon Web Services"])
	sites["Linkedin"] = "linkedin.com"
	fmt.Println(sites)

	delete(sites, "Google")
	fmt.Println(sites)
}

func Practice() {
	userNames := make([]string, 2, 5)
	userNames[0] = "Flavia"
	userNames[1] = "Lucas"
	fmt.Println(userNames)

	ratings := make(floatMap, 3)
	ratings["go"] = 4.7
	ratings["react"] = 4.8
	ratings["angular"] = 4.9
	ratings.output()

	for i := 0; i < len(userNames); i++ {
		fmt.Println(userNames[i])
	}

	for i, value := range userNames {
		fmt.Printf("%v: %v\n", i, value)
	}
	for i, value := range ratings {
		fmt.Printf("%v: %v\n", i, value)
	}
}
