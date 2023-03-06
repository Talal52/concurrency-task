// package main

// import (
// 	"fmt"
// 	"strconv"
// 	)
	
// 	func main() {
// 	  fmt.Println(sum([]any{9, 1, "8", "2"})) // this should output 20
// 	}
	
// 	func sum(arr []any) int {
// 	  n:=0
// 	  for _, v := range arr{
// 		temp:=strconv.Atoi(v) //err: cannot initialize 1 variables with 2 values
// 		n+=temp //err: mismatched types int and any
// 	  }
// 	  return n
// 	}


package main

import "fmt"

func main() {

    numarray := []int{15, 25, 35, 45, 55, 65, 75}

    arrSum := 0

    for _, a := range numarray {
        arrSum = arrSum + a
    }
    fmt.Println("The Sum of Array Items = ", arrSum)
}