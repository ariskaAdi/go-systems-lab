# 06 — Concurrency

ROADMAP: §20

## Checklist konsep
- [ ] Goroutine vs OS thread (GOMAXPROCS)
- [ ] Channel buffered vs unbuffered
- [ ] `sync.WaitGroup`
- [ ] `sync.Mutex` vs channel: kapan memakai yang mana
- [ ] Data race dan `go test -race`
- [ ] `context` untuk cancellation
- [ ] Worker pool
- [ ] Goroutine leak
- [ ] CPU-bound vs I/O-bound

## Exercises
| Folder | Isi |
|---|---|
| `exercises/01-waitgroup` | menunggu banyak goroutine |
| `exercises/02-mutex` | data race lalu diperbaiki |
| `exercises/03-pipeline` | producer → workers → collector |
| `exercises/04-context-cancel` | cancellation |

## Project
- [mygrep-workers](mygrep-workers/)

## Pertanyaan evaluasi fase
1. Kenapa `counter++` dari banyak goroutine tidak aman? Apa yang terjadi di level instruksi?
2. Siapa yang harus menutup channel, producer atau consumer? Kenapa?
3. Bagaimana cara mendeteksi goroutine leak?
4. Pada mygrep, apakah bottleneck-nya CPU atau disk? Bagaimana membuktikannya?
5. Kenapa workers = 1000 belum tentu lebih cepat dari workers = 8?
