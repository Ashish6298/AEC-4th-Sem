package main

import (
	"fmt"
	"strconv"
)

func isPalidrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}

	}
	return true
}
func largestPalidromeProduct() (int, int, int) {
	largestPalidrome := 0
	var multiplicand1, multiplicand2 int
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := i * j
			if product < largestPalidrome {
				break
			}
			if isPalidrome(product) && product > largestPalidrome {
				largestPalidrome = product
				multiplicand1 = i
				multiplicand2 = j

			}

		}

	}
	return largestPalidrome, multiplicand1, multiplicand2
}

func main() {
	result, multiplicand1, multiplicand2 := largestPalidromeProduct()
	fmt.Printf("The largest palidrome product is : %d\n", result)

	fmt.Printf("The multiplicands are :%d and %d\n", multiplicand1, multiplicand2)
}
