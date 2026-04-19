package file // must be in a folder of same name to allow creating the package

import (
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadFloatFromFile(fileName string) float64 {
	data, err := os.ReadFile(fileName)
	if err != nil {
		os.Create(fileName)
		return 0
	}
	result, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0
	}
	return result
}
