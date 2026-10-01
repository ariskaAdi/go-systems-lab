# byte-explorer

## Spec
```bash
byte-explorer "Hé🚀"
```
Contoh output:
```text
text : "Hé🚀"
bytes: 7   runes: 3

idx  dec  hex  binary     char
0    72   48   01001000   H
1    195  c3   11000011   é (byte 1/2)
2    169  a9   10101001     (byte 2/2)
3    240  f0   11110000   🚀 (byte 1/4)
...
```

## Constraints
- [ ] Tanpa library pihak ketiga
- [ ] Hitung jumlah rune **sendiri** dari bit prefix UTF-8, baru bandingkan dengan `utf8.RuneCountInString`
- [ ] Logika dipisah dari `main` (misal `explore(s string) []Row`) supaya bisa dites

## Acceptance Criteria
- [ ] ASCII murni: setiap byte = satu karakter
- [ ] `é` (2 byte), `你` (3 byte), `🚀` (4 byte) ditampilkan dengan posisi byte yang benar
- [ ] Jumlah byte dan rune benar
- [ ] Tanpa argumen → baca dari stdin
- [ ] `go test ./01-bytes/byte-explorer/...` lulus, table-driven

## Apa yang ada di bawahnya?
- Bagaimana decoder UTF-8 tahu sebuah karakter terdiri dari berapa byte?
- Apa yang terjadi kalau byte-nya tidak valid UTF-8 (misal `"\xff"`)?
- Apa beda `%c` pada `byte` vs pada `rune`?

## Observability
Bandingkan output dengan `echo -n "Hé🚀" | xxd` (WSL) atau `Format-Hex` (PowerShell).

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
