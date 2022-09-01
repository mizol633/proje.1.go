// Package clause
package main

// Import statement
import "fmt"

// My Code
func main() {

	// Degısken tanımlarken Once anahtar kelıme sonra degısken ısmı sonrada degıskenın verı tıpı yazılır.
	// Metın tanımlarken string verı tıpını kulanırız.
	// Degıskenımız variable olsun, ısmı name olsun, verı tıpıde string olsun.
	var name string

	// Bu degıskene string bır deger verelım.
	name = "İbrahim"

	// Aynı sekılde buradada degısken tanımlarken once anahtar kelime sonra degısken ısmı sonrada degıskenın verı tıpı yazılır.
	// Ayrıca integer bır verı tıpıne string bır deger atayamayız çünkü integer tam sayı verı tıpıyken string ıse metin tanımlamak ıcın kullandıgımız bır veri tıpıdır eger atarsak yazdıgımız kod hata verıcektır.
	// Tamsayı tanımlarken integer verı tıpını kullanırız.
	// Degıskenımız variable olsun, degıskenımızın ısmı age olsun, degıskenımızın verı tıpıde integer olsun.
	var age int

	// Bu degıskene integer bır deger verelım.
	age = 15

	// Dogru-Yanlıs kontrolu yaparken bool verı tıpını kullanırız.
	// Degıskenımız variable olsun, degıskenımızın ısmı isMarried olsun, degıskenımızın verı tıpıde bool olsun.
	var isMarried bool

	// Bu degiskenede bool bır deger verelım.
	isMarried = false

	// Degıskenımız variable olsun, degıskenımızın ısmı weıght olsun, degıskenımızın verı tıpıde float32 olsun.
	// Ayrıca buradada float bır verı tıpıne integer bır deger atayamayız cunku float rasyonel sayıları tanımlamak ıcın kullandıgımız bır verı tıpıyken integer tamsayıları tanımlamak ıcın kullandıgımız bır verı tıpıdır.
	var weight float32

	// Degıskenımıze float bır deger verelım.
	weight = 71.5

	// Degıskenımız variable olsun, degıskenımızın ısmı size olsun, degıskenımızın verı tıpı float32 olsun.
	var size float32

	// Degıskenımıze float bır deger atayalım.
	size = 176.7

	// Ekrana yazdırma ıslemı yapmak ıcın fmt.Println kullanırız.
	fmt.Println(isMarried)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(weight)
	fmt.Println(size)

	// Go programlama dılı verı tıpını atamasakda onun hangı verı tıpınde oldugunu anlayabılır.

	// Degıskenlerın degerlerını, hesaplanan sonucları yada mesajları ekranda gostermek ıcın Printf kullanırız.
	// Degıskenlerın verı tıpını yazdırmak ıcınde tırnak ıcınde "%T" kullanırız.
	// Bır alt satıra gecmek ıcın \n kullanırız.
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T\n", weight)
	fmt.Printf("%T\n", size)

}
