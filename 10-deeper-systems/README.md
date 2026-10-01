# 10 — Deeper Systems

ROADMAP: §28, §29, §30 · **Syscall dikerjakan di WSL**

Fase ini berupa eksperimen, bukan project. Hasilnya dipakai di **gosh final** (lihat [05-shell/gosh](../05-shell/gosh/)).

## Checklist konsep
- [ ] User space vs kernel space
- [ ] Syscall: `open`, `read`, `write`, `close`, `fork`/`clone`, `execve`, `socket`
- [ ] Bagaimana `os.File.Read` akhirnya menjadi syscall `read`
- [ ] Syscall berbeda per OS; build tag `//go:build linux`
- [ ] Stack vs heap, escape analysis
- [ ] Pointer, address, lifetime
- [ ] Struct layout, alignment, padding
- [ ] `unsafe.Pointer` dan aturan konversinya

## Exercises
| Folder | Isi | Verifikasi |
|---|---|---|
| `exercises/01-syscall` | menulis ke fd 1 tanpa `os` | `strace ./program` hanya menampilkan `write(1, ...)` milikmu |
| `exercises/02-memory` | escape analysis | `go build -gcflags=-m` → "moved to heap" |
| `exercises/03-unsafe` | layout struct | ubah urutan field → ukuran struct berubah |

## Pertanyaan evaluasi fase
1. Telusuri `os.File.Read` di source Go sampai ke syscall. Tulis rantai fungsinya.
2. Kenapa fungsi yang me-return `&x` membuat `x` pindah ke heap?
3. Struct `{a bool; b int64; c bool}` vs `{b int64; a bool; c bool}`: berapa ukuran masing-masing? Kenapa?
4. Kenapa GC Go membuat `unsafe.Pointer` → `uintptr` → `unsafe.Pointer` berbahaya?
5. Masalah memory apa yang akan diselesaikan oleh ownership & borrow checker di Rust? (jembatan ke ROADMAP §32)

## Observability
- `strace -c` pada mycat, mygrep, tcp-echo → bandingkan syscall yang dominan
- `go build -gcflags=-m=2`
- `GODEBUG=gctrace=1 go run .`
