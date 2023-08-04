package main

import "fmt"

// contoh data array

func main() {
	// contoh 1
	// membuat array secarah manual 

	var names [3]string

	names[0]="hasael"
	names[1] ="butar"
	names[2] = "butar"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// contoh 2
	// membuat array secarah langsung

	var nilai = [3]int{
		98,
		99,
		70,
	}

	fmt.Println(nilai)

	// untuk mengetahui panjang dari sebuah array kita dapat menggunakan method len

	fmt.Println(len(names))
	fmt.Println(len(nilai))

}

