# Balanced Bracket Checker

## Analisis Kompleksitas

Algoritma ini menggunakan struktur data stack untuk menyimpan karakter bracket.

- **Iterasi string**: Saya melakukan satu kali iterasi sepanjang `n` karakter string → **O(n)**.
- **Operasi stack (push/pop)**: Setiap operasi push dan pop memakan waktu konstan → **O(1) per operasi**.
- **Total kompleksitas**: Karena setiap karakter hanya diproses sekali, kompleksitas total adalah **O(n)**, yang merupakan solusi paling optimal untuk permasalahan ini.

## Penjelasan Implementasi
1. Gunakan stack untuk menyimpan karakter bracket yang dibuka.
2. Saat menemui bracket tutup, periksa apakah sesuai dengan top dari stack.
3. Jika stack kosong di akhir iterasi, berarti bracket seimbang (`YES`), jika tidak, maka tidak seimbang (`NO`).

