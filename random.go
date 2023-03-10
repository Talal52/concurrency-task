package main

// import ("fmt"
// )

// func sum(numarray []int,c chan int){
//     sum:=0
//     for _,i := range numarray{
//         sum+=i
//     }
//     c <- sum
// }

// func main() {

//     numarray := []int{15, 25, 35, 45, 55, 65, 75}

//     arrSum := 0

//     for _, a := range numarray {
//         arrSum = arrSum + a
//     }

//     fmt.Println("The Sum of Array Items = ", arrSum)

//     c :=make(chan int)
//     go sum(numarray[:len(numarray)/2],c)
//     go sum(numarray[len(numarray)/2:],c)
//     x,y := <-c, <-c

//     fmt.Println("sum of array: ", x+y)
// }
