# Go Systems Programming --- Step by Step

Roadmap belajar **Go dari application-level menuju systems
programming**.

Tujuan utama repo ini bukan membuat REST API, tetapi memahami apa yang
terjadi di bawah abstraksi framework:

``` text
Text
 ↓
Binary
 ↓
Bit / Byte
 ↓
[]byte
 ↓
Buffer
 ↓
io.Reader / io.Writer
 ↓
File
 ↓
stdin / stdout / stderr
 ↓
Process
 ↓
Pipe
 ↓
TCP
 ↓
Binary Protocol
 ↓
Concurrency
 ↓
Syscall
 ↓
Memory / unsafe
```

> **Target akhir:** mampu membuat CLI, memahami data sebagai bytes,
> berinteraksi dengan filesystem/process/network, lalu memahami dasar
> systems programming menggunakan Go.

------------------------------------------------------------------------

# 0. Prerequisites

Karena sudah memahami dasar Go, tidak perlu mengulang:

-   variable
-   function
-   struct
-   interface dasar
-   slice
-   map
-   error handling
-   goroutine dasar
-   package/module

Yang perlu dikuatkan:

-   pointer
-   slice
-   interface
-   `io.Reader`
-   `io.Writer`
-   `defer`
-   error handling
-   concurrency

------------------------------------------------------------------------

# 1. Memahami Binary

Sebelum menyentuh systems programming, pahami bagaimana komputer
merepresentasikan angka.

## 1.1 Decimal

Sistem angka sehari-hari adalah base-10.

``` text
123

1 × 100
2 × 10
3 × 1

= 123
```

## 1.2 Binary

Binary menggunakan hanya:

``` text
0
1
```

Nilai posisi:

``` text
2^7  2^6  2^5  2^4  2^3  2^2  2^1  2^0
128   64   32   16    8    4    2    1
```

Contoh:

``` text
00000001 = 1
00000010 = 2
00000011 = 3
00000100 = 4
```

Untuk `72`:

``` text
01001000

= 64 + 8
= 72
```

## 1.3 Latihan

Konversikan sendiri:

``` text
00000001 = ?
00000010 = ?
00000101 = ?
00001010 = ?
00100000 = ?
01001000 = ?
11111111 = ?
```

Kemudian coba sebaliknya:

``` text
1  → ?
5  → ?
10 → ?
32 → ?
65 → ?
72 → ?
255 → ?
```

Jangan menggunakan calculator pada tahap ini.

------------------------------------------------------------------------

# 2. Bit dan Byte

## 2.1 Bit

Bit hanya memiliki dua kemungkinan:

``` text
0
1
```

## 2.2 Byte

1 byte = 8 bit.

``` text
00000000
```

sampai:

``` text
11111111
```

Jumlah kombinasi:

``` text
2^8 = 256
```

Karena angka dimulai dari 0:

``` text
0 ... 255
```

Maka:

``` go
byte
```

di Go adalah alias dari:

``` go
uint8
```

## 2.3 Eksperimen Go

Buat:

``` text
01-binary/
└── main.go
```

Isi:

``` go
package main

import "fmt"

func main() {
    var b byte = 72

    fmt.Println(b)
    fmt.Printf("%08b\n", b)
    fmt.Printf("%c\n", b)
}
```

Output:

``` text
72
01001000
H
```

Perhatikan:

``` text
72
 ↓
01001000
 ↓
H
```

------------------------------------------------------------------------

# 3. Character → ASCII → Byte

Sekarang pahami kenapa:

``` text
H = 72
```

ASCII memberikan angka kepada karakter.

Contoh:

``` text
A = 65
B = 66
C = 67

a = 97
b = 98
c = 99
```

Eksperimen:

``` go
package main

import "fmt"

func main() {
    data := []byte("Hello")

    fmt.Println(data)

    for _, b := range data {
        fmt.Printf("%d = %c = %08b\n", b, b, b)
    }
}
```

Output:

``` text
[72 101 108 108 111]

72  = H = 01001000
101 = e = 01100101
108 = l = 01101100
108 = l = 01101100
111 = o = 01101111
```

Sekarang kamu mulai melihat:

``` text
"Hello"
   ↓
[]byte
   ↓
72 101 108 108 111
   ↓
binary
```

------------------------------------------------------------------------

