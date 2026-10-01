# tcp-chat

## Spec
```text
client A ──┐
client B ──┼──> TCP Server ── broadcast ──> semua client
client C ──┘
```
```text
$ /nick adi
* adi joined
adi: halo semua
budi: halo adi
$ /quit
```

## Constraints
- [ ] Goroutine per koneksi: `go handleConnection(conn)`
- [ ] Implementasi broadcast dua kali: (1) map + `sync.Mutex`, (2) satu goroutine "hub" + channel. Bandingkan.

## Acceptance Criteria
- [ ] Pesan dari satu client diterima semua client lain
- [ ] Join / leave diumumkan
- [ ] Client yang lambat / macet **tidak** menghambat client lain (buffered channel per client + drop / disconnect)
- [ ] Client putus tiba-tiba (kill process) dibersihkan dari daftar
- [ ] `go test -race` bersih
- [ ] **Challenge:** ganti framing `\n` dengan length-prefix (uint16 big endian + payload), memakai ilmu dari fase 07

## Apa yang ada di bawahnya?
- Mutex vs hub-channel: mana yang lebih mudah dibaca? Mana yang lebih sulit membuat deadlock?
- Apa yang terjadi kalau `conn.Write` ke client lambat dipanggil sambil memegang mutex?

## Observability
- Jalankan 100 client palsu (script / test), amati `runtime.NumGoroutine()` dan memory

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
