package main

import "fmt"

// Exercise: cetak byte 72 sebagai desimal, %08b, dan %c (ROADMAP §2.3)
// Ketik ulang sendiri, jangan copy-paste dari ROADMAP.md.
func main() {
	var b byte = 72
	fmt.Println(b)          // cetak sebagai desimal
	fmt.Printf("%08b\n", b) // cetak sebagai biner (gunakan %08b)
	fmt.Printf("%c\n", b)   // cetak sebagai karakter (gunakan %c)

	fmt.Println('H' == 72)    // apakah 'H' dan 72 itu sama?
	fmt.Println('H' + 1)      // angka berapa?
	fmt.Printf("%c\n", 'H'+1) // huruf apa?
	fmt.Printf("%c\n", 97)    // huruf apa?

}

// go run ./01-bytes/exercises/01-binary
