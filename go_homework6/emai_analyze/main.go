package main

import "fmt"

func main() {
	// Elektron pochta manzillarini tahlil qiling:
	// emailList o'zgaruvchidan domen nomlarini
	// (masalan, "example.com") chiqaradigan dasturni ishlab chiqing.
	// Keyin, har bir domen nomining takrorlanishini hisoblang va natijalarni chop eting.

	emailList := []string{
		"john@example.com",
		"alice@example.com",
		"bob@example.com",
		"sam@example.net",
		"lisa@example.org",
		"test@example.com",
		"example@example.com",
		"user1@gmail.com",
		"user2@gmail.com",
		"user3@gmail.com",
		"user4@yahoo.com",
		"user5@yahoo.com",
		"user6@outlook.com",
		"user7@outlook.com",
		"admin@example.com",
		"info@example.com",
		"contact@example.com",
		"support@example.com",
		"sales@example.com",
		"feedback@example.com",
		"webmaster@example.com",
	}

	// m - do'men nomi va uning takrorlanishlarini saqlaydigan map
	m := map[string]int{}

	// emailList ichidagi har bir elektron pochta manzilini tekshirib chiqish
	for _, str := range emailList {
		// convert funktsiyasini ishlatib, do'men nomini kesvolish
		nstr := convert(str)
		// m mapdagi ushbu do'men nomining occurance larini qo'shish
		m[nstr] = getOrDefault(m, nstr) + 1
	}

	fmt.Println("Berilgan ro'yxatda mavjud bo'lgan do'menlar soni:")

	// Natijalarni chiqarish
	for domain, occurrence := range m {
		fmt.Println(domain, ":", occurrence)
	}

}

// convert funksiya - berilgan elektron pochta manzilidan do'men nomini qirqib qaytaradi
func convert(s string) string {
	var i int

	// '@' belgigacha indeksni aniqlash -> chunki shu belgidan keyin domein nomi boshlanadi
	for s[i] != '@' {
		i++
	}
	// '@' belgisidan keyingi qismni qaytarish (do'men nomi)
	return s[i+1:]
}

// getOrDefault funksiya - mapdagi qo'shimcha qiymatni olish yoki 0 ga tenglash
func getOrDefault(m map[string]int, s string) int {
	// 
	// Bu degani agar map ichida so'ralgan string yo'q bo'lsa nil qaytishini oldini olish
	// 
	val, e := m[s]

	// Qiymatni olish
	if e {
		return val
	}
	// Agar qiymat topilmasa, nol qaytarish
	return 0
}