# 4. UTF-8 dan Kenapa Tidak Semua Character = 1 Byte

ASCII hanya mencakup karakter terbatas.

Coba:

``` go
package main

import "fmt"

func main() {
    examples := []string{
        "A",
        "é",
        "你",
        "🚀",
    }

    for _, text := range examples {
        fmt.Printf("%q -> %v -> %d bytes\n",
            text,
            []byte(text),
            len([]byte(text)),
        )
    }
}
```

Pahami perbedaan:

``` text
character
   ≠
byte
```

Satu character Unicode dapat membutuhkan beberapa byte dalam UTF-8.

Ini sangat penting ketika nanti bekerja dengan:

-   file
-   network
-   protocol
-   encoding
-   binary data

------------------------------------------------------------------------

# 5. `string` vs `[]byte`

Pelajari perbedaan:

``` go
text := "Hello"

data := []byte(text)
```

Sekarang:

``` text
string
 ↓
[]byte
```

dan:

``` go
textAgain := string(data)
```

kembali:

``` text
[]byte
 ↓
string
```

Eksperimen:

``` go
package main

import "fmt"

func main() {
    text := "Hello"

    data := []byte(text)

    fmt.Println(text)
    fmt.Println(data)

    data[0] = 'J'

    fmt.Println(string(data))
}
```

Output:

``` text
Hello
[72 101 108 108 111]
Jello
```

------------------------------------------------------------------------

# 6. Buffer

Sekarang masuk ke konsep penting dalam I/O.

Buffer adalah tempat sementara untuk menampung data.

``` go
buffer := make([]byte, 10)
```

Artinya kita memiliki ruang untuk 10 byte.

Visual:

``` text
┌────┬────┬────┬────┬────┬────┬────┬────┬────┬────┐
│    │    │    │    │    │    │    │    │    │    │
└────┴────┴────┴────┴────┴────┴────┴────┴────┴────┘
                     10 bytes
```

Pahami konsep:

``` text
source
  ↓
buffer
  ↓
destination
```

------------------------------------------------------------------------

# 7. `io.Reader` dan `io.Writer`

Ini salah satu konsep paling penting di Go.

## Reader

Sesuatu yang bisa dibaca:

``` go
io.Reader
```

Contoh:

``` text
file
stdin
network connection
buffer
```

## Writer

Sesuatu yang bisa menerima data:

``` go
io.Writer
```

Contoh:

``` text
file
stdout
network connection
buffer
```

Konsepnya:

``` text
Reader
   ↓
bytes
   ↓
Writer
```

Pelajari:

``` go
io.ReadAll()
io.Copy()
```

Tetapi jangan berhenti di helper tersebut.

Tujuan akhirnya adalah memahami apa yang dilakukan `Read()`.

------------------------------------------------------------------------

# 8. Project 1 --- `mycat`

Buat CLI:

``` bash
mycat hello.txt
```

Output:

``` text
Hello World
```

Jangan menggunakan library CLI seperti Cobra.

Gunakan:

``` go
os.Args
os.Open
file.Read
os.Stdout.Write
```

Target architecture:

``` text
hello.txt
   ↓
os.File
   ↓
[]byte buffer
   ↓
stdout
   ↓
terminal
```

## Challenge

Jangan membaca seluruh file sekaligus.

Gunakan buffer:

``` go
buffer := make([]byte, 1024)
```

Kemudian baca secara bertahap.

------------------------------------------------------------------------

# 9. Project 2 --- `mywc`

Buat:

``` bash
mywc hello.txt
```

Output:

``` text
Lines: 10
Words: 42
Bytes: 231
```

Pelajari:

-   byte counting
-   newline
-   whitespace
-   streaming
-   buffer

Challenge:

Jangan menggunakan `os.ReadFile()` untuk seluruh file.

------------------------------------------------------------------------

# 10. Project 3 --- `mygrep`

Buat:

``` bash
mygrep "TODO" .
```

Program mencari text pada file.

Contoh:

``` text
src/main.go:10: TODO: refactor this
src/service.go:22: TODO: add validation
```

Pelajari:

-   filesystem traversal
-   `filepath.Walk`
-   buffer
-   string processing
-   error handling

## Challenge

Tambahkan:

``` bash
mygrep -i "todo" .
```

`-i` berarti case-insensitive.

------------------------------------------------------------------------

