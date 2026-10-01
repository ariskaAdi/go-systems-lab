# tcp-hello

## Spec
Server listen di `:9000`, mengirim `Hello from server\n` ke setiap client lalu menutup koneksi.
Client connect, membaca sampai EOF, mencetak hasilnya.

## Acceptance Criteria
- [ ] Bekerja dengan client sendiri **dan** `nc localhost 9000` / `curl telnet://localhost:9000`
- [ ] Server tetap jalan setelah client pertama selesai (loop `Accept`)
- [ ] Server mencetak alamat client (`conn.RemoteAddr()`)
- [ ] Port sudah dipakai → pesan error jelas

## Apa yang ada di bawahnya?
- Apa yang terjadi di OS saat `Listen` dipanggil? Dan saat `Accept`?

## Observability
- `ss -tlnp | grep 9000` (WSL) atau `netstat -ano | findstr 9000` (Windows)

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
