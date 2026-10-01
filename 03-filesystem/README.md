# 03 — Filesystem

ROADMAP: §10, §11

## Checklist konsep
- [ ] File vs directory
- [ ] Path absolut vs relatif; `filepath.Join` vs string concat (`\` vs `/`)
- [ ] Metadata: size, mode, permission, modtime
- [ ] Permission Unix (`rwxr-xr-x`, `0644`) vs Windows
- [ ] Traversal: `os.ReadDir` rekursif vs `filepath.WalkDir`
- [ ] Symlink

## Exercises
| Folder | Isi |
|---|---|
| `exercises/01-stat` | metadata file |
| `exercises/02-readdir` | traversal manual vs `WalkDir` |
| `exercises/03-file-ops` | create / rename / remove |

## Projects
- [myfind](myfind/)
- [mygrep](mygrep/)

## Pertanyaan evaluasi fase
1. Apa beda `filepath.Walk` dan `filepath.WalkDir`? Kenapa `WalkDir` lebih cepat?
2. Apa arti `0644`? Tulis dalam binary.
3. Kenapa sebaiknya tidak menyusun path dengan `dir + "/" + name`?
4. Apa yang terjadi saat traversal bertemu symlink yang menunjuk ke parent-nya sendiri?
5. Kenapa `os.Stat` pada file yang tidak ada mengembalikan error, dan bagaimana mengeceknya (`errors.Is(err, fs.ErrNotExist)`)?