# 11. Filesystem

Sekarang pelajari filesystem secara serius.

Gunakan:

``` go
os.ReadDir()
os.Stat()
os.Open()
os.Create()
os.Remove()
os.Rename()
filepath.Join()
filepath.Walk()
```

Buat project:

``` text
myfind
```

Usage:

``` bash
myfind . -name "*.go"
```

Pelajari:

``` text
file
directory
path
metadata
permission
size
modification time
```

------------------------------------------------------------------------

# 12. stdin / stdout / stderr

CLI tidak hanya menerima argument.

Ada tiga stream penting:

``` text
stdin
stdout
stderr
```

Visual:

``` text
keyboard
   │
   ▼
stdin
   │
   ▼
program
   ├──────> stdout ──────> terminal
   │
   └──────> stderr ──────> terminal
```

Eksperimen:

``` go
os.Stdin
os.Stdout
os.Stderr
```

Coba:

``` bash
program > output.txt
```

dan:

``` bash
program 2> error.txt
```

Pahami perbedaannya.

------------------------------------------------------------------------

# 13. File Descriptor

Sekarang masuk lebih dalam.

Di Unix-like systems, program berinteraksi dengan resource melalui file
descriptor.

Konsep umum:

``` text
0 = stdin
1 = stdout
2 = stderr
```

Visual:

``` text
Process

FD 0 ──> stdin
FD 1 ──> stdout
FD 2 ──> stderr
```

Pelajari konsep ini sebelum masuk ke pipe.

> Catatan: detail implementasi OS berbeda antara Windows dan Unix-like
> systems. Gunakan Linux/WSL jika ingin mengikuti contoh process/shell
> berbasis Unix dengan lebih mudah.

------------------------------------------------------------------------

# 14. Process

Sekarang belajar menjalankan program lain dari Go.

Gunakan:

``` go
os/exec
```

Buat:

``` text
myrun
```

Usage:

``` bash
myrun ls
myrun go version
myrun python script.py
```

Pelajari:

``` go
exec.Command()
cmd.Run()
cmd.Start()
cmd.Wait()
```

Kemudian pahami:

``` text
parent process
      │
      └── child process
```

------------------------------------------------------------------------

# 15. Project 4 --- Process Runner

Buat:

``` bash
myrun sleep 5
```

Kemudian tambahkan:

``` bash
myrun --timeout 3 sleep 10
```

Jika process lebih lama dari 3 detik:

``` text
Process terminated
```

Pelajari:

-   process lifecycle
-   context
-   timeout
-   cancellation
-   signals

------------------------------------------------------------------------

# 16. Pipe

Ini salah satu konsep terpenting dalam Unix-style systems programming.

Command:

``` bash
ls | grep go
```

Secara konsep:

``` text
ls
 │
 │ stdout
 ▼
PIPE
 │
 │ stdin
 ▼
grep
```

Jadi output program pertama menjadi input program kedua.

Di Go, pelajari:

``` go
cmd.StdoutPipe()
cmd.Stdin
```

Buat:

``` text
mypipe
```

yang mampu menjalankan pipeline sederhana.

------------------------------------------------------------------------

# 17. Project 5 --- Mini Shell

Sekarang gabungkan semua yang sudah dipelajari.

Buat:

``` bash
gosh
```

Target:

``` text
$ pwd
/home/adi

$ ls
main.go
go.mod

$ echo hello
hello

$ cat file.txt
hello world

$ exit
```

Architecture:

``` text
input
  ↓
parser
  ↓
command
  ↓
process
  ↓
stdin/stdout/stderr
```

Jangan menggunakan shell parser library terlebih dahulu.

------------------------------------------------------------------------

# 18. Mini Shell Level 2 --- Pipe

Tambahkan:

``` bash
$ ls | grep go
```

Architecture:

``` text
           pipe
ls ─────────────────> grep
 │                     │
stdout                stdin
```

Pelajari bagaimana process saling berkomunikasi.

------------------------------------------------------------------------

# 19. Mini Shell Level 3 --- Redirect

Tambahkan:

``` bash
$ echo hello > output.txt
```

dan:

``` bash
$ cat < input.txt
```

Kemudian:

``` bash
$ command 2> error.log
```

Target:

``` text
>
<
2>
```

Pahami bagaimana stdin/stdout/stderr dapat diarahkan.

