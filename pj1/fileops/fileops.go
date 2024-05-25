package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName , []byte(valueText), 0644)
}

func GetFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 1000, errors.New("Fail to find file")
	}

	valueText := string(data)
	value, error :=  strconv.ParseFloat(valueText, 64)

	if error != nil {
		return 1000, errors.New("Fail to parse stored value")
	}

	return value, nil
}