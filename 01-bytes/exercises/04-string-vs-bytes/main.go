package main

import "fmt"

// Exercise: ubah []byte lalu kembali ke string: Hello -> Jello (ROADMAP §5)
// Ketik ulang sendiri, jangan copy-paste dari ROADMAP.md.
func main() {

	s := "Hello"
	b := []byte(s) // ubah string ke []byte
	fmt.Printf("%08b\n", b)
	b[0] = 'J'      // ubah byte pertama menjadi J
	s2 := string(b) // ubah []byte kembali ke string

	fmt.Println(s2) // cetak s2, harusnya Jello
}

// go run ./01-bytes/exercises/04-string-vs-bytes
