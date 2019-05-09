package main

	import 
		"fmt"

func main() {
	var harga = 100000.0

	if percent := harga / 100; percent >= 100 {
    	fmt.Printf("%.1f%s sangatbaik!\n", percent, "%")
	} else if percent >= 70 {
    	fmt.Printf("%.1f%s baik\n", percent, "%")
	} else {
    	fmt.Printf("%.1f%s buruk\n", percent, "%")
    }
}
