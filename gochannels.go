package main
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
var routines = 5
func main() {
	fileData, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	arr := strings.Fields(string(fileData))
	chunks := len(arr) / routines
	start := 0
	end := 25
	sum := 0
	for i := 0; i < routines; i++ {
		arr1 := arr[start : end-1]
		for _, i := range arr1 {
			number, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			sum += number
		}
		start = start + chunks
		end = end + chunks
	}
	fmt.Println("final sum: ", sum)
}
