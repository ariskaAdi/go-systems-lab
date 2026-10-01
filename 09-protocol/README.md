# 09 — Protocol (HTTP di atas TCP)

ROADMAP: §26, §27

## Checklist konsep
- [ ] HTTP/1.1 adalah teks di atas TCP
- [ ] Request line: method, path, version
- [ ] Header, `\r\n`, baris kosong sebagai pemisah
- [ ] `Content-Length` dan body
- [ ] Status line dan status code
- [ ] `Connection: close` vs keep-alive
- [ ] Perbandingan dengan `net/http`

## Exercises
| Folder | Isi |
|---|---|
| `exercises/01-raw-request-dump` | melihat request HTTP mentah dari curl / browser |

## Project
- [tiny-http](tiny-http/)

## Pertanyaan evaluasi fase
1. Bagaimana server tahu header sudah selesai? Bagaimana server tahu body sudah selesai?
2. Kenapa HTTP memakai `\r\n`, bukan `\n`?
3. Apa yang terjadi kalau `Content-Length` lebih besar dari body yang dikirim?
4. Apa yang dilakukan `net/http` yang **tidak** dilakukan tiny-http? (sebutkan minimal 5)
5. Gambar ulang diagram `TCP → bytes → parser → request → handler → response → bytes → TCP` dari ingatan.
