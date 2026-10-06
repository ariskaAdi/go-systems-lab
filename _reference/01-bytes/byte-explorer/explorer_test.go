package main

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

var update = flag.Bool("update", false, "tulis ulang file golden di testdata/")

func TestCharLen(t *testing.T) {
	tests := []struct {
		b    byte
		want int
	}{
		{'H', 1},  // 01001000
		{0x7F, 1}, // 01111111, ASCII terakhir
		{0xC3, 2}, // 11000011, byte pertama é
		{0xE4, 3}, // 11100100, byte pertama 你
		{0xF0, 4}, // 11110000, byte pertama 🚀
		{0xA9, 0}, // 10101001, byte lanjutan
		{0xFF, 0}, // 11111111, tidak pernah valid
	}

	for _, tt := range tests {
		if got := charLen(tt.b); got != tt.want {
			t.Errorf("charLen(%08b) = %d, want %d", tt.b, got, tt.want)
		}
	}
}

func TestCountRunes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"kosong", "", 0},
		{"ascii", "Hello", 5},
		{"2 byte", "é", 1},
		{"3 byte", "你", 1},
		{"4 byte", "🚀", 1},
		{"campuran", "Hé🚀", 3},
		{"byte tidak valid", "\xff", 1},
		{"byte lanjutan sendirian", "\xa9", 1},
		{"karakter terpotong", "\xe4\xbd", 2},
		{"overlong", "\xc0\x80", 2},
		{"surrogate", "\xed\xa0\x80", 3},
		{"di atas U+10FFFF", "\xf4\x90\x80\x80", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countRunes(tt.input)
			if got != tt.want {
				t.Errorf("countRunes(%q) = %d, want %d", tt.input, got, tt.want)
			}
			// Pembanding: hasil hitungan sendiri harus sama dengan standard library.
			if std := utf8.RuneCountInString(tt.input); got != std {
				t.Errorf("countRunes(%q) = %d, utf8.RuneCountInString = %d", tt.input, got, std)
			}
		})
	}
}

func TestExplore(t *testing.T) {
	got := explore("Hé\xff")
	want := []Row{
		{Index: 0, Value: 'H', Char: "H", Part: 1, Total: 1},
		{Index: 1, Value: 0xC3, Char: "é", Part: 1, Total: 2},
		{Index: 2, Value: 0xA9, Part: 2, Total: 2},
		{Index: 3, Value: 0xFF},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("explore:\n got  %+v\n want %+v", got, want)
	}
}

func TestInput(t *testing.T) {
	got, err := input([]string{"Hello", "World"}, strings.NewReader("diabaikan"))
	if err != nil || got != "Hello World" {
		t.Errorf("dari argumen: got %q, %v", got, err)
	}

	got, err = input(nil, strings.NewReader("Hi\n"))
	if err != nil || got != "Hi\n" {
		t.Errorf("dari stdin: got %q, %v", got, err)
	}
}

// TestPrintReport membandingkan output lengkap dengan file golden.
// Setelah sengaja mengubah format output, jalankan: go test ./_reference/01-bytes/byte-explorer -update
func TestPrintReport(t *testing.T) {
	var buf bytes.Buffer
	printReport(&buf, "Hé🚀\n")

	golden := filepath.Join("testdata", "report.golden")
	if *update {
		if err := os.WriteFile(golden, buf.Bytes(), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	want, err := os.ReadFile(golden)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Errorf("output berbeda dari %s:\n%s", golden, buf.String())
	}
}

// FuzzCountRunes mencoba input acak dan memastikan hasilnya selalu sama dengan utf8.RuneCountInString.
// Jalankan: go test ./_reference/01-bytes/byte-explorer -fuzz=FuzzCountRunes -fuzztime=30s
func FuzzCountRunes(f *testing.F) {
	for _, seed := range []string{"", "Hello", "Hé🚀", "\xff", "\xe4\xbd", "\xed\xa0\x80"} {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, s string) {
		if got, want := countRunes(s), utf8.RuneCountInString(s); got != want {
			t.Errorf("countRunes(%q) = %d, utf8.RuneCountInString = %d", s, got, want)
		}
	})
}
