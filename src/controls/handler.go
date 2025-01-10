package controls

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func printError(err error, message string) {
	if err != nil {
		fmt.Println(errors.New(strings.ToLower(fmt.Sprint("\033[31m", message, "\033[0m"))))
	}
}

func ConvertStringToInteger(data string) int {
	finalData, err := strconv.Atoi(data)
	if err != nil {
		printError(err, fmt.Sprint(err))
		return 0
	}
	return finalData
}

func ConvertStringToFloat(data string) float64 {
	finalData, err := strconv.ParseFloat(data, 64)
	if err != nil {
		printError(err, fmt.Sprint(err))
		return 0
	}
	return finalData
}

func PrintRedText(text string) {
	fmt.Println("\033[31m", text, "\033[0m")
}

func PrintGreenText(text string) {
	fmt.Println("\033[32m", text, "\033[0m")
}

func PrintYellowText(text string) {
	fmt.Println("\033[33m", text, "\033[0m")
}
