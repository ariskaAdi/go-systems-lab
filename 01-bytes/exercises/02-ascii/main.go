package main

import "fmt"

// Exercise: iterasi []byte("Hello") dan cetak %d / %c / %08b (ROADMAP §3)
// Ketik ulang sendiri, jangan copy-paste dari ROADMAP.md.
func main() {
	var s = []byte("Hello")
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])          // cetak sebagai desimal
		fmt.Printf("%c\n", s[i])   // cetak sebagai karakter (gunakan %c)
		fmt.Printf("%08b\n", s[i]) // cetak sebagai biner (gunakan %08b)
	}
}

// go run ./01-bytes/exercises/02-ascii
