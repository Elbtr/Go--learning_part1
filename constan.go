package main

import "fmt"

// contoh constan

func main() {
	// constan jika data nya suda di deklarasikan maka data nya tidak akan bisa di ubah
	const firstName string = "hasael"
	const lastName = "butar butar"
	fmt.Println(firstName)
	fmt.Println(lastName)

	// contoh constan yang multi fungsi
	const (
		boyName string = "jack"
		girlName = "angel"
	)
	
	fmt.Println(boyName)
	fmt.Println(girlName)

}

