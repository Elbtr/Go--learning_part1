package main

import "fmt"

// contoh operasi boolean

func main() {
	
	var ujian = 88
	var ipk = 3.0

	// cara pertama
	
	var nilaiUjian = ujian >= 85
	var nilaiIpk = ipk >= 3.2
	
	// penggunaan && (dan)
	var result bool = nilaiUjian && nilaiIpk
	// penggunaan || (atau)
	var result2 bool = nilaiUjian || nilaiIpk


	fmt.Println(result)
	fmt.Println(result2)

	// cara kedua

	// fmt.Println(ujian>=85 && ipk>=3.4)
	// fmt.Println(ujian>=85 || ipk>=3.4)
}

