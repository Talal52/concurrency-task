// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"sync"
// )

// var routines = 5
// func addition(arr []string, chunks int, start int, end int, c chan int) {
// 	sum1 := 0
// 	fmt.Println(start, end)
// 	arr = arr[start:end-1]
// 	for _, i := range arr {
// 		number, err := strconv.Atoi(i)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		sum1 += number
// 	}
// 	c <- sum1
// }

// func main() {
// 	c:=make(chan int)
// 	fileData, err := os.ReadFile("text.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	arr := strings.Fields(string(fileData))
// 	chunks := len(arr) / routines
// 	fmt.Println(chunks)
// 	var x sync.WaitGroup
// 	x.Add(routines)
// 	start := 0
// 	end := chunks
// 	for i := 0; i < routines; i++ {
// 		go func(i int, start int, end int) {
// 			addition(arr, chunks, start, end, c)
// 			x.Done()
// 		}(i ,start,end)

// 		start = start + chunks
// 		end = end + chunks
// 	}

// 	go func() {
// 		x.Wait()
// 		close(c)
// 	}()

// 	final := 0
// 	for sum := range c {
// 		final = final + sum
// 	}

// 	fmt.Println("grand sum: ", final)
// }
