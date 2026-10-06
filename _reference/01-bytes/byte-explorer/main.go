package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, err := input(os.Args[1:], os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "byte-explorer:", err)
		os.Exit(1)
	}

	printReport(os.Stdout, text)
}

// input memakai argumen jika ada, selain itu membaca seluruh stdin.
// Newline dari stdin sengaja tidak dibuang: byte itu memang bagian dari input.
func input(args []string, stdin io.Reader) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	data, err := io.ReadAll(stdin)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func printReport(w io.Writer, text string) {
	fmt.Fprintf(w, "text : %q\n", text)
	fmt.Fprintf(w, "bytes: %d   runes: %d\n\n", len(text), countRunes(text))
	fmt.Fprintln(w, "idx  dec  hex  binary     char")

	for _, r := range explore(text) {
		fmt.Fprintf(w, "%-4d %-4d %02x   %08b   %s\n", r.Index, r.Value, r.Value, r.Value, describe(r))
	}
}

// describe mengisi kolom char.
func describe(r Row) string {
	switch {
	case r.Total == 0:
		return "� (invalid)"
	case r.Total == 1:
		return printable(r.Char)
	case r.Part == 1:
		return fmt.Sprintf("%s (byte 1/%d)", printable(r.Char), r.Total)
	default:
		return fmt.Sprintf("  (byte %d/%d)", r.Part, r.Total)
	}
}

// printable mengubah karakter kontrol seperti "\n" menjadi teks `\n` agar tabel tidak rusak.
func printable(s string) string {
	q := strconv.Quote(s)
	return q[1 : len(q)-1]
}
