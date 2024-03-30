package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/4944945/homework-bootcamp-go-n9
*/
func main() {
	fmt.Print("Enter the elements of the first array : ")
	arr1, _ := getInput()
	fmt.Print("Enter the elements of the second array : ")
	arr2, _ := getInput()

	map2 := make(map[int]int)

	for _, num := range arr2 {
		map2[num] = getOrDefault(map2, num) + 1
	}

	var result []int

	for i := range arr1 {
		if num := getOrDefault(map2, arr1[i]); num > 0 {
			result = append(result, arr1[i])
			map2[arr1[i]] = num - 1
		}
	}
	fmt.Println("Intersection of two arrays is :", result)
}

func getInput() ([]int, bool) {
	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return []int{}, true
	}

	userInput = strings.TrimSpace(userInput)

	return convert(userInput), false
}

func getOrDefault(mapp map[int]int, val int) int {
	value, exists := mapp[val]

	if exists {
		return value
	}

	return 0
}

func convert(s string) []int {
	s = strings.Trim(s, " \n")

	ss := strings.Split(s, " ")
	arr := make([]int, len(ss))

	for i := range ss {
		arr[i], _ = strconv.Atoi(ss[i])
	}

	return arr
}