------------------------------------------------------------------------

# 20. Concurrency

Sekarang baru gunakan concurrency untuk project systems.

Pelajari:

``` go
goroutine
channel
sync.Mutex
sync.WaitGroup
context
```

Buat:

``` text
mygrep --workers 8 "TODO" .
```

Architecture:

``` text
                 ┌── worker 1
files ───────────┼── worker 2
                 ├── worker 3
                 └── worker 4
                         │
                         ▼
                       result
```

Pahami kapan concurrency membantu dan kapan justru membuat program lebih
rumit.

------------------------------------------------------------------------

# 21. Binary Data

Sekarang kembali ke byte, tetapi lebih serius.

Pelajari:

``` go
encoding/binary
```

Contoh:

``` go
package main

import (
    "encoding/binary"
    "fmt"
)

func main() {
    data := make([]byte, 4)

    binary.BigEndian.PutUint32(data, 123456)

    fmt.Println(data)

    value := binary.BigEndian.Uint32(data)

    fmt.Println(value)
}
```

Pahami:

``` text
integer
   ↓
bytes
   ↓
binary representation
```

Kemudian pelajari:

``` text
Big Endian
Little Endian
```

Ini akan sangat berguna ketika belajar network protocol dan binary file
format.

------------------------------------------------------------------------

# 22. Project 6 --- Binary File Format

Buat format file sederhana:

``` text
MYDB

HEADER
├── magic number
├── version
└── record count

RECORD
├── ID
├── name length
└── name
```

Contoh:

``` text
[magic]
[version]
[count]
[id]
[name length]
[name bytes]
```

Gunakan:

``` go
encoding/binary
```

Tujuannya bukan membuat database production.

Tujuannya memahami:

``` text
struct
 ↓
bytes
 ↓
file
 ↓
bytes
 ↓
struct
```

------------------------------------------------------------------------

# 23. TCP

Jangan langsung menggunakan HTTP.

Mulai dari TCP.

Server:

``` go
net.Listen()
```

Client:

``` go
net.Dial()
```

Architecture:

``` text
Client
  │
  │ TCP connection
  ▼
Server
```

Buat:

``` text
tcp-server
tcp-client
```

Server mengirim:

``` text
Hello from server
```

------------------------------------------------------------------------

# 24. Project 7 --- TCP Echo Server

Buat:

``` text
tcp-echo-server
```

Client mengirim:

``` text
hello
```

Server mengembalikan:

``` text
hello
```

Kemudian:

``` text
hello
world
test
```

Pelajari:

``` go
net.Conn
Read()
Write()
Close()
```

Sekarang perhatikan:

``` text
TCP
 ↓
bytes
```

Kamu sudah kembali ke konsep awal.

------------------------------------------------------------------------

# 25. TCP Chat

Buat:

``` text
client A ──┐
           │
client B ──┼──> TCP Server
           │
client C ──┘
```

Setiap client mendapatkan goroutine:

``` go
go handleConnection(conn)
```

Pelajari:

-   concurrent connections
-   shared state
-   mutex
-   channels
-   connection lifecycle

------------------------------------------------------------------------

# 26. HTTP dari Raw TCP

Sekarang lakukan sesuatu yang menarik.

Jangan menggunakan:

``` go
net/http
```

terlebih dahulu.

Gunakan:

``` go
net.Listen()
```

Terima koneksi TCP.

Baca bytes.

Kemudian parsing:

``` http
GET /hello HTTP/1.1
Host: localhost:8080
```

Dan kirim:

``` http
HTTP/1.1 200 OK
Content-Type: text/plain
Content-Length: 5

Hello
```

Architecture:

``` text
TCP
 ↓
bytes
 ↓
HTTP parser
 ↓
request
 ↓
handler
 ↓
HTTP response
 ↓
bytes
 ↓
TCP
```

Setelah selesai, bandingkan dengan:

``` go
http.ListenAndServe()
```

Kamu akan mulai memahami apa yang sebenarnya dilakukan HTTP framework.

------------------------------------------------------------------------

# 27. Project 8 --- Tiny HTTP Server

Target:

``` bash
curl http://localhost:8080/hello
```

Output:

``` text
Hello
```

Tambahkan:

``` text
GET /hello
GET /users
POST /users
```

Jangan membuat framework.

