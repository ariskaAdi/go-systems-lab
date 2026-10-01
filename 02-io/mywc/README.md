# mywc

## Spec
```bash
mywc hello.txt
```
```text
Lines: 10
Words: 42
Bytes: 231
```
Tambahan: flag `-l`, `-w`, `-c` (pakai package `flag`), dan stdin jika tanpa file.

## Constraints
- [ ] Tanpa `os.ReadFile` untuk seluruh file
- [ ] Streaming dengan buffer tetap
- [ ] Logika hitung terpisah sebagai `count(r io.Reader) (Counts, error)` supaya bisa dites dengan `strings.NewReader`

## Acceptance Criteria
- [ ] Angka identik dengan `wc` (WSL) untuk semua file di `testdata/`
- [ ] Kata yang **terpotong batas buffer** tetap dihitung satu (uji dengan buffer kecil, misal 4 byte)
- [ ] File tanpa newline di akhir: lines sama dengan perilaku `wc`
- [ ] Whitespace beruntun (`"a   b\n\n\tc"`) dihitung benar
- [ ] Tab, `\r\n` (file Windows) ditangani
- [ ] `go test ./02-io/mywc/...` lulus, table-driven

## Apa yang ada di bawahnya?
- State apa yang harus dibawa dari satu chunk buffer ke chunk berikutnya?
- `Bytes` vs jumlah karakter: apa bedanya untuk file UTF-8? (`wc -c` vs `wc -m`)
- Kenapa `wc` menghitung `\n`, bukan "baris"?

## Observability
- Uji pada file 1 GB; ukur waktu & memory dibanding `wc`

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
