package main

import "fmt"

// contoh conversion tipe data int

func main() {
	// var nilai32 = 127
	
	// // nilai32 di conversion ke dalam int yang lebih tinggi 
	// var nilai64 = int64(nilai32)

	// // nilai32 di conversion ke dalam int yang lebih rendah
	// var nilai8 = int8(nilai32)



	// hati hati dalam melakukan conversion karna jika nilai yang diconversi melebihin batas int nya maka nilai nya akan turun ke minimun nya (nilai akan turun ke dalam bentuk minus "-")

	// example
	var nilai32 = 128
	
	// nilai32 di conversion ke dalam int yang lebih tinggi 
	var nilai64 = int64(nilai32)

	// nilai32 di conversion ke dalam int yang lebih rendah
	var nilai8 = int8(nilai32)


	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

}

