package main

import "fmt"

/*
    *
   ***
  *****
 *******
*********

2 * i + 1

4개의 층을 쌓아야 함
첫 번째는 공백으로 시작해야한다.
마지막 공백은 4개 (층이 5개일 때)
 *******
  *****
   ***
    *
*/

func main() {
	height := 4
	for i := 0; i < height; i++ {
		// i == 현재 층 (0층부터 4층까지)
		for j := height - i - 1; j > 0; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
