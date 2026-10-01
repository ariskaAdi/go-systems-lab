# Go Systems Programming

Repo belajar Go dari application-level menuju systems programming.
Materi lengkap ada di [ROADMAP.md](ROADMAP.md); cara menilai setiap project ada di [EVALUATION.md](EVALUATION.md).

> **Acuan urutan belajar adalah nomor folder di bawah ini.**
> Urutan di ROADMAP.md (isi bab, Phase list, Recommended Order) sempat berbeda-beda; yang dipakai adalah urutan folder.

## Cara menjalankan

Satu `go.mod` di root, jadi semua project dijalankan dari root:

```bash
go run ./02-io/mycat hello.txt
go test ./02-io/...
go test ./...            # semua
go vet ./...
```

## Struktur setiap fase

```text
NN-fase/
├── README.md       ← konsep, checklist, pertanyaan evaluasi fase
├── NOTES.md        ← catatanmu sendiri (tulis dengan kata-katamu)
├── exercises/      ← eksperimen kecil, tidak dinilai
└── <project>/      ← project yang dinilai, punya README.md sendiri
```

## Progress

Status: ⬜ belum · 🟨 sedang dikerjakan · ✅ lulus evaluasi

| # | Fase | Project | Status | OS |
|---|---|---|---|---|
| 01 | [Bytes](01-bytes/) | [byte-explorer](01-bytes/byte-explorer/) | ⬜ | Windows |
| 02 | [I/O](02-io/) | [mycat](02-io/mycat/) | ⬜ | Windows |
| 02 | | [mywc](02-io/mywc/) | ⬜ | Windows |
| 03 | [Filesystem](03-filesystem/) | [myfind](03-filesystem/myfind/) | ⬜ | Windows |
| 03 | | [mygrep](03-filesystem/mygrep/) | ⬜ | Windows |
| 04 | [Process](04-process/) | [myrun](04-process/myrun/) | ⬜ | WSL |
| 05 | [Pipe & Shell](05-shell/) | [mypipe](05-shell/mypipe/) | ⬜ | WSL |
| 05 | | [gosh](05-shell/gosh/) — L1 basic | ⬜ | WSL |
| 05 | | gosh — L2 pipe | ⬜ | WSL |
| 05 | | gosh — L3 redirect | ⬜ | WSL |
| 06 | [Concurrency](06-concurrency/) | [mygrep-workers](06-concurrency/mygrep-workers/) | ⬜ | Windows |
| 07 | [Binary Format](07-binary-format/) | [mydb](07-binary-format/mydb/) | ⬜ | Windows |
| 08 | [Networking](08-networking/) | [tcp-hello](08-networking/tcp-hello/) | ⬜ | Windows |
| 08 | | [tcp-echo](08-networking/tcp-echo/) | ⬜ | Windows |
| 08 | | [tcp-chat](08-networking/tcp-chat/) | ⬜ | Windows |
| 09 | [Protocol](09-protocol/) | [tiny-http](09-protocol/tiny-http/) | ⬜ | Windows |
| 10 | [Deeper Systems](10-deeper-systems/) | exercises: syscall, memory, unsafe | ⬜ | WSL |
| 11 | Final | gosh — final (`cd`, `&`, `jobs`) | ⬜ | WSL |

## Konvensi git

- Commit kecil per langkah, pesan menjelaskan *apa yang dipelajari*, bukan hanya *apa yang diubah*.
- Tag setiap project yang lulus evaluasi: `git tag mycat-done`.
- gosh dikembangkan di satu folder; setiap level ditandai tag:
  `gosh-L1-basic`, `gosh-L2-pipe`, `gosh-L3-redirect`, `gosh-final`.
  Evaluasi perubahan arsitektur dengan `git diff gosh-L1-basic gosh-L2-pipe`.

## Catatan OS

Fase 1–3 dan 6–9 aman di Windows. Fase 4, 5, 10, dan final gosh sebaiknya dikerjakan di **WSL**,
karena contoh di roadmap (`ls`, `sleep`, `grep`, signal, file descriptor, `strace`) berbasis Unix.