Tujuannya memahami protocol.

------------------------------------------------------------------------

# 28. Syscall

Setelah filesystem, process, pipe, dan TCP sudah nyaman, mulai pelajari
syscall.

Konsep:

``` text
Your Go program
      │
      ▼
Go standard library
      │
      ▼
OS syscall
      │
      ▼
Kernel
      │
      ▼
Hardware
```

Pelajari secara konseptual:

``` text
open
read
write
close
fork/process creation
exec
socket
```

Detail syscall berbeda antara:

``` text
Linux
Windows
macOS
```

Karena itu jangan menganggap satu syscall API berlaku identik di semua
OS.

------------------------------------------------------------------------

# 29. Memory

Mulai memahami:

``` text
stack
heap
pointer
address
allocation
lifetime
```

Gunakan Go sebagai alat eksperimen.

Pelajari:

``` go
&
*
new()
make()
```

Contoh:

``` go
x := 10

p := &x

fmt.Println(x)
fmt.Println(p)
fmt.Println(*p)
```

Pahami:

``` text
x
 ↓
memory address
 ↓
value
```

------------------------------------------------------------------------

# 30. `unsafe`

**Jangan mulai dari sini.**

Setelah memahami memory, baru lihat:

``` go
unsafe.Pointer
```

Tujuannya bukan menggunakan `unsafe` dalam aplikasi biasa.

Tujuannya memahami:

``` text
type safety
memory layout
pointer conversion
low-level representation
```

------------------------------------------------------------------------

# 31. Final Project --- `gosh`

Gabungkan semuanya.

Target:

``` bash
gosh
```

Support:

``` text
$ pwd
$ cd ..
$ ls
$ echo hello
$ cat file.txt

$ ls | grep go

$ echo hello > output.txt

$ cat < input.txt

$ command &
$ jobs

$ command 2> error.log

$ exit
```

Architecture:

``` text
                 ┌───────────────┐
                 │   gosh shell  │
                 └───────┬───────┘
                         │
                    command parser
                         │
             ┌───────────┼───────────┐
             ▼           ▼           ▼
          process       pipe      redirect
             │           │           │
             └───────────┼───────────┘
                         ▼
                  stdin/stdout/stderr
                         │
                         ▼
                        OS
```

------------------------------------------------------------------------

# 32. Setelah Go Systems Programming

Setelah project di atas selesai, baru masuk Rust.

Roadmap:

``` text
Go Systems
    │
    ├── bytes
    ├── memory
    ├── process
    ├── filesystem
    ├── TCP
    ├── concurrency
    └── syscall
            │
            ▼
          Rust
            │
            ├── ownership
            ├── borrowing
            ├── lifetime
            ├── references
            ├── slices
            ├── traits
            ├── Result / Option
            ├── threads
            └── async
```

Dengan urutan ini, Rust akan lebih mudah dipahami karena kamu sudah tahu
**masalah apa yang ingin diselesaikan oleh ownership dan borrow
checker**.

------------------------------------------------------------------------

# 33. Urutan Belajar yang Disarankan

Jangan mengerjakan semuanya sekaligus.

## Phase 1 --- Binary & Bytes

-   [ ] Decimal
-   [ ] Binary
-   [ ] Bit
-   [ ] Byte
-   [ ] ASCII
-   [ ] Unicode
-   [ ] UTF-8
-   [ ] `byte`
-   [ ] `[]byte`

**Project:**

``` text
byte-explorer
```

------------------------------------------------------------------------

## Phase 2 --- I/O

-   [ ] Buffer
-   [ ] `io.Reader`
-   [ ] `io.Writer`
-   [ ] file
-   [ ] stdin
-   [ ] stdout
-   [ ] stderr

**Projects:**

``` text
mycat
mywc
```

------------------------------------------------------------------------

## Phase 3 --- Filesystem

-   [ ] directory
-   [ ] path
-   [ ] file metadata
-   [ ] permissions
-   [ ] traversal

**Project:**

``` text
myfind
mygrep
```

------------------------------------------------------------------------

## Phase 4 --- Process

-   [ ] PID
-   [ ] parent process
-   [ ] child process
-   [ ] exec
-   [ ] wait
-   [ ] timeout
-   [ ] signal

**Project:**

``` text
myrun
```

------------------------------------------------------------------------

## Phase 5 --- Pipe & Shell

