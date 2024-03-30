package main

import "fmt"

/*

https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9
https://leetcode.com/problems/valid-palindrome/solutions/4945002/homewor-bootcamp-go-n9

*/

func main() {
	var s string
	fmt.Print("Enter the string : ")
	fmt.Scan(&s)
	x := isPalindrome(s)

	if x {
		fmt.Println(s, "is a palindrome !")
	} else {
		fmt.Println(s, "is not a palindrome !")
	}
}

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
