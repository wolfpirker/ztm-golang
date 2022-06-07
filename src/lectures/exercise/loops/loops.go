//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {		
		if i % 3 == 0 {
			fmt.Print("Fizz")
		} 
		if i % 5 == 0 {
				fmt.Print("Buzz")
		} else if i % 3 != 0 {
			fmt.Print("integer ", i)
		}
		fmt.Println("")
	}

}
