# 01 — Binary & Bytes

ROADMAP: §1–§5

## Checklist konsep
- [ ] Decimal (base-10) dan nilai posisi
- [ ] Binary (base-2), konversi manual dua arah
- [ ] Bit dan byte; kenapa 1 byte = 0..255
- [ ] `byte` adalah alias `uint8`
- [ ] ASCII: karakter → angka
- [ ] Unicode vs UTF-8; satu karakter bisa > 1 byte
- [ ] `rune` adalah alias `int32`
- [ ] `string` vs `[]byte`; konversi dan mutabilitas

## Exercises
| Folder | Isi |
|---|---|
| `exercises/01-binary` | byte 72 → `01001000` → `H` |
| `exercises/02-ascii` | `"Hello"` → `[72 101 108 108 111]` |
| `exercises/03-utf8` | jumlah byte `A`, `é`, `你`, `🚀` |
| `exercises/04-string-vs-bytes` | `Hello` → `Jello` |

## Project
- [byte-explorer](byte-explorer/)

## Pertanyaan evaluasi fase
Jawab di `NOTES.md` **tanpa** membuka referensi:
1. Kenapa `len("🚀")` bukan 1? Berapa nilainya, dan kenapa?
2. Apa beda `for i := range s` dengan `for i := 0; i < len(s); i++` pada string UTF-8?
3. Kenapa `s[0] = 'J'` error untuk string, tapi boleh untuk `[]byte`?
4. Apakah `[]byte(s)` menyalin data? Kenapa itu penting?
5. Byte pertama `é` dalam UTF-8 adalah `11000011`. Apa arti bit `110` di depannya?
