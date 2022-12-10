/*
 To print the pyramid below
i/p: n=6
 1
 22
 333
 55555
88888888
13131313131313131313131313
I/P: n=5
 1
 22
 333
 55555
 88888888
I/P: n=7
 1
 22
 333
 55555
 88888888
 13131313131313131313131313
 212121212121212121212121212121212121212121
 */
 
 package main

import (
	"fmt"
)

func main() {
	var a, b, temp int = 1, 2, 0
	var i, j, n int
	fmt.Println("Enter the value of n: ")
	fmt.Scan(&n)
	fmt.Println(a)
	fmt.Print(b)
	fmt.Print(b)
	fmt.Println()
	for i = 3; i <= n; i++ {
		temp = a + b
		for j = 1; j <= temp; j++ {
			fmt.Print(temp)
		}
		fmt.Println()
		a = b
		b = temp
	}

}


 
