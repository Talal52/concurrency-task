package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileData, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	arr := strings.Fields(string(fileData))
	arr = append(arr)

	sum := 0
	for _, i := range arr {
		number, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Println("The Sum of Array Items = ", sum)
}
