package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var routines = 5

func addition(arr []string, chunks int, start int, end int, c chan int) {
	sum1 := 0
	arr = arr[start : end-1]
	for _, i := range arr {
		number, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		sum1 += number
	}
	fmt.Println("sum of chunk: ",sum1)
	c <- sum1
}
func main() {
	c := make(chan int)
	fileData, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	arr := strings.Fields(string(fileData))
	chunks := len(arr) / routines
	final := 0
	start := 0
	end := chunks

	for i := 0; i < routines; i++ {
		go addition(arr, chunks, start, end, c)
		start = start + chunks
		end = end + chunks
	}

	for j := 0; j < routines; j++ {
		x := <-c
		final = final + x
	}
	fmt.Println("______________________")
	fmt.Println("sum of array: ", final)
	fmt.Println("______________________")
}
