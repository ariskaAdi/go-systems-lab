# 08 — Networking (TCP)

ROADMAP: §23, §24, §25

## Checklist konsep
- [ ] IP, port, socket (konsep)
- [ ] `net.Listen` → `Accept` → `net.Conn`
- [ ] `net.Dial`
- [ ] TCP adalah **stream of bytes**, bukan pesan
- [ ] Message framing: delimiter (`\n`) vs length-prefix
- [ ] Satu goroutine per koneksi
- [ ] Read/write deadline
- [ ] Connection lifecycle: connect, half-close, close, `io.EOF`
- [ ] Shared state antar koneksi (mutex / channel)

## Projects
- [tcp-hello](tcp-hello/) — `server/` + `client/`
- [tcp-echo](tcp-echo/) — `server/` + `client/`
- [tcp-chat](tcp-chat/) — `server/` + `client/`

Jalankan:
```bash
go run ./08-networking/tcp-echo/server
go run ./08-networking/tcp-echo/client     # terminal lain
```

## Pertanyaan evaluasi fase
1. Client mengirim `"hello"` lalu `"world"`. Kenapa server bisa menerimanya sebagai satu `Read` `"helloworld"`, atau `"hel"` + `"loworld"`?
2. Apa yang dikembalikan `conn.Read` saat client menutup koneksi?
3. Kenapa satu koneksi lambat tidak boleh menghambat koneksi lain? Bagaimana Go mencegahnya?
4. Apa itu `TIME_WAIT`, dan kenapa terkadang server tidak bisa langsung bind ulang ke port yang sama?
5. Hubungkan dengan `ROADMAP §35`: jelaskan setiap baris `n, err := conn.Read(data); conn.Write(data[:n])`.
