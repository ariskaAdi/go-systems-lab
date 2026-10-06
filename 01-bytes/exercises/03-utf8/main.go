package main

import "fmt"

// Exercise: bandingkan jumlah byte A, é, 你, 🚀 (ROADMAP §4)
// Ketik ulang sendiri, jangan copy-paste dari ROADMAP.md.
func main() {

	A := "A"
	é := "é"
	你 := "你"
	rocket := "🚀"

	fmt.Println(len(A))      // berapa byte A?
	fmt.Println(len(é))      // berapa byte é?
	fmt.Println(len(你))      // berapa byte 你?
	fmt.Println(len(rocket)) // berapa byte 🚀?
}

// go run ./01-bytes/exercises/03-utf8
