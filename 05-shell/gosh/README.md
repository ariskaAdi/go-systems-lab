# gosh

Shell sederhana. **Satu folder, berkembang per level.** Setelah level lulus, buat git tag.

## Struktur yang disarankan
```text
gosh/
├── main.go          ← REPL loop: prompt → baca baris → parse → eksekusi
├── lexer.go         ← string → token (word, |, >, <, 2>, &)
├── parser.go        ← token → Pipeline{ Commands []Command, Background bool }
├── parser_test.go   ← table-driven, ini test paling penting
├── builtins.go      ← cd, pwd, exit, jobs
└── exec.go          ← menjalankan Pipeline
```

## Constraints (semua level)
- [ ] Tanpa library shell parser
- [ ] Tanpa `sh -c` / `bash -c`
- [ ] Parser terpisah dari eksekusi dan dites tanpa menjalankan proses

---

## Level 1 — Basic · tag `gosh-L1-basic`
```text
$ pwd
$ ls
$ echo hello
$ cat file.txt
$ exit
```
- [ ] Prompt `$ `, baca baris dengan `bufio`
- [ ] Builtin: `pwd`, `exit`
- [ ] External command dijalankan sebagai child
- [ ] Baris kosong diabaikan; Ctrl+D (EOF) keluar dengan rapi
- [ ] Command tidak ditemukan → `gosh: foo: command not found`, shell tetap jalan
- [ ] Mendukung quote: `echo "hello   world"`

## Level 2 — Pipe · tag `gosh-L2-pipe`
```text
$ ls | grep go
$ cat file.txt | grep a | wc -l
```
- [ ] N command dalam pipeline (reuse pemahaman dari `mypipe`)
- [ ] Builtin di dalam pipeline tidak membuat shell crash

## Level 3 — Redirect · tag `gosh-L3-redirect`
```text
$ echo hello > output.txt
$ echo again >> output.txt
$ cat < input.txt
$ command 2> error.log
```
- [ ] `>`, `>>`, `<`, `2>`
- [ ] Redirect + pipe bersamaan: `cat < in.txt | grep a > out.txt`
- [ ] File redirect di-close setelah command selesai

## Final — tag `gosh-final` (setelah fase 10)
```text
$ cd ..
$ sleep 10 &
$ jobs
```
- [ ] Builtin `cd` (termasuk `cd` tanpa argumen → home)
- [ ] Background job `&` + builtin `jobs`
- [ ] Ctrl+C membunuh foreground job, **bukan** shell
- [ ] Tidak ada zombie process setelah background job selesai

---

## Apa yang ada di bawahnya?
- Kenapa `cd` harus builtin, tapi `ls` tidak?
- Apa beda pipeline dengan builtin di bash (`echo hi | read x; echo $x`)?
- Bagaimana bash membuat Ctrl+C hanya mengenai foreground job? (process group, terminal foreground)

## Evaluasi perkembangan arsitektur
```bash
git diff gosh-L1-basic gosh-L2-pipe -- 05-shell/gosh
git diff gosh-L2-pipe gosh-L3-redirect -- 05-shell/gosh
```
Catat di sini: apakah desain `Command` / `Pipeline` di L1 bertahan, atau harus dirombak? Kenapa?

## Observability
- `pstree -p` saat pipeline berjalan
- `strace -f -e trace=process,dup2,pipe2 go run .`

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
