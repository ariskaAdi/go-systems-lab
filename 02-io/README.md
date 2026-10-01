# 02 — I/O

ROADMAP: §6, §7, §8, §9, §12

## Checklist konsep
- [ ] Buffer: ruang sementara berukuran tetap
- [ ] `io.Reader`: kontrak `Read(p []byte) (n int, err error)`
- [ ] `io.Writer`: kontrak `Write(p []byte) (n int, err error)`
- [ ] `io.EOF` bukan error sungguhan
- [ ] `n` bisa lebih kecil dari `len(p)` (short read)
- [ ] `os.File` adalah Reader dan Writer
- [ ] stdin / stdout / stderr; redirect `>` dan `2>`
- [ ] `defer f.Close()`

## Exercises
| Folder | Isi |
|---|---|
| `exercises/01-buffer` | `buf` vs `buf[:n]` |
| `exercises/02-reader-writer` | loop `Read()` manual, lalu bandingkan dengan `io.Copy` |
| `exercises/03-std-streams` | stdout vs stderr + redirect |

## Projects
- [mycat](mycat/)
- [mywc](mywc/)

## Pertanyaan evaluasi fase
1. Tuliskan dari ingatan: bentuk loop `Read()` yang benar, termasuk penanganan `n > 0` bersamaan dengan `io.EOF`.
2. Kenapa `io.Copy` tidak perlu tahu ukuran file?
3. Kenapa `program > out.txt` masih menampilkan pesan error di terminal?
4. Apa risikonya `os.ReadFile` pada file 10 GB?
5. Bagaimana `bufio.Reader` mengurangi jumlah pemanggilan `Read` ke OS?
