# myfind

## Spec
```bash
myfind . -name "*.go"
myfind . -type d
myfind . -size +1M
```

## Constraints
- [ ] Parsing argumen sendiri dari `os.Args` (gaya `find` tidak cocok dengan package `flag`)
- [ ] Pattern pakai `filepath.Match`
- [ ] Traversal pakai `os.ReadDir` rekursif **dulu**, baru refactor ke `filepath.WalkDir`

## Acceptance Criteria
- [ ] `-name` dengan glob bekerja
- [ ] `-type f|d` bekerja
- [ ] `-size` bekerja (`+N`, `-N`, satuan `k`/`M`)
- [ ] Direktori tanpa permission → pesan ke stderr, traversal **lanjut**
- [ ] Output sama dengan `find` (WSL) untuk kasus yang sama (urutan boleh beda; bandingkan setelah `sort`)
- [ ] Test memakai `t.TempDir()` untuk membuat struktur direktori palsu

## Apa yang ada di bawahnya?
- Dari mana `os.ReadDir` mendapat daftar file? (syscall `getdents64` di Linux)
- Kenapa `DirEntry.Type()` lebih murah dari `os.Stat`?

## Observability
- WSL: `strace -c go run . / -name "*.go"` → syscall apa yang paling sering?

## Refleksi
- Bagian tersulit:
- Bug yang saya temui:
- Kalau mengulang, saya akan:

## Status: ⬜
