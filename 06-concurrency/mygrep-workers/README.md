# mygrep-workers

Salinan [mygrep](../../03-filesystem/mygrep/) yang dibuat concurrent. **Sengaja dipisah** agar bisa dibandingkan dengan versi sequential.

Mulai dengan menyalin kode mygrep yang sudah lulus evaluasi:
```bash
cp ../../03-filesystem/mygrep/*.go .
```

## Spec
```bash
mygrep-workers --workers 8 "TODO" .
mygrep-workers --workers 8 -i "todo" .
```

## Arsitektur
```text
walker ──paths──► [worker 1..N] ──matches──► printer
```

## Constraints
- [ ] Worker pool dengan jumlah tetap (bukan satu goroutine per file)
- [ ] Hanya **satu** goroutine yang menulis ke stdout
- [ ] Tanpa library concurrency pihak ketiga (`errgroup` boleh **setelah** versi manual jadi)

## Acceptance Criteria
- [ ] Hasil sama dengan mygrep sequential (setelah `sort`)
- [ ] Baris output tidak pernah tercampur / terpotong
- [ ] `go test -race ./06-concurrency/...` bersih
- [ ] Opsi `--sorted`: output urut sesuai path walaupun diproses paralel
- [ ] Tidak ada goroutine leak saat error / Ctrl+C (cek `runtime.NumGoroutine()` di test)
- [ ] Benchmark workers = 1, 2, 4, 8, 16, 64 dibandingkan dengan baseline mygrep

## Hasil benchmark
| Workers | Waktu | vs sequential |
|---|---|---|
| sequential | | 1.0x |
| 1 | | |
| 2 | | |
| 4 | | |
| 8 | | |
| 16 | | |
| 64 | | |

Kesimpulan (kapan concurrency membantu, kapan tidak):

## Apa yang ada di bawahnya?
- Run pertama (cold cache) vs run kedua (warm page cache): apa bedanya, dan kenapa?
- Apa yang dilakukan Go scheduler saat goroutine menunggu disk?

## Observability
- `go test -bench . -cpuprofile cpu.prof` → `go tool pprof cpu.prof`
- `go test -trace trace.out` → `go tool trace trace.out`

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
