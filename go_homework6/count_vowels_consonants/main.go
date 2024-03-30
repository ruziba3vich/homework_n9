package main

import (
	"fmt"
	"strings"
)

var str = `
	Alice birinchi singluqni olib, 
	ning, bir yoki ikki marta, yengasi bilan birgalikda otirishdan juda charchab bo'lib chiqdi, 
	ya'ni u qizining o'qiyotgan kitobiga qarashini ishtirok etib, 
	lekin kitobda hech qanday rasmlar yoki muloqotlar yo'q edi, 
	'Alice o'yladi, 'kitobda rasmlar yoki muloqotlar bo'lmagan bo'lsa, uni qanday qilib foydalanish mumkin?'
	Shuning uchun u o'ziga maqulroq fikr qilmoqda (u shu qadar yaxshi ko'radi, 
	qayg'uli va og'irlik his qilish uchun), 
	qanday qilib qo'zg'alagan kasso uchun qiyinlikga loyiq daisi-iplar yaratish xushxoinlikning 
	qiyinlikiga ehtiyot bo'lishi mumkinmi? Qo'zg'alagan daislarni to'plashingizni tushunishni boshlaganida, 
	tez-tez yuqoriga qizil ko'zli oq sochib chiqqan oq yilancha bilan o'zgacha yo'g'irib ketdi.
`

func main() {
	// Unli va undoshlarni sanash:
	// Str o'zgaruvchi ni ichida unli va undoshlar sonini hisoblaydigan dastur tuzing.
	// Bo'shliqlarga e'tibor bermang va katta va kichik harflarni ekvivalent deb hisoblang.

	vowels := map[byte]byte{
		'a': 'a',
		'o': 'o',
		'e': 'e',
		'i': 'i',
		'u': 'u',
	}
	/*
	Go dasturlash tilida `Set` datatype yo'qligi uchun `map` ishlatdim
	chunki map ham set ham birxil Time Complexity da element qidiradi
	key ham value ham bir bo'lsa, set bilan bir xil ish bajaradi
	*/

	var vowelscnt int // unlilar sanog'i
	var consonantscnt int // undoshlar sanog'i

	for i := range str {
		if char := str[i]; char != ' ' { // probel emasligini tekshiramiz
			if _, exists := vowels[char]; exists { // harfni unli ekanligini tekshiramiz
				vowelscnt ++
			} else { // aks holda undosh deb qaraymiz
				consonantscnt ++
			}
		}
	}

	fmt.Printf("Berilgan matnda unli harflarning soni %d ta, va undosh harflarning soni %d ta\n", vowelscnt, consonantscnt)
	fmt.Println(len(str))
}

func ToLower(letter byte) string {
	/*
	kelayotgan char lowercase mi uppercase mi farqi yo'q, lowerga o'tkazib qaytaradi
	P.s -> berilgan `str` da sonlar yoki raqamlar ishlatilmagani uchun shunday qildim
	bo'lmasa IsDigit ligiga ham tekshirgan bo'lar edim
	*/
	lowercaseLetter := strings.ToLower(string(letter))

	return lowercaseLetter
}
