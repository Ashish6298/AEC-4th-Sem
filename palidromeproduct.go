package main

import "fmt"

func isPalindrome(n int) bool {
    original := n
    reverse := 0

    for n > 0 {
        remainder := n % 10
        reverse = reverse*10 + remainder
        n = n / 10
    }

    return original == reverse
}

func main() {
    largestPalindrome := 0

    for i := 100; i <= 999; i++ {
        for j := 100; j <= 999; j++ {
            product := i * j
            if isPalindrome(product) && product > largestPalindrome {
                largestPalindrome = product
            }
        }
    }

    fmt.Println("The largest palindrome product of two three-digit numbers is:", largestPalindrome)
}


// output 906609
