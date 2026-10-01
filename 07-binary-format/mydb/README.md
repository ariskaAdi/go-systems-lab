# mydb

Format file biner sederhana. Tujuannya memahami `struct → bytes → file → bytes → struct`, bukan membuat database.

## Format (semua integer big endian)

### Header (10 byte)
| Offset | Ukuran | Field | Nilai |
|---|---|---|---|
| 0 | 4 | magic | `MYDB` (`4D 59 44 42`) |
| 4 | 2 | version | `1` |
| 6 | 4 | record count | uint32 |

### Record (berulang `count` kali)
| Ukuran | Field |
|---|---|
| 4 | id (uint32) |
| 2 | name length (uint16) |
| N | name (UTF-8 bytes) |

## Spec
```bash
mydb init data.db
mydb add data.db 1 adi
mydb add data.db 2 "budi santoso"
mydb list data.db
mydb get data.db 2
```

## Constraints
- [ ] Pakai `encoding/binary`, tanpa `encoding/gob` / JSON
- [ ] Encode/decode terpisah dari I/O: `Encode(w io.Writer, db DB) error`, `Decode(r io.Reader) (DB, error)`

## Acceptance Criteria
- [ ] `xxd data.db` sesuai tabel format di atas (tempel hasilnya di bawah)
- [ ] Round-trip test: `Decode(Encode(x)) == x`
- [ ] Magic salah → error `not a mydb file`
- [ ] Version tidak dikenal → error yang jelas
- [ ] File terpotong di tengah record → error, **bukan** panic
- [ ] `name length` yang tidak masuk akal tidak membuat alokasi raksasa
- [ ] Fuzz test: `go test -fuzz=FuzzDecode ./07-binary-format/mydb` berjalan 1 menit tanpa panic
- [ ] Nama UTF-8 (`"你好"`) tersimpan benar (panjang = jumlah byte, bukan karakter)

## Hasil xxd
```text
(tempel di sini)
```

## Apa yang ada di bawahnya?
- Kenapa `binary.Write(w, binary.BigEndian, myStruct)` tidak bisa dipakai untuk struct yang berisi `string`?
- Bagaimana menambah field baru di version 2 tanpa merusak file version 1?
- Bagaimana format ini dibandingkan dengan SQLite file header? (baca dokumentasinya)

## Observability
- `xxd data.db` sebelum dan sesudah `add`

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
