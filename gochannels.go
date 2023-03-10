package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var routines = 5

func addition(arr []string, chunks int, start int, end int) int {
	sum1 := 0
	fmt.Println(start, end)
	arr = arr[start:end]
	for _, i := range arr {
		number, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		sum1 += number
	}
	return sum1

}
func main() {

	fileData, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	arr := strings.Fields(string(fileData))
	chunks := len(arr) / routines
	fmt.Println(chunks)

	final := 0
	start := 0
	end := chunks
	for i := 0; i < routines; i++ {
		sum := addition(arr, chunks, start, end)
		start = start + chunks
		end = end + chunks
		fmt.Println("final sum: ", sum)

		final = final + sum
	}
	fmt.Println("grand sum: ", final)
}
