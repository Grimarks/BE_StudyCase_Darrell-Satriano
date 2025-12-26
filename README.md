Study Case 3

Simple RESTful API untuk platform pemesanan tiket event online. Dibangun menggunakan **Go (Golang)**, **Fiber**, **GORM**, dan **JWT** untuk otentikasi.

## Tech

* **Language:** Go 1.24
* **Framework:** Fiber v2
* **Database:** MySQL
* **ORM:** GORM
* **Auth:** JWT (JSON Web Token)

## cara run program

1. **Clone Repository**
```bash
git clone https://github.com/Grimarks/BE_StudyCase_Darrell-Satriano.git
cd StudyCase3

```


2. **config Database**
Pastikan kamu memiliki MySQL yang berjalan. Ubah konfigurasi database di file `config/database.go` jika diperlukan (username/password):
```go
// config/database.go
dsn := "root:password@tcp(localhost:3306)/ticket_db?charset=utf8mb4&parseTime=True&loc=Local"

```


3. **Install Dependencies**
```bash
go mod tidy

```


4. **Jalankan Server**
```bash
go run BE_Program_DarrellSatriano.go

```


Server akan berjalan di `http://localhost:8080`.

---

## Dokumentasi API Endpoints

### Authentication

#### 1. Register User

Mendaftarkan pengguna baru (Admin atau User).

* **URL:** `/register`
* **Method:** `POST`
* **Body (JSON):**
```json
{
    "name": "Darrell Satriano",
    "email": "darrell@example.com",
    "password": "password123",
    "role": "admin"  // atau "user"
}

```



#### 2. Login

Masuk untuk mendapatkan **JWT Token**. Token ini diperlukan untuk mengakses endpoint yang terkunci.

* **URL:** `/login`
* **Method:** `POST`
* **Body (JSON):**
```json
{
    "email": "darrell@example.com",
    "password": "password123"
}

```


* **Response Success:**
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

```



---

### Events

#### 3. Get All Events

Melihat daftar semua event yang tersedia.

* **URL:** `/events`
* **Method:** `GET`
* **Auth:** *Public* (Tidak perlu token)
* **Response Example:**
```json
[
    {
        "ID": 1,
        "Title": "Konser Musik A",
        "Capacity": 100,
        "TicketsSold": 10
    }
]

```



#### 4. Create Event (Admin Only)

Membuat event baru. Hanya bisa diakses oleh user dengan role `admin`.

* **URL:** `/events`
* **Method:** `POST`
* **Headers:**
* `Authorization`: `Bearer <token_jwt_admin>`


* **Body (JSON):**
```json
{
    "title": "Seminar Teknologi Go",
    "capacity": 50
}

```



---

### Transactions

#### 5. Buy Ticket

Membeli tiket untuk event tertentu. Sistem akan mengecek ketersediaan kapasitas (Overselling Protection).

* **URL:** `/transactions`
* **Method:** `POST`
* **Headers:**
* `Authorization`: `Bearer <token_jwt_user_atau_admin>`


* **Body (JSON):**
```json
{
    "event_id": 1,
    "quantity": 2
}

```


* **Response Success:**
Mengembalikan detail transaksi.
* **Response Error (Jika Penuh):**
```json
{
    "error": "Ticket sold out"
}

```



---

## Struktur Project

```
StudyCase3/
├── config/
│   ├── database.go   # Koneksi Database MySQL
│   └── jwt.go        # Secret Key JWT
├── controllers/
│   ├── auth_controller.go        # Login & Register Logic
│   ├── event_controller.go       # CRUD Event
│   └── transaction_controller.go # Pembelian Tiket
├── middleware/
│   ├── auth.go       # Validasi JWT Token
│   └── role.go       # Validasi Role Admin
├── models/
│   ├── event.go
│   ├── transaction.go
│   └── user.go
├── routes/
│   └── routes.go     # Definisi Endpoint Fiber
├── BE_Program_DarrellSatriano.go # Entry Point (Main)
└── go.mod

```
