/* 1. Fizz Buzz
   - 1부터 100까지 숫자를 다음 규칙에 맞게 하나씩 출력하는 프로그램입니다.
   - 만약 3으로 나눠지는 숫자라면 숫자 대신 Fizz가 출력되게 해주세요.
   - 만약 5로 나눠지는 숫자라면 숫자 대신 Buzz가 출력되게 해주세요.
   - 만약 두 숫자 모두 나눌 수 있는 수라면 Fizz Buzz가 출력되게 해주세요. */

package main

import "fmt"

//  / 괄호 위치가 중요하다..!!! ㅠㅠㅠ/
func main() {
	for num := 1; num <= 15; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")
		} else if num%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(num)
		}
	}
}

/*if num%3 == 0 && num%5 == 0 {
	   fmt.Println("Fizz Buzz")
   } else if num%5 == 0 {
	   fmt.Println("Buzz")
   } else if num%3 == 0 {
	   fmt.Println("Fizz")
   } else {
	   fmt.Println(num)
   }
*/
