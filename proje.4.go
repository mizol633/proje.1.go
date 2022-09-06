package main

import "fmt"

func main() {

	// Grade 3 olarak bir degisken olusturalım.
	grade := 3

	// Switch anahtar kelimesini yazalim.
	// Sonra karsilastiricagimiz degeri yazalım.
	// Sonrada suslu parantezimizi acıp caseleri yani karsilastigimiz durumlari belirtelim.
	// Switch anahtar kelimesini kullandikdan sonra yanina neyi karsilastircaksak onu yazariz.
	switch grade {

	case 5:
		fmt.Println("pekiyi")

	case 4:
		fmt.Println("iyi")

	case 3:
		fmt.Println("Orta")

	case 2:
		fmt.Println("Gecer")

	case 1:
		fmt.Println("Basarisiz")
	}

	// Switch case yapisi ile if else if yapisi arasindaki fark switch case yapisi fazla koşul olduğu zaman if else yapisina göre daha okunur hale gelmesi if else if arasnidaki tek fark bu.

	// Ayrica burada switch anahtar kelimesi yerine if else yapisi kullansakda kodumuz yinede ayni sonucu verecektir.
	// Cunku buradaki caseler if elseif yapisina denk geliyor.
	// Biz bunu birde if elseif olarak gösterelim.
	// Kodumuzu yazdik ve yine ayni sonucu aldik.
	if grade == 5 {
		fmt.Println("Pekiyi")
	} else if grade == 4 {
		fmt.Println("iyi")
	} else if grade == 3 {
		fmt.Println("Orta")
	} else if grade == 2 {
		fmt.Println("Geçer")
	} else if grade == 1 {
		fmt.Println("Basarisiz")
	}

}
