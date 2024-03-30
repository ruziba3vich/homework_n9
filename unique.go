package main

import (
	"fmt"
	"math/rand"
)

func main3() {
	// 1. 1 dan 50 gacha bo'lgan tasodifiy sonni yarating (ishora: rand.Intn(max-min+1) + min)
	randNum := rand.Intn(49) + 1

	fmt.Println("Generated random value from 1 - 50 is ", randNum)

	// 2. Tasodifiy sonni length 20 bo'lgan slice yarating

	var slice [20]int
	// 3. Har bir tasodifiy sonni slice ichiga qo'shib chiqin

	for i := 0; i < 20; i++ {
		slice[i] = rand.Intn(5)
	}
	fmt.Println("Generated slice :", slice)
	// 4. Faqat noyob elementlardan iborat slice ni qaytaring

	var unique []int

	dict := make(map[int]int)

	for _, num := range slice {
		dict[num] = num
	}

	for key := range dict {
		unique = append(unique, key)
	}

	fmt.Println("A slice of unique elements :", unique)
}
