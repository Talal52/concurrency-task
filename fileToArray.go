package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"

)

// func findArraySum(arr []int) int{
// 	res := 0
// 	for i:=0; i<len(arr); i++ {
// 	   res += arr[i]
// 	}
// 	fmt.Println(res)
//  }

func findArraySum(arr []int) int{
	res := 0
	for i:=0; i<len(arr); i++ {
	   res += arr[i]
	}
	return res
 }

func main() {
	talal, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	arr := strings.Fields(string(talal))
	fmt.Println(arr)
	arr = append(arr,)

	// for i := 0; i < len(talal); i++ {
	// 	intVar, err := strconv.Atoi(arr[i])
	// 	fmt.Println(intVar, err, reflect.TypeOf(intVar))

		
	// }
	//findArraySum(arr )
    var t2 = []int{}

    for _, i := range arr {
        arr, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        t2 = append(t2, arr)
    }
    fmt.Println(t2)
	
	numarray := []int{15, 25, 35, 45, 55, 65, 75}

    fmt.Println("The Sum of Array Items = ",findArraySum(numarray) )
	fmt.Println(findArraySum([]int{1, 2, 3, 4, 5})) 
}
