# myrun

## Spec
```bash
myrun ls -la
myrun go version
myrun sleep 5
myrun --timeout 3 sleep 10     # → "Process terminated" setelah 3 detik
myrun -v go version            # cetak PID, durasi, exit code ke stderr
```

## Constraints
- [ ] Pakai `os/exec`, `context`, `os/signal`
- [ ] Parsing flag dengan package `flag`; argumen setelah flag diteruskan apa adanya ke child

## Acceptance Criteria
- [ ] stdin / stdout / stderr child tersambung langsung ke terminal (interaktif: `myrun python3` bisa dipakai)
- [ ] Exit code child diteruskan sebagai exit code `myrun` (`myrun false; echo $?` → 1)
- [ ] Command tidak ditemukan → pesan jelas ke stderr, exit code 127
- [ ] `--timeout` mematikan child dan mencetak `Process terminated`
- [ ] Timeout mengirim `SIGTERM` dulu, lalu `SIGKILL` jika child belum berhenti dalam 2 detik (lihat `cmd.Cancel` dan `cmd.WaitDelay`)
- [ ] Ctrl+C pada `myrun` diteruskan ke child, bukan membunuh `myrun` duluan
- [ ] Test: menjalankan helper process (pola `TestHelperProcess` dari source `os/exec`)

## Apa yang ada di bawahnya?
- Gambar pohon proses saat `myrun sleep 10` berjalan (`pstree -p`).
- Siapa yang menerima Ctrl+C: `myrun`, child, atau keduanya? Kenapa? (process group)
- Apa yang dilakukan `exec.LookPath`?

## Observability
- `strace -f -e trace=process go run . sleep 1` → temukan `clone`/`execve`/`wait4`
- `ps -o pid,ppid,stat,cmd` saat child berjalan

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
