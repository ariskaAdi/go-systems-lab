# Evaluasi

Satu rubrik untuk semua project. Project dianggap **lulus (✅)** jika setiap aspek minimal skor 2.

## Rubrik

| Aspek | 0 | 1 | 2 | 3 |
|---|---|---|---|---|
| **Correctness** | tidak jalan | jalan untuk happy path | semua acceptance criteria lulus | output identik dengan tool asli, termasuk edge case |
| **Understanding** | tidak bisa menjelaskan | bisa menjelaskan sambil membaca kode | bisa menjawab "apa yang ada di bawahnya?" tanpa membuka kode | bisa menjelaskan ke orang lain + menggambar alur datanya |
| **Constraints** | memakai helper yang dilarang | sebagian constraint dilanggar | semua constraint dipatuhi | + bisa menjelaskan kenapa constraint itu ada |
| **Error handling** | error diabaikan / panic | error dicetak | error ke stderr, exit code benar, resource di-close | + pesan error jelas & konsisten seperti tool asli |
| **Testing** | tidak ada test | test happy path | test + `testdata/` + edge case | + table-driven / golden file / benchmark bila relevan |
| **Observability** | tidak dicek | dijalankan manual saja | dicek dengan tool yang relevan (lihat di bawah) | + temuan dicatat di README project |

## Edge case umum yang wajib diuji

- file kosong
- file tanpa newline di akhir
- file besar (> buffer, idealnya > 100 MB)
- karakter UTF-8 multi-byte (`é`, `你`, `🚀`)
- file / direktori tidak ada
- tidak punya permission
- input dari stdin (bukan file)

## Tool observability per fase

| Fase | Tool |
|---|---|
| 01–03 | `xxd` / `hexdump` (WSL) atau `Format-Hex` (PowerShell), `go test -cover` |
| 04–05 | `ps`, `pstree`, `strace -f`, `lsof -p <pid>` |
| 06 | `go test -race`, `go test -bench`, `go tool pprof`, `go tool trace` |
| 07 | `xxd` untuk memverifikasi layout file byte per byte |
| 08–09 | `curl -v`, `nc`, `ss -tnp` / `netstat`, Wireshark |
| 10 | `strace`, `go build -gcflags=-m` (escape analysis) |

## Cara verifikasi terhadap tool asli (WSL)

```bash
diff <(go run ./02-io/mycat f.txt) <(cat f.txt) && echo OK
diff <(go run ./02-io/mywc f.txt)  <(wc f.txt)
```

## Log evaluasi

| Tanggal | Project | Correct | Underst. | Constr. | Error | Test | Observ. | Lulus? |
|---|---|---|---|---|---|---|---|---|
| | | | | | | | | |
