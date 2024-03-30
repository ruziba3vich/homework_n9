package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Ketma-ket takrorlanuvchi belgilarni takrorlash sonidan keyin belgi bilan almashtirib,
	// siqib chiqaradigan dastur yarating.
	// Masalan, "aaabbbccc" "a3b3c3" ga aylanishi kerak.

	var s string
	fmt.Print("Enter the string : ") // Foydalanuvchidan matnni kiritishni so'raymiz
	fmt.Scan(& s) // Kiritilgan matnni o'qib olish

	var char byte = s[0] // Birinchi belgini olamiz
	var result strings.Builder // Yangi matn uchun strings.Builder obyekti yaratamiz
	/*
	`string` tipidagi o'zgaruvchiga element qo'shish bu har safar
	hotirning boshqa joyidan yangi string yaratib, eskini copy qilib yangi qo'shilyatgan 
	element bilan yaratish degani, ya'ni har safar yangi `string` yaratish bu hotiradan ham
	yutkizadi, tezlikdan ham yutkizadi, shu sababli StringBuilder ishlatish afzal, chunki 
	StringBuilder `slice` kabi ishlaydi, eski string ni oxiga yangi elementni append qib ishlaydi . . .
	*/
	count := 1 // `char` o'zgaruvchining o'zi birinchi deb hisoblanadi

	// Matnning birinchi belgisidan boshlab o'ngga ko'chib chiqamiz
	for i := 1; i < len(s); i++ {
		if s[i] == char { // Agar hozirgi belgi bir oldingi bilan bir xil bo'lsa
			count++ // Belgi sanog'ini oshiramiz
		} else { // Aks holda
			result.WriteByte(char) // Avvalgi belgini natijaga qo'shamiz
			result.WriteString(strconv.Itoa(count)) // Belgi sonini `result` matnga qo'shamiz
			count = 1 // Belgi sanasini qayta hisoblash uchun qaytadan 1 ga tenglaymiz
			/*
			`char` har safar birinchi sanoq bo'ladi, ya'ni "aaa" ning no'linchi indeksdagi chari
			bu `a` ning birinchi occurance i hisoblanadi, shu sababli `count` ni 1 ga tenglayapman
			*/
			char = s[i] // Hozirgi belgini yangi belgiga almashtiramiz
		}
	}

	result.WriteByte(char) // Oxirgi belgini natijaga qo'shamiz
	result.WriteString(strconv.Itoa(count)) // Oxirgi belgi sanasini matnga qo'shamiz
	/*
	Chunki oxirida `for` loop tugab qoladi, oxirgi natija `result` ga qo'shilmay qoladi
	*/

	fmt.Println("The result is:", result.String()) // Natijani chiqaramiz
}