-   [ ] pipe
-   [ ] stdin redirection
-   [ ] stdout redirection
-   [ ] stderr redirection
-   [ ] command parsing

**Project:**

``` text
gosh
```

------------------------------------------------------------------------

## Phase 6 --- Networking

-   [ ] TCP
-   [ ] connection
-   [ ] socket concept
-   [ ] binary protocol
-   [ ] concurrency

**Projects:**

``` text
tcp-echo
tcp-chat
```

------------------------------------------------------------------------

## Phase 7 --- Protocol

-   [ ] HTTP
-   [ ] request parsing
-   [ ] response parsing
-   [ ] headers
-   [ ] content length

**Project:**

``` text
tiny-http
```

------------------------------------------------------------------------

## Phase 8 --- Deeper Systems

-   [ ] syscall
-   [ ] memory
-   [ ] pointer
-   [ ] stack
-   [ ] heap
-   [ ] binary format
-   [ ] endianness
-   [ ] `unsafe`

------------------------------------------------------------------------

# 34. Aturan Belajar

## Rule 1 --- Jangan langsung menggunakan library

Kalau ingin memahami sesuatu:

``` text
buat versi sederhananya sendiri
        ↓
baru gunakan library
        ↓
bandingkan implementasinya
```

Contoh:

``` text
Raw TCP
 ↓
HTTP parser sederhana
 ↓
net/http
 ↓
Fiber
```

------------------------------------------------------------------------

## Rule 2 --- Selalu tanyakan "apa yang ada di bawahnya?"

Misalnya:

``` go
os.ReadFile()
```

Jangan berhenti di:

> "Ini membaca file."

Tanyakan:

``` text
Bagaimana file dibuka?
Bagaimana data dibaca?
Data itu masuk ke mana?
Kenapa perlu buffer?
Siapa yang mengalokasikan memory?
Bagaimana OS tahu file mana yang dibaca?
```

------------------------------------------------------------------------

## Rule 3 --- Gunakan debugger dan observability

Biasakan melihat:

``` text
input
 ↓
bytes
 ↓
memory
 ↓
output
```

Gunakan:

``` bash
go test
go test -race
go tool pprof
go tool trace
```

Kemudian, jika menggunakan Linux:

``` bash
strace
lsof
ps
top
hexdump
xxd
```

Tools ini akan membantu melihat hubungan antara program dan OS.

------------------------------------------------------------------------

# 35. Definition of Done

Kamu tidak perlu hafal semua API.

Target akhirnya adalah ketika melihat:

``` go
data := make([]byte, 4096)

n, err := conn.Read(data)

conn.Write(data[:n])
```

kamu bisa menjelaskan:

``` text
4096
  ↓
buffer 4096 byte
  ↓
connection membaca bytes
  ↓
n = jumlah byte yang benar-benar dibaca
  ↓
data[:n]
  ↓
hanya bagian valid yang ditulis kembali
```

Dan ketika melihat:

``` text
ls | grep go
```

kamu bisa membayangkan:

``` text
ls process
    │
    │ stdout
    ▼
  pipe
    │
    │ stdin
    ▼
grep process
```

Dan ketika melihat:

``` text
HTTP
```

kamu tidak lagi menganggapnya sebagai sesuatu yang ajaib:

``` text
HTTP
 ↓
TCP
 ↓
bytes
 ↓
parser
 ↓
request
 ↓
response
 ↓
bytes
 ↓
TCP
```

Itulah titik ketika kamu mulai berpindah dari sekadar **menggunakan
framework** menjadi memahami **bagaimana software bekerja di bawah
framework**.

------------------------------------------------------------------------

# Recommended Learning Order

``` text
01 Binary
02 Bit & Byte
03 ASCII
04 UTF-8
05 []byte
06 Buffer
07 io.Reader / io.Writer
08 File
09 stdin/stdout/stderr
10 Filesystem
11 Process
12 Pipe
13 Mini Shell
14 Concurrency
15 Binary Format
16 TCP
17 TCP Chat
18 HTTP over TCP
19 Syscall
20 Memory
21 unsafe
22 Rust
```

**Jangan terburu-buru ke Rust.**

Kalau bagian `byte → buffer → file → process → TCP` sudah benar-benar
dipahami, masuk Rust akan jauh lebih bermakna.
