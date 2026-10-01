# mycat

## Spec
```bash
mycat hello.txt            # cetak isi file
mycat a.txt b.txt          # beberapa file, berurutan
echo hi | mycat            # tanpa argumen → stdin
```

## Constraints
- [ ] Tanpa Cobra / library CLI; pakai `os.Args`
- [ ] Tanpa `os.ReadFile` / `io.ReadAll` / `io.Copy`
- [ ] Pakai `os.Open`, `file.Read`, `os.Stdout.Write`
- [ ] Buffer `make([]byte, 1024)`, baca bertahap

## Acceptance Criteria
- [ ] Output identik dengan `cat` (`diff` kosong)
- [ ] File > 1 GB bisa dibaca tanpa memory naik (cek Task Manager / `top`)
- [ ] File tidak ada → pesan ke **stderr**, exit code 1, file lain tetap diproses
- [ ] File kosong dan file tanpa newline di akhir ditangani benar
- [ ] File biner (misal gambar) tersalin byte-per-byte: `mycat img.png > copy.png` identik
- [ ] `go test ./02-io/mycat/...` lulus

## Apa yang ada di bawahnya?
- Apa arti `n` dari `file.Read`? Kenapa harus menulis `buf[:n]`, bukan `buf`?
- Kapan `io.EOF` muncul? Bisakah `n > 0` bersamaan dengan `io.EOF`?
- Apa yang terjadi kalau file tidak di-`Close`, lalu `mycat` dipanggil dengan 10.000 file?
- Apa efek ukuran buffer 1 byte vs 1024 vs 64 KB terhadap kecepatan? (ukur!)

## Observability
- WSL: `strace -e trace=openat,read,write,close go run . f.txt` → hitung jumlah `read`
- Benchmark ukuran buffer dengan `go test -bench`

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
