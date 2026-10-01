# mypipe

## Spec
```bash
mypipe "ls" "grep go"
mypipe "cat big.txt" "grep error" "wc -l"
```
Setiap argumen adalah satu command; output satu menjadi input berikutnya.

## Constraints
- [ ] Pakai `os/exec` + `os.Pipe` atau `cmd.StdoutPipe`, **tanpa** memanggil `sh -c`
- [ ] Split argumen command cukup dengan `strings.Fields` (parsing beneran di gosh)

## Acceptance Criteria
- [ ] Mendukung N command, bukan hanya 2
- [ ] Semua proses berjalan **bersamaan** (uji: `mypipe "yes" "head -3"` harus selesai)
- [ ] Data tidak ditampung seluruhnya di memory Go (uji dengan file 1 GB)
- [ ] Exit code = exit code command terakhir (seperti bash)
- [ ] stderr setiap command tetap ke terminal

## Apa yang ada di bawahnya?
- Gambar diagram FD untuk 3 command: siapa memegang ujung baca/tulis yang mana?
- Kenapa `cmd.Output()` pada command pertama adalah implementasi yang salah?

## Observability
- `ls -l /proc/<pid>/fd` pada setiap proses saat pipeline berjalan (pakai `sleep` agar sempat dilihat)

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
