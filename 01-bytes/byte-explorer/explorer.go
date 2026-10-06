package main

type Row struct {
	Index int
	Value byte
	Char  string
	Part  int
	Total int
}

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

func isContinuation(b byte) bool {
	return b&0b1100_0000 == 0b1000_0000
}

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

/*
main.go                          explorer.go
───────                          ───────────
printReport ──┬──► countRunes ──┐
              │                 ├──► seqLen ──┬──► charLen          (baris 52)
              └──► explore ─────┘             ├──► secondByteOK     (baris 56)
                   (pakai Row)                └──► isContinuation   (baris 60)
*/
