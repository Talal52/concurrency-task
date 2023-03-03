package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	talal, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	arr := strings.Fields(string(talal))
	fmt.Println(arr)
	arr = append(arr,)
}
