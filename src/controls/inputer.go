package controls

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewReader(os.Stdin)
	reader, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return strings.TrimSpace(reader)
}

func InputStrictly(prompt string) string {
	for {
		input := Input(prompt)
		if CheckStrings(input) {
			return input
		}
		PrintRedText("Invalid Input")
	}
}
