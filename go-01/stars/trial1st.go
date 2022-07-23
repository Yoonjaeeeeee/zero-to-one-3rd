package main

import "fmt"

// var height int = 9

func main() {
	for i := 0; i < 3; i++ {
		for k := 0; k <= 1-i; k++ {
			fmt.Print(" ")
		}
		for j := 0; j <= 2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//merge 실험하기
