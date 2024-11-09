
# Inventaris Management System

Sistem Manajemen Inventaris ini adalah aplikasi berbasis web yang dibangun dengan bahasa Go (Golang) menggunakan Chi sebagai router, dengan PostgreSQL sebagai database. Aplikasi ini menyediakan fitur manajemen pengguna, kategori, dan item dengan fitur tambahan untuk melihat investasi dan item yang membutuhkan penggantian.

## Fitur

- **Manajemen Pengguna**: Registrasi, login, dan logout pengguna.
- **Manajemen Kategori**: CRUD kategori untuk mengelompokkan item.
- **Manajemen Item**: CRUD item dengan informasi seperti harga, tanggal pembelian, kategori, dll.
- **Total Investasi**: Menghitung total investasi berdasarkan item yang terdaftar.
- **Peringatan Penggantian Barang**: Menampilkan daftar item yang membutuhkan penggantian.

## Teknologi yang Digunakan

- **Golang** dengan **Chi** untuk routing.
- **PostgreSQL** sebagai database utama.
- **Chi Middleware** untuk logging dan middleware tambahan.
  
## Struktur Proyek


## Instalasi dan Penggunaan

1. **Clone repositori ini**:

   ```sh
   git clone https://github.com/username/inventaris.git
   cd inventaris
   ```

2. **Setup Database**:

   - Buat database PostgreSQL baru, misalnya `inventaris_db`.
   - Jalankan migrasi jika diperlukan.

3. **Konfigurasi Database**:

   Pastikan konfigurasi database di file `database/config.go` sesuai dengan pengaturan PostgreSQL Anda.

4. **Jalankan Aplikasi**:

   ```sh
   go run main.go
   ```

5. **Akses API**:

   API akan tersedia di `http://localhost:8080`. Beberapa endpoint utama:
   - `/login` - Login pengguna
   - `/register` - Registrasi pengguna
   - `/api/categories` - CRUD untuk kategori
   - `/api/items` - CRUD untuk item
   - `/api/items/replacement-needed` - Daftar item yang perlu diganti
   - `/api/items/investment` - Total investasi dari item

## Endpoint API

### Autentikasi

- `POST /login` - Login pengguna.
- `POST /register` - Registrasi pengguna baru.
- `POST /logout` - Logout pengguna.

### Manajemen Kategori

- `GET /api/categories` - Mendapatkan semua kategori.
- `GET /api/categories/{id}` - Mendapatkan kategori berdasarkan ID.
- `POST /api/categories/add` - Menambah kategori baru.
- `PUT /api/categories/{id}` - Memperbarui kategori berdasarkan ID.
- `DELETE /api/categories/{id}` - Menghapus kategori berdasarkan ID.

### Manajemen Item

- `GET /api/items` - Mendapatkan semua item.
- `GET /api/items/{id}` - Mendapatkan item berdasarkan ID.
- `POST /api/items/add` - Menambah item baru.
- `PUT /api/items/{id}` - Memperbarui item berdasarkan ID.
- `DELETE /api/items/{id}` - Menghapus item berdasarkan ID.
- `GET /api/items/replacement-needed` - Mendapatkan daftar item yang memerlukan penggantian.

### Total Investasi

- `GET /api/items/investment` - Mendapatkan total investasi.
- `GET /api/items/investment/{id}` - Mendapatkan investasi per item berdasarkan ID.

## Middleware

- **Logger**: Chi Middleware yang menangani logging untuk setiap request.
- **Session Middleware** (opsional): Untuk autentikasi session (dapat ditambahkan di komentar `middleware_auth.SessionMiddleware`).
