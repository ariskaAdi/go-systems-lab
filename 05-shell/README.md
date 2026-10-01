# 05 — Pipe & Shell

ROADMAP: §16, §17, §18, §19, §31 · **Kerjakan di WSL**

## Checklist konsep
- [ ] Pipe: stdout satu proses → stdin proses lain
- [ ] `os.Pipe()` vs `cmd.StdoutPipe()`
- [ ] Kenapa ujung pipe yang tidak dipakai harus di-close (EOF tidak pernah datang)
- [ ] Redirect `>`, `>>`, `<`, `2>`
- [ ] Builtin vs external command (kenapa `cd` harus builtin)
- [ ] Tokenizing dan parsing input shell
- [ ] Background job `&`

## Exercises
| Folder | Isi |
|---|---|
| `exercises/01-stdout-pipe` | streaming output child |
| `exercises/02-os-pipe` | `ls \| grep go` manual |

## Projects
- [mypipe](mypipe/)
- [gosh](gosh/) — dikembangkan bertahap, setiap level ditandai git tag

## Pertanyaan evaluasi fase
1. Pada `ls | grep go`, apakah `ls` menunggu `grep` selesai, atau berjalan bersamaan? Buktikan.
2. Kenapa `yes | head -1` selesai, padahal `yes` tidak pernah berhenti? (SIGPIPE)
3. Kenapa `cd` tidak bisa dijalankan sebagai child process?
4. Pada `cmd > out.txt 2>&1`, apa yang sebenarnya terjadi pada file descriptor?
5. Apa yang terjadi kalau parent lupa menutup ujung tulis pipe?
