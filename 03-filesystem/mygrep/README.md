# mygrep

## Spec
```bash
mygrep "TODO" .
mygrep -i "todo" .
```
```text
src/main.go:10: TODO: refactor this
src/service.go:22: TODO: add validation
```

## Constraints
- [ ] Sequential, **tanpa goroutine** (versi concurrent ada di fase 06)
- [ ] Baca file per baris secara streaming (`bufio.Scanner` boleh; pahami batas 64 KB-nya)
- [ ] Logika dipisah: `searchFile(path, pattern string, ignoreCase bool) ([]Match, error)`

## Acceptance Criteria
- [ ] Output format `path:line: isi`
- [ ] `-i` case-insensitive
- [ ] File biner dilewati (deteksi byte `0x00` di 512 byte pertama)
- [ ] Folder `.git` dilewati
- [ ] Baris lebih panjang dari 64 KB tidak membuat program gagal diam-diam
- [ ] Exit code: 0 jika ada match, 1 jika tidak ada, 2 jika error (sama seperti `grep`)
- [ ] Ada `BenchmarkSearch` (dipakai sebagai baseline di fase 06)

## Apa yang ada di bawahnya?
- Kenapa `strings.ToLower` per baris itu mahal? Alternatifnya?
- Bagaimana `bufio.Scanner` menemukan batas baris di dalam buffer?

## Observability
- Ukur waktu pada repo besar (misal `$(go env GOROOT)/src`), catat sebagai baseline untuk fase 06

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
