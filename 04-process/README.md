# 04 — Process

ROADMAP: §13, §14, §15 · **Kerjakan di WSL**

## Checklist konsep
- [ ] File descriptor 0 / 1 / 2
- [ ] PID dan PPID
- [ ] Parent process vs child process
- [ ] fork + exec (konsep), dan bagaimana `os/exec` membungkusnya
- [ ] `Run` vs `Start` + `Wait`
- [ ] Exit code
- [ ] Signal: `SIGINT`, `SIGTERM`, `SIGKILL`
- [ ] `context.WithTimeout` + `exec.CommandContext`
- [ ] Zombie process

## Exercises
| Folder | Isi |
|---|---|
| `exercises/01-exec` | menjalankan command dan membaca output |
| `exercises/02-start-wait` | PID, `Start`/`Wait`, exit code |
| `exercises/03-signal` | menangkap Ctrl+C |

## Project
- [myrun](myrun/)

## Pertanyaan evaluasi fase
1. Apa beda `cmd.Run()` dengan `cmd.Start()` lalu `cmd.Wait()`? Kapan butuh yang kedua?
2. Apa yang terjadi kalau `Start()` dipanggil tapi `Wait()` tidak pernah dipanggil?
3. Kenapa `SIGKILL` tidak bisa ditangkap, sedangkan `SIGTERM` bisa?
4. Child process mewarisi apa saja dari parent? (env, cwd, file descriptor, ...)
5. Kenapa `exec.Command("ls | grep go")` tidak bekerja?
