package main

// Row adalah satu baris tabel: satu byte beserta perannya di dalam karakter UTF-8.
type Row struct {
	Index int
	Value byte
	Char  string // karakter utuh, hanya diisi pada byte pertama karakter
	Part  int    // byte ke-berapa di dalam karakter (mulai dari 1)
	Total int    // jumlah byte karakter ini; 0 jika byte tidak valid
}

// charLen membaca bit depan b dan mengembalikan panjang karakter UTF-8 yang diawali b.
//
//	0xxxxxxx → 1    110xxxxx → 2    1110xxxx → 3    11110xxx → 4
//
// Mengembalikan 0 untuk byte lanjutan (10xxxxxx) dan byte yang tidak mungkin jadi awal karakter.
func charLen(b byte) int {
	switch {
	case b&0b1000_0000 == 0b0000_0000:
		return 1
	case b&0b1110_0000 == 0b1100_0000:
		return 2
	case b&0b1111_0000 == 0b1110_0000:
		return 3
	case b&0b1111_1000 == 0b1111_0000:
		return 4
	default:
		return 0
	}
}

// isContinuation melaporkan apakah b berbentuk 10xxxxxx.
func isContinuation(b byte) bool {
	return b&0b1100_0000 == 0b1000_0000
}

// secondByteOK menolak urutan yang pola bitnya terlihat benar tetapi dilarang oleh UTF-8:
//   - overlong: karakter ditulis dengan byte lebih banyak dari seharusnya (E0 80..9F, F0 80..8F)
//   - surrogate: U+D800..U+DFFF, khusus milik UTF-16 (ED A0..BF)
//   - di atas U+10FFFF, batas maksimum Unicode (F4 90..BF)
func secondByteOK(lead, b byte) bool {
	switch lead {
	case 0xE0:
		return b >= 0xA0
	case 0xED:
		return b <= 0x9F
	case 0xF0:
		return b >= 0x90
	case 0xF4:
		return b <= 0x8F
	}
	return true
}

// seqLen mengembalikan panjang karakter UTF-8 valid yang dimulai di s[i],
// atau 0 jika urutan byte di posisi itu tidak valid.
func seqLen(s string, i int) int {
	lead := s[i]

	// C0 dan C1 selalu overlong; F5..FF melewati U+10FFFF.
	if lead == 0xC0 || lead == 0xC1 || lead >= 0xF5 {
		return 0
	}

	n := charLen(lead)
	if n == 0 || i+n > len(s) {
		return 0
	}
	if n > 1 && !secondByteOK(lead, s[i+1]) {
		return 0
	}
	for j := i + 1; j < i+n; j++ {
		if !isContinuation(s[j]) {
			return 0
		}
	}
	return n
}

// explore memecah s menjadi satu Row per byte.
// Byte yang tidak valid menjadi Row dengan Total 0, lalu pembacaan lanjut ke byte berikutnya.
func explore(s string) []Row {
	rows := make([]Row, 0, len(s))

	for i := 0; i < len(s); {
		n := seqLen(s, i)
		if n == 0 {
			rows = append(rows, Row{Index: i, Value: s[i]})
			i++
			continue
		}

		for p := 0; p < n; p++ {
			row := Row{Index: i + p, Value: s[i+p], Part: p + 1, Total: n}
			if p == 0 {
				row.Char = s[i : i+n]
			}
			rows = append(rows, row)
		}
		i += n
	}

	return rows
}

// countRunes menghitung jumlah karakter di s tanpa package unicode/utf8.
// Sama seperti utf8.RuneCountInString, setiap byte tidak valid dihitung sebagai satu karakter.
func countRunes(s string) int {
	count := 0
	for i := 0; i < len(s); {
		n := seqLen(s, i)
		if n == 0 {
			n = 1
		}
		count++
		i += n
	}
	return count
}
