package main 

import "fmt"

func main () {

// Bir x degiskeni olusturalim.
// Bu degiskene bir deger atayalim.

x := 06


// if eger anlamina gelmektedir.
// else degilse anlamına gelmektedir.
// Eger x sayisinin 2 ile bolumunun kalani 0 ise bize x cift sayidir yazdirsin.
// Eger x sayisinin 2 ile bolumunun kalani 0 degilse bize x tek sayidir yazdirsin.

      if x%2 == 0 {
   
	      fmt.Println(x, "cift sayidir")
		} else {
			
		  fmt.Println(x, "tek sayidir")
		}
        


}
