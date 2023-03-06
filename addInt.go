package main

import ("fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum =sum + v
	}
	c <- sum
}

func main(){
	s:=[]int {8,38,99,96,30,55,11,35,29,40,2,94}


	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x+y)
}