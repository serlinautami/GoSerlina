package main

	import
		"fmt"

func main() {
	var serlina = 11

	if serlina == 15 {
		fmt.Println("sangat baik bahasa inggris")
	} else if serlina > 10{
		fmt.Println("baik bahasa inggris")
	} else if serlina == 8 {
	  	fmt.Println("cukup bahasa inggris")	
	} else {
		fmt.Printf("tidak bisa bahasa inggris. nilai serlina %d\n", serlina)	
	}
}