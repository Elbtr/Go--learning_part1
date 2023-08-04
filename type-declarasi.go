package main

import "fmt"

// contoh string

func main() {
// kita dapat mendeclarasikan tipe suatu data dengan menggunakan type

type NoKTP string
type married bool


var hasaelNIK NoKTP = "123399172312"
var hasaelMarried married = false
fmt.Println(hasaelNIK) 
fmt.Println(hasaelMarried)


}

