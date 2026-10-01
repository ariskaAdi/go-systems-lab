# tcp-echo

## Spec
```text
client> hello
server: hello
client> world
server: world
```

## Constraints
- [ ] Server memakai `conn.Read` / `conn.Write` manual dulu, baru bandingkan dengan `io.Copy(conn, conn)`

## Acceptance Criteria
- [ ] Banyak client sekaligus (goroutine per koneksi)
- [ ] Client mengirim dan menerima **bersamaan** (dua goroutine), tidak menunggu bergantian
- [ ] Client menutup → server mendeteksi `io.EOF`, mencetak log, menutup koneksi
- [ ] Idle timeout 60 detik dengan `SetReadDeadline`
- [ ] Graceful shutdown: Ctrl+C pada server menutup listener dan menunggu koneksi aktif selesai
- [ ] Test memakai `net.Pipe()` atau listener di `127.0.0.1:0`

## Apa yang ada di bawahnya?
- Kirim file 100 MB lewat echo server. Berapa kali `Read` dipanggil, dan berapa ukuran rata-rata `n`?
- Kenapa `io.Copy(conn, conn)` bisa bekerja padahal sumber dan tujuannya sama?

## Observability
- Wireshark / `tcpdump -i lo port 9000 -X`: lihat handshake SYN, SYN-ACK, ACK dan payload-nya

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
