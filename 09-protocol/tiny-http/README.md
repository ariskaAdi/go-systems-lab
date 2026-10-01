# tiny-http

## Spec
```bash
curl http://localhost:8080/hello                         # → Hello
curl http://localhost:8080/users                         # → daftar user (JSON)
curl -X POST -d '{"name":"adi"}' http://localhost:8080/users   # → 201 Created
```

## Struktur yang disarankan
```text
tiny-http/
├── main.go         ← listen, accept, goroutine per koneksi
├── request.go      ← ParseRequest(r *bufio.Reader) (*Request, error)
├── request_test.go ← table-driven: request valid, rusak, terpotong
├── response.go     ← (*Response).WriteTo(w io.Writer)
└── router.go       ← map method+path → handler (sederhana, bukan framework)
```

## Constraints
- [ ] **Tanpa** `net/http` maupun `net/http/httptest`; hanya `net` + `bufio`
- [ ] Jangan membuat framework: tidak ada middleware, tidak ada path param

## Acceptance Criteria
- [ ] `GET /hello` → `200`, body `Hello`, `Content-Length` benar
- [ ] `GET /users` → `200`, JSON
- [ ] `POST /users` membaca body sesuai `Content-Length` → `201`
- [ ] Path tidak dikenal → `404`
- [ ] Method salah (`DELETE /hello`) → `405`
- [ ] Request rusak → `400`, server tidak crash
- [ ] Header dibaca case-insensitive (`content-length` = `Content-Length`)
- [ ] Bekerja dengan `curl -v` **dan** browser
- [ ] Body > batas (misal 1 MB) → `413`
- [ ] **Challenge:** keep-alive (beberapa request dalam satu koneksi)

## Perbandingan dengan net/http
Setelah selesai, tulis ulang dengan `http.ListenAndServe` di file terpisah (`compare/main.go`) dan catat:
| Aspek | tiny-http | net/http |
|---|---|---|
| Baris kode | | |
| Keep-alive | | |
| Chunked encoding | | |
| Timeout | | |
| Benchmark (`hey` / `wrk`) | | |

## Apa yang ada di bawahnya?
- Di mana di source `net/http` request line di-parse? (cari `readRequest` di `$(go env GOROOT)/src/net/http`)

## Observability
- `curl -v`, Wireshark pada port 8080

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
