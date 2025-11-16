# Sirekam Medis Pasien

Sirekam Medis Pasien adalah sebuah sistem manajemen rekam medis yang dirancang untuk membantu rumah sakit, klinik, atau fasilitas kesehatan lainnya dalam mengelola data pasien, jadwal, dan informasi medis secara efisien.

## Fitur Utama

- **Manajemen Pasien**: Menyimpan dan mengelola data pasien, termasuk informasi pribadi dan riwayat medis.
- **Jadwal dan Janji Temu**: Mengatur jadwal dokter dan janji temu pasien.
- **Manajemen Obat**: Melacak stok obat dan resep.
- **Rekam Medis Elektronik**: Menyimpan catatan medis pasien secara digital.
- **Laporan dan Statistik**: Menghasilkan laporan dan analisis data untuk pengambilan keputusan.

## Teknologi yang Digunakan

- **Backend**: Dibangun menggunakan bahasa pemrograman Go untuk memastikan performa tinggi dan skalabilitas.
- **Database**: Menggunakan MySQL untuk penyimpanan data yang andal.
- **Frontend**: (Opsional, tambahkan jika ada informasi terkait frontend).

## Cara Menjalankan Proyek

1. **Persiapan Lingkungan**:
   - Pastikan Go telah terinstal di sistem Anda.
   - Pastikan MySQL telah dikonfigurasi dengan benar.

2. **Menjalankan Backend**:

   ```bash
   cd backend
   go run cmd/server/main.go
   ```

3. **Pengujian**:

   - Jalankan skrip pengujian menggunakan PowerShell:

     ```powershell
     .\test.ps1
     ```

   - Atau gunakan perintah berikut untuk menjalankan semua pengujian:

     ```bash
     go test ./...
     ```

## Kontribusi

Kami menyambut kontribusi dari siapa pun yang tertarik untuk meningkatkan proyek ini. Silakan buat pull request atau ajukan issue untuk diskusi lebih lanjut.

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).