# 07 — Binary Format

ROADMAP: §21, §22

## Checklist konsep
- [ ] Integer → bytes → integer
- [ ] Big endian vs little endian; CPU-mu yang mana?
- [ ] Fixed-size vs variable-length field
- [ ] Length-prefix (cara menyimpan string di format biner)
- [ ] Magic number dan version
- [ ] `binary.BigEndian.PutUint32` vs `binary.Write`
- [ ] Membaca format yang rusak / terpotong dengan aman

## Exercises
| Folder | Isi |
|---|---|
| `exercises/01-endianness` | urutan byte BE vs LE |
| `exercises/02-binary-write` | struct ↔ bytes |

## Project
- [mydb](mydb/)

## Pertanyaan evaluasi fase
1. `123456` dalam BigEndian 4 byte = ? Dalam LittleEndian = ? (hitung manual dulu)
2. Kenapa network protocol umumnya memakai big endian ("network byte order")?
3. Kenapa string disimpan dengan length-prefix, bukan diakhiri `\0`?
4. Apa yang terjadi kalau `name length` di file rusak bernilai 4.000.000.000?
5. Kenapa magic number berguna? Contoh magic number nyata: PNG, ELF, PDF, ZIP.
