package main

import "fmt"

// contoh variable

func main() {
	// variable yang di declarasikan
	var name string

	name = "hasael"
	fmt.Println(name)
	// jika name di buat satu lagi maka go akan menghasilkan dua nila
	name = "hasael butar butar"
	fmt.Println(name)

	// variable yang tidak di declarasikan

	var marga = "butar butar"
	fmt.Println(marga)

	// variable yang berbentuk int

	var age = 32
	fmt.Println(age)

	// kita dapat mengganti var dengan := untuk menginisialisasikan sebuah type data

	country := "indonesia"
	fmt.Println(country)
	country = "belanda"
	fmt.Println(country)

	// variable yang multi fungsi
	var(
		firstName = "hasel"
		lastName = "butar butar"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}

