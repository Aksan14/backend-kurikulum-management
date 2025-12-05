# Kurikulum Management System - Backend API# Kurikulum Management System - Backend API



Backend API untuk Sistem Manajemen Kurikulum menggunakan Golang, Gin Framework, dan MySQL.Backend API untuk Sistem Manajemen Kurikulum menggunakan Golang, Gin Framework, dan MySQL.



## 📋 Table of Contents## 📋 Table of Contents



- [Tech Stack](#tech-stack)- [Tech Stack](#tech-stack)

- [Getting Started](#getting-started)- [Getting Started](#getting-started)

- [API Documentation](#api-documentation)- [API Documentation](#api-documentation)

  - [Authentication](#1-authentication)  - [Authentication](#1-authentication)

  - [Users](#2-users-management)  - [Users](#2-users-management)

  - [CPL](#3-cpl-capaian-pembelajaran-lulusan)  - [CPL](#3-cpl-capaian-pembelajaran-lulusan)

  - [Mata Kuliah](#4-mata-kuliah)  - [Mata Kuliah](#4-mata-kuliah)

  - [CPL Assignment](#5-cpl-assignment)  - [CPL Assignment](#5-cpl-assignment)

  - [RPS](#6-rps-rencana-pembelajaran-semester)  - [RPS](#6-rps-rencana-pembelajaran-semester)

  - [Notifications](#7-notifications)  - [Notifications](#7-notifications)

  - [Dashboard](#8-dashboard)  - [Dashboard](#8-dashboard)

  - [Documents](#9-documents)  - [Documents](#9-documents)

  - [Files](#10-file-upload)  - [Files](#10-file-upload)

- [Environment Variables](#environment-variables)- [Environment Variables](#environment-variables)



------



## Tech Stack## Tech Stack



- **Language**: Go 1.23+- **Language**: Go 1.21+

- **Framework**: Gin (HTTP Router)- **Framework**: Gin (HTTP Router)

- **ORM**: GORM- **ORM**: GORM

- **Database**: MySQL 8.0+- **Database**: MySQL 8.0+

- **Authentication**: JWT (Access & Refresh Token)- **Authentication**: JWT (Access & Refresh Token)

- **Password Hashing**: bcrypt- **Password Hashing**: bcrypt



------



## Getting Started## Getting Started



### Installation### Installation



```bash```bash

# Clone repository# Clone repository

git clone <repository-url>git clone <repository-url>

cd backend-kurikulum-appscd backend-kurikulum-apps



# Copy environment file# Copy environment file

cp .env.example .envcp .env.example .env



# Install dependencies# Install dependencies

go mod downloadgo mod download



# Run the server# Create database

go run main.gomysql -u root -p -e "CREATE DATABASE kurikulum_db"

```mysql -u root -p kurikulum_db < migrations/001_init_schema.sql



### Running Tests# Run application

go run main.go

```bash```

# Run the test script

bash test_endpoints.sh---

```

## API Documentation

---

**Base URL:** `http://localhost:8080/api/v1`

## API Documentation

### Response Format

**Base URL**: `http://localhost:8080/api/v1`

**Success Response:**

### Response Format```json

{

Semua response API menggunakan format standar:  "success": true,

  "message": "Operation successful",

```json  "data": { ... }

{}

  "success": true,```

  "message": "Success message",

  "data": { ... }**Error Response:**

}```json

```{

  "success": false,

Error response:  "message": "Error message",

```json  "error": "Detailed error"

{}

  "success": false,```

  "message": "Error message",

  "error": "Detailed error"**Paginated Response:**

}```json

```{

  "success": true,

---  "message": "Data retrieved",

  "data": {

### 1. Authentication    "data": [ ... ],

    "page": 1,

#### Health Check    "limit": 10,

```    "total_items": 100,

GET /health    "total_pages": 10

```  }

}

**Response:**```

```json

{---

  "status": "ok",

  "message": "Kurikulum API is running"## 1. Authentication

}

```### 1.1 Register



---**Endpoint:** `POST /auth/register`



#### Register**Request:**

``````json

POST /auth/register{

```  "nama": "Dr. Budi Santoso, M.Kom",

  "email": "budi.santoso@university.ac.id",

**Request Body:**  "password": "password123",

```json  "role": "dosen",

{  "nip": "198501152010121002",

  "nama": "John Doe",  "phone": "081234567890"

  "email": "john@example.com",}

  "password": "Password123!",```

  "role": "dosen",

  "nip": "198501012020011001",| Field | Type | Required | Description |

  "phone": "081234567890"|-------|------|----------|-------------|

}| nama | string | ✅ | Min 2 characters |

```| email | string | ✅ | Valid email format |

| password | string | ✅ | Min 8 characters |

| Field | Type | Required | Description || role | string | ✅ | `kaprodi` or `dosen` |

|-------|------|----------|-------------|| nip | string | ❌ | Nomor Induk Pegawai |

| nama | string | Yes | Min 2 characters || phone | string | ❌ | Phone number |

| email | string | Yes | Valid email format |

| password | string | Yes | Min 8 characters |**Response (201):**

| role | string | Yes | `kaprodi` or `dosen` |```json

| nip | string | No | Nomor Induk Pegawai |{

| phone | string | No | Phone number |  "success": true,

  "message": "Registration successful",

**Response:**  "data": {

```json    "id": "550e8400-e29b-41d4-a716-446655440000",

{    "nama": "Dr. Budi Santoso, M.Kom",

  "success": true,    "email": "budi.santoso@university.ac.id",

  "message": "Registrasi berhasil",    "nip": "198501152010121002",

  "data": {    "role": "dosen",

    "id": "uuid",    "status": "active",

    "nama": "John Doe",    "phone": "081234567890",

    "email": "john@example.com",    "avatar_url": null,

    "role": "dosen",    "last_login": null,

    "status": "active",    "created_at": "2025-12-05T10:00:00Z"

    "created_at": "2025-12-05T00:00:00Z"  }

  }}

}```

```

---

---

### 1.2 Login

#### Login

```**Endpoint:** `POST /auth/login`

POST /auth/login

```**Request:**

```json

**Request Body:**{

```json  "email": "budi.santoso@university.ac.id",

{  "password": "password123"

  "email": "john@example.com",}

  "password": "Password123!"```

}

```**Response (200):**

```json

**Response:**{

```json  "success": true,

{  "message": "Login successful",

  "success": true,  "data": {

  "message": "Login berhasil",    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",

  "data": {    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",

    "access_token": "eyJhbGciOiJIUzI1NiIs...",    "expires_in": 86400,

    "refresh_token": "36acd783d7756d970958...",    "token_type": "Bearer",

    "expires_in": 3600,    "user": {

    "token_type": "Bearer",      "id": "550e8400-e29b-41d4-a716-446655440000",

    "user": {      "nama": "Dr. Budi Santoso, M.Kom",

      "id": "uuid",      "email": "budi.santoso@university.ac.id",

      "nama": "John Doe",      "nip": "198501152010121002",

      "email": "john@example.com",      "role": "dosen",

      "role": "dosen",      "status": "active",

      "status": "active"      "phone": "081234567890",

    }      "avatar_url": null,

  }      "last_login": "2025-12-05T10:00:00Z",

}      "created_at": "2025-12-05T09:00:00Z"

```    }

  }

---}

```

#### Refresh Token

```---

POST /auth/refresh

```### 1.3 Refresh Token



**Request Body:****Endpoint:** `POST /auth/refresh`

```json

{**Request:**

  "refresh_token": "36acd783d7756d970958..."```json

}{

```  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

}

**Response:**```

```json

{**Response (200):**

  "success": true,```json

  "message": "Token berhasil diperbarui",{

  "data": {  "success": true,

    "access_token": "eyJhbGciOiJIUzI1NiIs...",  "message": "Token refreshed",

    "refresh_token": "new_refresh_token...",  "data": {

    "expires_in": 3600,    "access_token": "new_access_token...",

    "token_type": "Bearer"    "refresh_token": "new_refresh_token...",

  }    "expires_in": 86400,

}    "token_type": "Bearer"

```  }

}

---```



#### Logout (Protected)---

```

POST /auth/logout### 1.4 Logout

Authorization: Bearer <access_token>

```**Endpoint:** `POST /auth/logout`



**Response:****Headers:**

```json```

{Authorization: Bearer <access_token>

  "success": true,```

  "message": "Logout berhasil"

}**Response (200):**

``````json

{

---  "success": true,

  "message": "Logout successful"

#### Get Profile (Protected)}

``````

GET /auth/profile

Authorization: Bearer <access_token>---

```

### 1.5 Get Profile

**Response:**

```json**Endpoint:** `GET /auth/profile`

{

  "success": true,**Headers:**

  "data": {```

    "id": "uuid",Authorization: Bearer <access_token>

    "nama": "John Doe",```

    "email": "john@example.com",

    "nip": "198501012020011001",**Response (200):**

    "role": "dosen",```json

    "status": "active",{

    "phone": "081234567890",  "success": true,

    "avatar_url": null,  "message": "Profile retrieved",

    "last_login": "2025-12-05T00:00:00Z",  "data": {

    "created_at": "2025-12-05T00:00:00Z"    "id": "550e8400-e29b-41d4-a716-446655440000",

  }    "nama": "Dr. Budi Santoso, M.Kom",

}    "email": "budi.santoso@university.ac.id",

```    "nip": "198501152010121002",

    "role": "dosen",

---    "status": "active",

    "phone": "081234567890",

#### Update Profile (Protected)    "avatar_url": null,

```    "last_login": "2025-12-05T10:00:00Z",

PUT /auth/profile    "created_at": "2025-12-05T09:00:00Z"

Authorization: Bearer <access_token>  }

```}

```

**Request Body:**

```json---

{

  "nama": "John Doe Updated",### 1.6 Update Profile

  "phone": "081234567891",

  "avatar_url": "https://example.com/avatar.jpg"**Endpoint:** `PUT /auth/profile`

}

```**Headers:**

```

---Authorization: Bearer <access_token>

```

#### Change Password (Protected)

```**Request:**

POST /auth/change-password```json

Authorization: Bearer <access_token>{

```  "nama": "Dr. Budi Santoso, M.Kom (Updated)",

  "phone": "081234567899",

**Request Body:**  "avatar_url": "https://example.com/avatar.jpg"

```json}

{```

  "old_password": "OldPassword123!",

  "new_password": "NewPassword123!"**Response (200):**

}```json

```{

  "success": true,

---  "message": "Profile updated",

  "data": {

### 2. Users Management    "id": "550e8400-e29b-41d4-a716-446655440000",

    "nama": "Dr. Budi Santoso, M.Kom (Updated)",

> **Note**: Create, Update, Delete operations require `kaprodi` role.    "email": "budi.santoso@university.ac.id",

    "nip": "198501152010121002",

#### Get All Users (Protected)    "role": "dosen",

```    "status": "active",

GET /users    "phone": "081234567899",

Authorization: Bearer <access_token>    "avatar_url": "https://example.com/avatar.jpg",

```    "last_login": "2025-12-05T10:00:00Z",

    "created_at": "2025-12-05T09:00:00Z"

**Query Parameters:**  }

| Parameter | Type | Description |}

|-----------|------|-------------|```

| page | int | Page number (default: 1) |

| limit | int | Items per page (default: 10, max: 100) |---

| search | string | Search by name or email |

| role | string | Filter by role: `kaprodi`, `dosen` |### 1.7 Change Password

| status | string | Filter by status: `active`, `inactive` |

| sort_by | string | Sort field: `nama`, `email`, `created_at` |**Endpoint:** `POST /auth/change-password`

| sort_order | string | `asc` or `desc` |

**Headers:**

---```

Authorization: Bearer <access_token>

#### Get Dosen List (Protected)```

```

GET /users/dosen**Request:**

Authorization: Bearer <access_token>```json

```{

  "old_password": "password123",

Returns list of active dosen users.  "new_password": "newpassword456"

}

---```



#### Get User By ID (Protected)**Response (200):**

``````json

GET /users/:id{

Authorization: Bearer <access_token>  "success": true,

```  "message": "Password changed successfully"

}

---```



#### Create User (Kaprodi Only)---

```

POST /users## 2. Users Management

Authorization: Bearer <access_token>

```> **Note:** Requires **Kaprodi** role for most endpoints.



**Request Body:**### 2.1 Get All Users

```json

{**Endpoint:** `GET /users`

  "nama": "New User",

  "email": "newuser@example.com",**Headers:**

  "password": "Password123!",```

  "role": "dosen",Authorization: Bearer <access_token>

  "nip": "198501012020011001",```

  "phone": "081234567890"

}**Query Parameters:**

```| Parameter | Type | Description |

|-----------|------|-------------|

---| page | int | Page number (default: 1) |

| limit | int | Items per page (default: 10, max: 100) |

#### Update User (Kaprodi Only)| search | string | Search by nama, email, nip |

```| role | string | Filter: `kaprodi`, `dosen` |

PUT /users/:id| status | string | Filter: `active`, `inactive` |

Authorization: Bearer <access_token>| sort_by | string | Sort: `nama`, `email`, `created_at` |

```| sort_order | string | Order: `asc`, `desc` |



**Request Body:****Example:** `GET /users?page=1&limit=10&role=dosen&status=active`

```json

{**Response (200):**

  "nama": "Updated Name",```json

  "email": "updated@example.com",{

  "role": "dosen",  "success": true,

  "status": "active"  "message": "Users retrieved",

}  "data": {

```    "data": [

      {

---        "id": "550e8400-e29b-41d4-a716-446655440000",

        "nama": "Dr. Budi Santoso, M.Kom",

#### Delete User (Kaprodi Only)        "email": "budi.santoso@university.ac.id",

```        "nip": "198501152010121002",

DELETE /users/:id        "role": "dosen",

Authorization: Bearer <access_token>        "status": "active",

```        "phone": "081234567890",

        "avatar_url": null,

---        "last_login": "2025-12-05T10:00:00Z",

        "created_at": "2025-12-05T09:00:00Z"

#### Toggle User Status (Kaprodi Only)      }

```    ],

PATCH /users/:id/toggle-status    "page": 1,

Authorization: Bearer <access_token>    "limit": 10,

```    "total_items": 50,

    "total_pages": 5

---  }

}

### 3. CPL (Capaian Pembelajaran Lulusan)```



> **Note**: Create, Update, Delete operations require `kaprodi` role.---



#### Get All CPL (Protected)### 2.2 Get User by ID

```

GET /cpl**Endpoint:** `GET /users/:id`

Authorization: Bearer <access_token>

```**Response (200):**

```json

**Query Parameters:**{

| Parameter | Type | Description |  "success": true,

|-----------|------|-------------|  "message": "User retrieved",

| page | int | Page number (default: 1) |  "data": {

| limit | int | Items per page (default: 10, max: 100) |    "id": "550e8400-e29b-41d4-a716-446655440000",

| search | string | Search by kode or nama |    "nama": "Dr. Budi Santoso, M.Kom",

| status | string | Filter: `draft`, `published`, `archived` |    "email": "budi.santoso@university.ac.id",

| sort_by | string | Sort field: `kode`, `nama`, `created_at` |    "nip": "198501152010121002",

| sort_order | string | `asc` or `desc` |    "role": "dosen",

    "status": "active",

---    "phone": "081234567890",

    "avatar_url": null,

#### Get CPL By ID (Protected)    "last_login": "2025-12-05T10:00:00Z",

```    "created_at": "2025-12-05T09:00:00Z"

GET /cpl/:id  }

Authorization: Bearer <access_token>}

``````



------



#### Get Active CPL (Protected)### 2.3 Get All Dosen

```

GET /cpl/active**Endpoint:** `GET /users/dosen`

Authorization: Bearer <access_token>

```**Response (200):**

```json

Returns CPL with status `published`.{

  "success": true,

---  "message": "Dosen list retrieved",

  "data": {

#### Get CPL Statistics (Protected)    "data": [

```      {

GET /cpl/statistics        "id": "550e8400-e29b-41d4-a716-446655440000",

Authorization: Bearer <access_token>        "nama": "Dr. Budi Santoso, M.Kom",

```        "email": "budi.santoso@university.ac.id",

        "nip": "198501152010121002",

**Response:**        "role": "dosen",

```json        "status": "active"

{      }

  "success": true,    ],

  "data": {    "page": 1,

    "total_cpl": 10,    "limit": 100,

    "published": 5,    "total_items": 45,

    "draft": 3,    "total_pages": 1

    "archived": 2  }

  }}

}```

```

---

---

### 2.4 Create User (Kaprodi Only)

#### Create CPL (Kaprodi Only)

```**Endpoint:** `POST /users`

POST /cpl

Authorization: Bearer <access_token>**Request:**

``````json

{

**Request Body:**  "nama": "Dr. Ani Wijaya, M.T",

```json  "email": "ani.wijaya@university.ac.id",

{  "password": "password123",

  "kode": "CPL-001",  "role": "dosen",

  "nama": "Capaian Pembelajaran Lulusan 1",  "nip": "197812102005011003",

  "deskripsi": "Deskripsi lengkap CPL",  "phone": "081234567899"

  "status": "draft"}

}```

```

**Response (201):**

| Field | Type | Required | Description |```json

|-------|------|----------|-------------|{

| kode | string | Yes | Min 2, Max 20 characters |  "success": true,

| nama | string | Yes | Min 5, Max 255 characters |  "message": "User created",

| deskripsi | string | No | Full description |  "data": {

| status | string | No | `draft`, `published`, `archived` (default: `draft`) |    "id": "660e8400-e29b-41d4-a716-446655440001",

    "nama": "Dr. Ani Wijaya, M.T",

---    "email": "ani.wijaya@university.ac.id",

    "nip": "197812102005011003",

#### Update CPL (Kaprodi Only)    "role": "dosen",

```    "status": "active",

PUT /cpl/:id    "phone": "081234567899",

Authorization: Bearer <access_token>    "avatar_url": null,

```    "last_login": null,

    "created_at": "2025-12-05T10:00:00Z"

**Request Body:**  }

```json}

{```

  "kode": "CPL-001",

  "nama": "Capaian Pembelajaran Lulusan 1 Updated",---

  "deskripsi": "Deskripsi updated",

  "status": "published"### 2.5 Update User (Kaprodi Only)

}

```**Endpoint:** `PUT /users/:id`



---**Request:**

```json

#### Update CPL Status (Kaprodi Only){

```  "nama": "Dr. Ani Wijaya, M.T (Updated)",

PATCH /cpl/:id/status  "email": "ani.wijaya@university.ac.id",

Authorization: Bearer <access_token>  "role": "dosen",

```  "nip": "197812102005011003",

  "phone": "081234567800",

**Request Body:**  "status": "active"

```json}

{```

  "status": "published"

}**Response (200):**

``````json

{

---  "success": true,

  "message": "User updated",

#### Delete CPL (Kaprodi Only)  "data": {

```    "id": "660e8400-e29b-41d4-a716-446655440001",

DELETE /cpl/:id    "nama": "Dr. Ani Wijaya, M.T (Updated)",

Authorization: Bearer <access_token>    "email": "ani.wijaya@university.ac.id",

```    "nip": "197812102005011003",

    "role": "dosen",

---    "status": "active",

    "phone": "081234567800",

### 4. Mata Kuliah    "avatar_url": null,

    "last_login": null,

> **Note**: Create, Update, Delete operations require `kaprodi` role.    "created_at": "2025-12-05T10:00:00Z"

  }

#### Get All Mata Kuliah (Protected)}

``````

GET /mata-kuliah

Authorization: Bearer <access_token>---

```

### 2.6 Delete User (Kaprodi Only)

**Query Parameters:**

| Parameter | Type | Description |**Endpoint:** `DELETE /users/:id`

|-----------|------|-------------|

| page | int | Page number (default: 1) |**Response (200):**

| limit | int | Items per page (default: 10, max: 100) |```json

| search | string | Search by kode or nama |{

| semester | int | Filter by semester (1-8) |  "success": true,

| is_active | bool | Filter by active status |  "message": "User deleted"

| sort_by | string | Sort: `kode`, `nama`, `sks`, `semester`, `created_at` |}

| sort_order | string | `asc` or `desc` |```



------



#### Get My Mata Kuliah (Protected)### 2.7 Toggle User Status (Kaprodi Only)

```

GET /mata-kuliah/my**Endpoint:** `PATCH /users/:id/toggle-status`

Authorization: Bearer <access_token>

```**Response (200):**

```json

Returns mata kuliah assigned to current user.{

  "success": true,

---  "message": "User status toggled",

  "data": {

#### Get Mata Kuliah By Semester (Protected)    "id": "660e8400-e29b-41d4-a716-446655440001",

```    "nama": "Dr. Ani Wijaya, M.T",

GET /mata-kuliah/semester/:semester    "status": "inactive"

Authorization: Bearer <access_token>  }

```}

```

---

---

#### Get Mata Kuliah By Dosen (Protected)

```## 3. CPL (Capaian Pembelajaran Lulusan)

GET /mata-kuliah/dosen/:dosen_id

Authorization: Bearer <access_token>### 3.1 Get All CPL

```

**Endpoint:** `GET /cpl`

---

**Query Parameters:**

#### Get Mata Kuliah By ID (Protected)| Parameter | Type | Description |

```|-----------|------|-------------|

GET /mata-kuliah/:id| page | int | Page number |

Authorization: Bearer <access_token>| limit | int | Items per page |

```| search | string | Search by kode, judul |

| status | string | `draft`, `published`, `archived` |

---| aspek | string | `sikap`, `pengetahuan`, `keterampilan_umum`, `keterampilan_khusus` |

| kategori | string | Filter by kategori |

#### Create Mata Kuliah (Kaprodi Only)| sort_by | string | `kode`, `judul`, `created_at` |

```| sort_order | string | `asc`, `desc` |

POST /mata-kuliah

Authorization: Bearer <access_token>**Response (200):**

``````json

{

**Request Body:**  "success": true,

```json  "message": "CPL retrieved",

{  "data": {

  "kode": "MK-001",    "data": [

  "nama": "Pemrograman Web",      {

  "sks": 3,        "id": "cpl-uuid-001",

  "semester": 3,        "kode": "CPL-01",

  "deskripsi": "Deskripsi mata kuliah",        "judul": "Kemampuan Berpikir Kritis",

  "prasyarat": ["MK-DASAR-001", "MK-DASAR-002"],        "deskripsi": "Mampu menerapkan pemikiran logis, kritis, sistematis...",

  "dosen_pengampu_id": "uuid",        "aspek": "pengetahuan",

  "koordinator_id": "uuid",        "kategori": "Kompetensi Utama",

  "is_active": true        "status": "published",

}        "version": 1,

```        "created_at": "2025-12-05T09:00:00Z",

        "updated_at": "2025-12-05T09:00:00Z",

| Field | Type | Required | Description |        "created_by": "user-uuid",

|-------|------|----------|-------------|        "published_at": "2025-12-05T10:00:00Z"

| kode | string | Yes | Min 2, Max 20 characters |      }

| nama | string | Yes | Min 3, Max 255 characters |    ],

| sks | int | Yes | 1-12 |    "page": 1,

| semester | int | Yes | 1-8 |    "limit": 10,

| deskripsi | string | No | Full description |    "total_items": 10,

| prasyarat | array | No | List of prerequisite course codes |    "total_pages": 1

| dosen_pengampu_id | string | No | UUID of assigned dosen |  }

| koordinator_id | string | No | UUID of coordinator |}

| is_active | bool | No | Active status (default: true) |```



------



#### Update Mata Kuliah (Kaprodi Only)### 3.2 Get CPL by ID

```

PUT /mata-kuliah/:id**Endpoint:** `GET /cpl/:id`

Authorization: Bearer <access_token>

```**Response (200):**

```json

---{

  "success": true,

#### Toggle Mata Kuliah Status (Kaprodi Only)  "message": "CPL retrieved",

```  "data": {

PATCH /mata-kuliah/:id/toggle-status    "id": "cpl-uuid-001",

Authorization: Bearer <access_token>    "kode": "CPL-01",

```    "judul": "Kemampuan Berpikir Kritis",

    "deskripsi": "Mampu menerapkan pemikiran logis, kritis, sistematis...",

---    "aspek": "pengetahuan",

    "kategori": "Kompetensi Utama",

#### Delete Mata Kuliah (Kaprodi Only)    "status": "published",

```    "version": 1,

DELETE /mata-kuliah/:id    "created_at": "2025-12-05T09:00:00Z",

Authorization: Bearer <access_token>    "updated_at": "2025-12-05T09:00:00Z",

```    "created_by": "user-uuid",

    "published_at": "2025-12-05T10:00:00Z",

---    "creator": {

      "id": "user-uuid",

### 5. CPL Assignment      "nama": "Kaprodi",

      "email": "kaprodi@university.ac.id"

> **Note**: Create, Delete operations require `kaprodi` role.    }

  }

#### Get All Assignments (Protected)}

``````

GET /cpl-assignments

Authorization: Bearer <access_token>---

```

### 3.3 Get Active CPL

**Query Parameters:**

| Parameter | Type | Description |**Endpoint:** `GET /cpl/active`

|-----------|------|-------------|

| page | int | Page number |**Response (200):** Same as Get All CPL with status=published

| limit | int | Items per page |

| cpl_id | string | Filter by CPL ID |---

| dosen_id | string | Filter by Dosen ID |

| status | string | `assigned`, `accepted`, `rejected`, `done`, `cancelled` |### 3.4 Get CPL Statistics

| sort_by | string | `assigned_at`, `deadline`, `status` |

| sort_order | string | `asc` or `desc` |**Endpoint:** `GET /cpl/statistics`



---**Response (200):**

```json

#### Get My Assignments (Protected){

```  "success": true,

GET /cpl-assignments/my  "message": "CPL statistics retrieved",

Authorization: Bearer <access_token>  "data": {

```    "total_cpl": 10,

    "published": 7,

Returns assignments for current user.    "draft": 2,

    "archived": 1

---  }

}

#### Get Assignments By CPL (Protected)```

```

GET /cpl-assignments/cpl/:cpl_id---

Authorization: Bearer <access_token>

```### 3.5 Create CPL (Kaprodi Only)



---**Endpoint:** `POST /cpl`



#### Get Assignment By ID (Protected)**Request:**

``````json

GET /cpl-assignments/:id{

Authorization: Bearer <access_token>  "kode": "CPL-01",

```  "judul": "Kemampuan Berpikir Kritis",

  "deskripsi": "Mampu menerapkan pemikiran logis, kritis, sistematis, dan inovatif dalam konteks pengembangan ilmu pengetahuan.",

---  "aspek": "pengetahuan",

  "kategori": "Kompetensi Utama",

#### Create Assignment (Kaprodi Only)  "status": "draft"

```}

POST /cpl-assignments```

Authorization: Bearer <access_token>

```| Field | Type | Required | Description |

|-------|------|----------|-------------|

**Request Body:**| kode | string | ✅ | 2-20 characters |

```json| judul | string | ✅ | Min 5 characters |

{| deskripsi | string | ✅ | Min 10 characters |

  "cpl_id": "uuid",| aspek | string | ✅ | `sikap`, `pengetahuan`, `keterampilan_umum`, `keterampilan_khusus` |

  "dosen_id": "uuid",| kategori | string | ✅ | Category |

  "mata_kuliah": "Pemrograman Web",| status | string | ❌ | `draft`, `published`, `archived` (default: draft) |

  "mata_kuliah_id": "uuid",

  "deadline": "2025-12-31T23:59:59Z",**Response (201):**

  "comment": "Silakan dikerjakan"```json

}{

```  "success": true,

  "message": "CPL created",

---  "data": {

    "id": "cpl-uuid-001",

#### Update Assignment Status (Protected)    "kode": "CPL-01",

```    "judul": "Kemampuan Berpikir Kritis",

PATCH /cpl-assignments/:id/status    "deskripsi": "Mampu menerapkan pemikiran logis...",

Authorization: Bearer <access_token>    "aspek": "pengetahuan",

```    "kategori": "Kompetensi Utama",

    "status": "draft",

**Request Body:**    "version": 1,

```json    "created_at": "2025-12-05T09:00:00Z",

{    "updated_at": "2025-12-05T09:00:00Z",

  "status": "accepted",    "created_by": "user-uuid"

  "comment": "Saya terima tugas ini"  }

}}

``````



Status values: `accepted`, `rejected`, `done`, `cancelled`---



---### 3.6 Update CPL (Kaprodi Only)



#### Delete Assignment (Kaprodi Only)**Endpoint:** `PUT /cpl/:id`

```

DELETE /cpl-assignments/:id**Request:**

Authorization: Bearer <access_token>```json

```{

  "kode": "CPL-01",

---  "judul": "Kemampuan Berpikir Kritis (Updated)",

  "deskripsi": "Updated description...",

### 6. RPS (Rencana Pembelajaran Semester)  "aspek": "pengetahuan",

  "kategori": "Kompetensi Utama",

#### Get All RPS (Protected)  "status": "published"

```}

GET /rps```

Authorization: Bearer <access_token>

```---



**Query Parameters:**### 3.7 Delete CPL (Kaprodi Only)

| Parameter | Type | Description |

|-----------|------|-------------|**Endpoint:** `DELETE /cpl/:id`

| page | int | Page number |

| limit | int | Items per page |**Response (200):**

| search | string | Search term |```json

| mata_kuliah_id | string | Filter by Mata Kuliah ID |{

| dosen_id | string | Filter by Dosen ID |  "success": true,

| status | string | `draft`, `submitted`, `approved`, `rejected`, `published` |  "message": "CPL deleted"

| tahun_akademik | string | Filter by academic year |}

| semester | int | Filter by semester |```

| sort_by | string | `mata_kuliah_nama`, `created_at`, `updated_at` |

| sort_order | string | `asc` or `desc` |---



---### 3.8 Update CPL Status (Kaprodi Only)



#### Get My RPS (Protected)**Endpoint:** `PATCH /cpl/:id/status`

```

GET /rps/my**Request:**

Authorization: Bearer <access_token>```json

```{

  "status": "published"

---}

```

#### Get RPS By Mata Kuliah (Protected)

```**Response (200):**

GET /rps/mata-kuliah/:mata_kuliah_id```json

Authorization: Bearer <access_token>{

```  "success": true,

  "message": "CPL status updated",

---  "data": {

    "id": "cpl-uuid-001",

#### Get RPS By ID (Protected)    "status": "published",

```    "published_at": "2025-12-05T10:00:00Z"

GET /rps/:id  }

Authorization: Bearer <access_token>}

``````



------



#### Create RPS (Protected)## 4. Mata Kuliah

```

POST /rps### 4.1 Get All Mata Kuliah

Authorization: Bearer <access_token>

```**Endpoint:** `GET /mata-kuliah`



**Request Body:****Query Parameters:**

```json| Parameter | Type | Description |

{|-----------|------|-------------|

  "mata_kuliah_id": "uuid",| page | int | Page number |

  "tahun_akademik": "2024/2025",| limit | int | Items per page |

  "deskripsi": "Deskripsi RPS",| search | string | Search by kode, nama |

  "tujuan": "Tujuan pembelajaran",| semester | int | Filter by semester (1-8) |

  "metode": ["Ceramah", "Diskusi", "Praktikum"],| jenis | string | `wajib`, `pilihan` |

  "bobot_nilai": {| status | string | `aktif`, `nonaktif`, `dihapus` |

    "tugas": 20,| sort_by | string | `kode`, `nama`, `sks`, `semester`, `created_at` |

    "uts": 25,| sort_order | string | `asc`, `desc` |

    "uas": 35,

    "kehadiran": 10,**Response (200):**

    "praktikum": 10```json

  }{

}  "success": true,

```  "message": "Mata kuliah retrieved",

  "data": {

---    "data": [

      {

#### Update RPS (Protected)        "id": "mk-uuid-001",

```        "kode": "IF101",

PUT /rps/:id        "nama": "Algoritma dan Pemrograman",

Authorization: Bearer <access_token>        "sks": 3,

```        "semester": 1,

        "jenis": "wajib",

---        "deskripsi": "Mata kuliah dasar pemrograman...",

        "prasyarat": [],

#### Submit RPS (Protected)        "status": "aktif",

```        "created_at": "2025-12-05T09:00:00Z",

PATCH /rps/:id/submit        "updated_at": "2025-12-05T09:00:00Z",

Authorization: Bearer <access_token>        "created_by": "user-uuid"

```      }

    ],

---    "page": 1,

    "limit": 10,

#### Approve RPS (Kaprodi Only)    "total_items": 30,

```    "total_pages": 3

PATCH /rps/:id/approve  }

Authorization: Bearer <access_token>}

``````



**Request Body:**---

```json

{### 4.2 Get Mata Kuliah by ID

  "review_notes": "RPS sudah sesuai standar"

}**Endpoint:** `GET /mata-kuliah/:id`

```

---

---

### 4.3 Get My Mata Kuliah (Dosen)

#### Reject RPS (Kaprodi Only)

```**Endpoint:** `GET /mata-kuliah/my`

PATCH /rps/:id/reject

Authorization: Bearer <access_token>**Description:** Get mata kuliah assigned to current logged-in dosen.

```

---

**Request Body:**

```json### 4.4 Get Mata Kuliah by Semester

{

  "review_notes": "Perlu perbaikan pada bagian..."**Endpoint:** `GET /mata-kuliah/semester/:semester`

}

```**Example:** `GET /mata-kuliah/semester/3`



------



#### Request Revision (Kaprodi Only)### 4.5 Get Mata Kuliah by Dosen

```

PATCH /rps/:id/request-revision**Endpoint:** `GET /mata-kuliah/dosen/:dosen_id`

Authorization: Bearer <access_token>

```---



---### 4.6 Create Mata Kuliah (Kaprodi Only)



#### Delete RPS (Protected)**Endpoint:** `POST /mata-kuliah`

```

DELETE /rps/:id**Request:**

Authorization: Bearer <access_token>```json

```{

  "kode": "IF101",

---  "nama": "Algoritma dan Pemrograman",

  "sks": 3,

### RPS Sub-Resources  "semester": 1,

  "deskripsi": "Mata kuliah ini membahas konsep dasar algoritma...",

#### CPMK (Capaian Pembelajaran Mata Kuliah)  "prasyarat": [],

  "dosen_pengampu_id": "dosen-uuid",

**Add CPMK:**  "koordinator_id": "dosen-uuid"

```}

POST /rps/:rps_id/cpmk```

Authorization: Bearer <access_token>

```| Field | Type | Required | Description |

|-------|------|----------|-------------|

**Request Body:**| kode | string | ✅ | 2-20 characters |

```json| nama | string | ✅ | 3-255 characters |

{| sks | int | ✅ | 1-12 |

  "kode": "CPMK-1",| semester | int | ✅ | 1-8 |

  "deskripsi": "Deskripsi CPMK",| deskripsi | string | ❌ | Description |

  "cpl_ids": ["uuid1", "uuid2"],| prasyarat | array | ❌ | Prerequisites |

  "urutan": 1| dosen_pengampu_id | string | ❌ | Dosen UUID |

}| koordinator_id | string | ❌ | Koordinator UUID |

```

**Response (201):**

**Update CPMK:**```json

```{

PUT /rps/cpmk/:cpmk_id  "success": true,

Authorization: Bearer <access_token>  "message": "Mata kuliah created",

```  "data": {

    "id": "mk-uuid-001",

**Delete CPMK:**    "kode": "IF101",

```    "nama": "Algoritma dan Pemrograman",

DELETE /rps/cpmk/:cpmk_id    "sks": 3,

Authorization: Bearer <access_token>    "semester": 1,

```    "jenis": "wajib",

    "deskripsi": "Mata kuliah ini membahas...",

---    "prasyarat": [],

    "status": "aktif",

#### Rencana Pembelajaran    "created_at": "2025-12-05T09:00:00Z"

  }

**Add Rencana Pembelajaran:**}

``````

POST /rps/:rps_id/rencana-pembelajaran

Authorization: Bearer <access_token>---

```

### 4.7 Update Mata Kuliah (Kaprodi Only)

**Request Body:**

```json**Endpoint:** `PUT /mata-kuliah/:id`

{

  "pertemuan": 1,**Request:**

  "topik": "Pengenalan Web Development",```json

  "sub_topik": ["HTML", "CSS", "JavaScript"],{

  "metode": "Ceramah dan Praktikum",  "kode": "IF101",

  "waktu": 150,  "nama": "Algoritma dan Pemrograman Dasar",

  "cpmk_ids": ["uuid1"],  "sks": 4,

  "materi": "Link materi"  "semester": 1,

}  "deskripsi": "Updated description...",

```  "prasyarat": [],

  "dosen_pengampu_id": "new-dosen-uuid",

**Update Rencana Pembelajaran:**  "is_active": true

```}

PUT /rps/rencana-pembelajaran/:rencana_id```

Authorization: Bearer <access_token>

```---



**Delete Rencana Pembelajaran:**### 4.8 Delete Mata Kuliah (Kaprodi Only)

```

DELETE /rps/rencana-pembelajaran/:rencana_id**Endpoint:** `DELETE /mata-kuliah/:id`

Authorization: Bearer <access_token>

```---



---### 4.9 Toggle Mata Kuliah Status



#### Bahan Bacaan**Endpoint:** `PATCH /mata-kuliah/:id/toggle-status`



**Add Bahan Bacaan:**---

```

POST /rps/:rps_id/bahan-bacaan## 5. CPL Assignment

Authorization: Bearer <access_token>

```### 5.1 Get All Assignments



**Request Body:****Endpoint:** `GET /cpl-assignments`

```json

{**Query Parameters:**

  "judul": "Web Development with Go",| Parameter | Type | Description |

  "penulis": "John Doe",|-----------|------|-------------|

  "tahun": 2024,| page | int | Page number |

  "jenis": "buku",| limit | int | Items per page |

  "url": "https://example.com/book",| cpl_id | string | Filter by CPL |

  "isbn": "978-3-16-148410-0",| dosen_id | string | Filter by Dosen |

  "halaman": "1-100",| status | string | `assigned`, `accepted`, `rejected`, `done`, `cancelled` |

  "urutan": 1| sort_by | string | `assigned_at`, `deadline`, `status` |

}| sort_order | string | `asc`, `desc` |

```

**Response (200):**

Jenis: `buku`, `jurnal`, `artikel`, `website`, `modul````json

{

**Update Bahan Bacaan:**  "success": true,

```  "message": "Assignments retrieved",

PUT /rps/bahan-bacaan/:bahan_id  "data": {

Authorization: Bearer <access_token>    "data": [

```      {

        "id": "assignment-uuid",

**Delete Bahan Bacaan:**        "cpl_id": "cpl-uuid",

```        "dosen_id": "dosen-uuid",

DELETE /rps/bahan-bacaan/:bahan_id        "mata_kuliah": "IF101 - Algoritma",

Authorization: Bearer <access_token>        "mata_kuliah_id": "mk-uuid",

```        "status": "assigned",

        "assigned_at": "2025-12-05T09:00:00Z",

---        "accepted_at": null,

        "completed_at": null,

#### Evaluasi        "deadline": "2025-12-31T23:59:59Z",

        "comment": "Please complete this assignment",

**Add Evaluasi:**        "rejection_reason": null,

```        "created_by": "kaprodi-uuid",

POST /rps/:rps_id/evaluasi        "cpl": {

Authorization: Bearer <access_token>          "id": "cpl-uuid",

```          "kode": "CPL-01",

          "judul": "Kemampuan Berpikir Kritis"

**Request Body:**        },

```json        "dosen": {

{          "id": "dosen-uuid",

  "jenis": "UTS",          "nama": "Dr. Budi Santoso"

  "bobot": 25,        }

  "deskripsi": "Ujian Tengah Semester",      }

  "minggu_pelaksanaan": [8],    ],

  "kriteria_penilaian": "Kriteria penilaian",    "page": 1,

  "rubrik_penilaian": "Rubrik penilaian"    "limit": 10,

}    "total_items": 20,

```    "total_pages": 2

  }

**Update Evaluasi:**}

``````

PUT /rps/evaluasi/:evaluasi_id

Authorization: Bearer <access_token>---

```

### 5.2 Get Assignment by ID

**Delete Evaluasi:**

```**Endpoint:** `GET /cpl-assignments/:id`

DELETE /rps/evaluasi/:evaluasi_id

Authorization: Bearer <access_token>---

```

### 5.3 Get My Assignments (Dosen)

---

**Endpoint:** `GET /cpl-assignments/my`

### 7. Notifications

---

#### Get My Notifications (Protected)

```### 5.4 Get Assignments by CPL

GET /notifications

Authorization: Bearer <access_token>**Endpoint:** `GET /cpl-assignments/cpl/:cpl_id`

```

---

**Query Parameters:**

| Parameter | Type | Description |### 5.5 Create Assignment (Kaprodi Only)

|-----------|------|-------------|

| page | int | Page number |**Endpoint:** `POST /cpl-assignments`

| limit | int | Items per page |

| type | string | `assignment`, `approval`, `rejection`, `document`, `info`, `deadline`, `system` |**Request:**

| is_read | bool | Filter by read status |```json

| sort_order | string | `asc` or `desc` |{

  "cpl_id": "cpl-uuid",

---  "dosen_id": "dosen-uuid",

  "mata_kuliah": "IF101 - Algoritma dan Pemrograman",

#### Get Unread Count (Protected)  "mata_kuliah_id": "mk-uuid",

```  "deadline": "2025-12-31T23:59:59Z",

GET /notifications/unread-count  "comment": "Please implement this CPL in your course"

Authorization: Bearer <access_token>}

``````



**Response:**| Field | Type | Required | Description |

```json|-------|------|----------|-------------|

{| cpl_id | string | ✅ | CPL UUID |

  "success": true,| dosen_id | string | ✅ | Dosen UUID |

  "data": {| mata_kuliah | string | ❌ | Mata kuliah name |

    "count": 5| mata_kuliah_id | string | ❌ | Mata kuliah UUID |

  }| deadline | datetime | ❌ | Deadline |

}| comment | string | ❌ | Comment |

```

**Response (201):**

---```json

{

#### Create Notification (Kaprodi Only)  "success": true,

```  "message": "Assignment created",

POST /notifications  "data": {

Authorization: Bearer <access_token>    "id": "assignment-uuid",

```    "cpl_id": "cpl-uuid",

    "dosen_id": "dosen-uuid",

**Request Body:**    "mata_kuliah": "IF101 - Algoritma",

```json    "status": "assigned",

{    "assigned_at": "2025-12-05T09:00:00Z",

  "user_id": "uuid",    "deadline": "2025-12-31T23:59:59Z"

  "title": "New Assignment",  }

  "message": "You have a new CPL assignment",}

  "type": "assignment",```

  "action_url": "/cpl-assignments/123",

  "related_id": "uuid",---

  "related_type": "cpl_assignment",

  "priority": "high"### 5.6 Update Assignment Status

}

```**Endpoint:** `PATCH /cpl-assignments/:id/status`



Type values: `assignment`, `approval`, `rejection`, `document`, `info`, `deadline`, `system`**Request:**

Priority values: `low`, `normal`, `high`, `urgent````json

{

---  "status": "accepted",

  "comment": "I accept this assignment"

#### Mark as Read (Protected)}

``````

PATCH /notifications/:id/read

Authorization: Bearer <access_token>**Or Reject:**

``````json

{

---  "status": "rejected",

  "rejection_reason": "Cannot complete due to schedule conflict"

#### Mark All as Read (Protected)}

``````

PATCH /notifications/read-all

Authorization: Bearer <access_token>| Field | Type | Required | Description |

```|-------|------|----------|-------------|

| status | string | ✅ | `accepted`, `rejected`, `completed`, `done` |

---| comment | string | ❌ | Comment |

| rejection_reason | string | ❌ | Required if rejecting |

#### Delete Notification (Protected)

```---

DELETE /notifications/:id

Authorization: Bearer <access_token>### 5.7 Delete Assignment (Kaprodi Only)

```

**Endpoint:** `DELETE /cpl-assignments/:id`

---

---

### 8. Dashboard

## 6. RPS (Rencana Pembelajaran Semester)

#### Get My Dashboard (Protected)

```### 6.1 Get All RPS

GET /dashboard

Authorization: Bearer <access_token>**Endpoint:** `GET /rps`

```

**Query Parameters:**

Returns dashboard data based on user role.| Parameter | Type | Description |

|-----------|------|-------------|

---| page | int | Page number |

| limit | int | Items per page |

#### Get Kaprodi Dashboard (Kaprodi Only)| search | string | Search |

```| mata_kuliah_id | string | Filter by mata kuliah |

GET /dashboard/kaprodi| dosen_id | string | Filter by dosen |

Authorization: Bearer <access_token>| status | string | `draft`, `submitted`, `approved`, `rejected`, `published` |

```| tahun_akademik | string | Filter by year |

| semester | int | Filter by semester |

**Response:**

```json**Response (200):**

{```json

  "success": true,{

  "data": {  "success": true,

    "total_cpl": 10,  "message": "RPS retrieved",

    "published_cpl": 5,  "data": {

    "draft_cpl": 5,    "data": [

    "total_rps": 20,      {

    "approved_rps": 15,        "id": "rps-uuid",

    "submitted_rps": 3,        "mata_kuliah_id": "mk-uuid",

    "rejected_rps": 2,        "mata_kuliah_nama": "Algoritma dan Pemrograman",

    "active_dosen": 10,        "kode_mk": "IF101",

    "accepted_assignments": 5,        "sks": 3,

    "pending_assignments": 3,        "semester": 1,

    "completed_assignments": 2,        "tahun_akademik": "2024/2025",

    "ready_documents": 8        "dosen_id": "dosen-uuid",

  }        "dosen_nama": "Dr. Budi Santoso",

}        "deskripsi": "Deskripsi mata kuliah...",

```        "tujuan": "Tujuan pembelajaran...",

        "metode": ["Ceramah", "Diskusi", "Praktikum"],

---        "bobot_nilai": {

          "tugas": 20,

#### Get Dosen Dashboard (Protected)          "uts": 25,

```          "uas": 35,

GET /dashboard/dosen          "kehadiran": 10,

Authorization: Bearer <access_token>          "praktikum": 10

```        },

        "status": "draft",

---        "created_at": "2025-12-05T09:00:00Z",

        "updated_at": "2025-12-05T09:00:00Z"

### 9. Documents      }

    ],

#### Get All Documents (Protected)    "page": 1,

```    "limit": 10,

GET /documents    "total_items": 50,

Authorization: Bearer <access_token>    "total_pages": 5

```  }

}

**Query Parameters:**```

| Parameter | Type | Description |

|-----------|------|-------------|---

| page | int | Page number |

| limit | int | Items per page |### 6.2 Get RPS by ID (Detail)

| template_id | string | Filter by template |

| status | string | `processing`, `ready`, `failed`, `archived` |**Endpoint:** `GET /rps/:id`

| tahun | string | Filter by year |

| file_type | string | `docx`, `pdf`, `xlsx` |**Response (200):**

| sort_by | string | `created_at`, `completed_at` |```json

| sort_order | string | `asc` or `desc` |{

  "success": true,

---  "message": "RPS retrieved",

  "data": {

#### Get Document By ID (Protected)    "id": "rps-uuid",

```    "mata_kuliah_id": "mk-uuid",

GET /documents/:id    "mata_kuliah_nama": "Algoritma dan Pemrograman",

Authorization: Bearer <access_token>    "kode_mk": "IF101",

```    "sks": 3,

    "semester": 1,

---    "tahun_akademik": "2024/2025",

    "dosen_id": "dosen-uuid",

#### Generate Document (Protected)    "dosen_nama": "Dr. Budi Santoso",

```    "deskripsi": "...",

POST /documents/generate    "tujuan": "...",

Authorization: Bearer <access_token>    "metode": ["Ceramah", "Diskusi"],

```    "bobot_nilai": { ... },

    "status": "draft",

**Request Body:**    "cpmk": [

```json      {

{        "id": "cpmk-uuid",

  "template_id": "uuid",        "rps_id": "rps-uuid",

  "tahun": "2024",        "kode": "CPMK-01",

  "sections": ["section1", "section2"],        "deskripsi": "Mahasiswa mampu...",

  "file_type": "pdf",        "cpl_ids": ["cpl-uuid"],

  "generation_data": {        "urutan": 1

    "key": "value"      }

  }    ],

}    "rencana_pembelajaran": [

```      {

        "id": "rp-uuid",

---        "rps_id": "rps-uuid",

        "pertemuan": 1,

#### Delete Document (Protected)        "topik": "Pengantar Algoritma",

```        "sub_topik": ["Definisi", "Karakteristik"],

DELETE /documents/:id        "metode": "Ceramah",

Authorization: Bearer <access_token>        "waktu": 150

```      }

    ],

---    "bahan_bacaan": [ ... ],

    "evaluasi": [ ... ]

### Document Templates  }

}

#### Get All Templates (Protected)```

```

GET /documents/templates---

Authorization: Bearer <access_token>

```### 6.3 Get My RPS (Dosen)



---**Endpoint:** `GET /rps/my`



#### Get Template By ID (Protected)---

```

GET /documents/templates/:id### 6.4 Get RPS by Mata Kuliah

Authorization: Bearer <access_token>

```**Endpoint:** `GET /rps/mata-kuliah/:mata_kuliah_id`



------



#### Create Template (Kaprodi Only)### 6.5 Create RPS

```

POST /documents/templates**Endpoint:** `POST /rps`

Authorization: Bearer <access_token>

```**Request:**

```json

**Request Body:**{

```json  "mata_kuliah_id": "mk-uuid",

{  "tahun_ajaran": "2024/2025",

  "nama": "Template RPS",  "semester_type": "ganjil",

  "deskripsi": "Template untuk RPS",  "deskripsi_mk": "Mata kuliah ini membahas...",

  "sections": ["identitas", "cpl", "cpmk", "rencana_pembelajaran"],  "capaian_pembelajaran": "Setelah menyelesaikan...",

  "file_url": "https://example.com/template.docx",  "metode_pembelajaran": ["Ceramah", "Diskusi", "Praktikum"],

  "version": "1.0",  "media_pembelajaran": ["Slide", "Video", "Laboratorium"]

  "is_active": true}

}```

```

| Field | Type | Required | Description |

---|-------|------|----------|-------------|

| mata_kuliah_id | string | ✅ | Mata kuliah UUID |

#### Update Template (Kaprodi Only)| tahun_ajaran | string | ✅ | e.g., "2024/2025" |

```| semester_type | string | ✅ | `ganjil`, `genap` |

PUT /documents/templates/:id| deskripsi_mk | string | ❌ | Description |

Authorization: Bearer <access_token>| capaian_pembelajaran | string | ❌ | Learning outcomes |

```| metode_pembelajaran | array | ❌ | Methods |

| media_pembelajaran | array | ❌ | Media |

---

---

#### Delete Template (Kaprodi Only)

```### 6.6 Update RPS

DELETE /documents/templates/:id

Authorization: Bearer <access_token>**Endpoint:** `PUT /rps/:id`

```

**Request:**

---```json

{

### 10. File Upload  "tahun_ajaran": "2024/2025",

  "semester_type": "ganjil",

#### Upload File (Protected)  "deskripsi_mk": "Updated description...",

```  "capaian_pembelajaran": "Updated learning outcomes...",

POST /files/upload  "metode_pembelajaran": ["Ceramah", "Diskusi"],

Authorization: Bearer <access_token>  "media_pembelajaran": ["Slide", "Video"]

Content-Type: multipart/form-data}

``````



**Form Data:**---

| Field | Type | Description |

|-------|------|-------------|### 6.7 Delete RPS

| file | file | File to upload |

| type | string | Optional: file type category |**Endpoint:** `DELETE /rps/:id`



**Response:**---

```json

{### 6.8 RPS Workflow

  "success": true,

  "message": "File uploaded successfully",#### Submit RPS for Review

  "data": {

    "filename": "document.pdf",**Endpoint:** `PATCH /rps/:id/submit`

    "url": "/uploads/document.pdf",

    "size": 102400,**Description:** Changes status from `draft` to `submitted`.

    "mime_type": "application/pdf"

  }**Response (200):**

}```json

```{

  "success": true,

---  "message": "RPS submitted for review",

  "data": {

#### Get File Info (Protected)    "id": "rps-uuid",

```    "status": "submitted",

GET /files/info?path=/uploads/document.pdf    "submitted_at": "2025-12-05T10:00:00Z"

Authorization: Bearer <access_token>  }

```}

```

---

---

#### Delete File (Protected)

```#### Approve RPS (Kaprodi Only)

DELETE /files?path=/uploads/document.pdf

Authorization: Bearer <access_token>**Endpoint:** `PATCH /rps/:id/approve`

```

**Request:**

---```json

{

## Environment Variables  "catatan": "RPS sudah lengkap dan sesuai kurikulum. Disetujui."

}

Create a `.env` file with the following variables:```



```env**Response (200):**

# Server```json

PORT=8080{

GIN_MODE=release  "success": true,

  "message": "RPS approved",

# Database  "data": {

DB_HOST=localhost    "id": "rps-uuid",

DB_PORT=3306    "status": "approved",

DB_USER=root    "reviewed_at": "2025-12-05T11:00:00Z",

DB_PASSWORD=your_password    "review_notes": "RPS sudah lengkap..."

DB_NAME=akademik-management  }

}

# JWT```

JWT_SECRET=your-super-secret-key

JWT_ACCESS_EXPIRY=1h---

JWT_REFRESH_EXPIRY=168h

#### Reject RPS (Kaprodi Only)

# CORS

ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173**Endpoint:** `PATCH /rps/:id/reject`



# Upload**Request:**

UPLOAD_DIR=./uploads```json

MAX_UPLOAD_SIZE=10485760{

```  "alasan": "Evaluasi belum lengkap, mohon tambahkan rubrik penilaian."

}

---```



## Error Codes---



| HTTP Status | Description |#### Request Revision (Kaprodi Only)

|-------------|-------------|

| 200 | Success |**Endpoint:** `PATCH /rps/:id/request-revision`

| 201 | Created |

| 400 | Bad Request - Invalid input |**Request:**

| 401 | Unauthorized - Invalid or missing token |```json

| 403 | Forbidden - Insufficient permissions |{

| 404 | Not Found - Resource not found |  "catatan": "Mohon revisi bagian CPMK agar lebih spesifik dan terukur."

| 409 | Conflict - Resource already exists |}

| 500 | Internal Server Error |```



------



## Rate Limiting### 6.9 CPMK (Capaian Pembelajaran Mata Kuliah)



API memiliki rate limit **100 requests per minute** per IP address.#### Get CPMK by RPS



---**Endpoint:** `GET /rps/:id/cpmk`



## Authentication**Response (200):**

```json

Semua endpoint yang memerlukan authentication harus menyertakan header:{

  "success": true,

```  "data": [

Authorization: Bearer <access_token>    {

```      "id": "cpmk-uuid",

      "rps_id": "rps-uuid",

Access token memiliki masa berlaku 1 jam. Gunakan refresh token untuk mendapatkan access token baru.      "kode": "CPMK-01",

      "deskripsi": "Mahasiswa mampu memahami konsep dasar algoritma",

---      "cpl_ids": ["cpl-uuid-01"],

      "urutan": 1,

## Role-Based Access Control      "created_at": "2025-12-05T09:00:00Z"

    }

| Role | Description |  ]

|------|-------------|}

| `kaprodi` | Full access to all resources |```

| `dosen` | Limited access - can manage own RPS and view assigned CPL |

---

### Kaprodi Only Endpoints:

- User management (CUD)#### Create CPMK

- CPL management (CUD)

- Mata Kuliah management (CUD)**Endpoint:** `POST /rps/:id/cpmk`

- CPL Assignment (Create, Delete)

- RPS approval/rejection**Request:**

- Document template management (CUD)```json

- Create notifications{

  "kode": "CPMK-01",

---  "deskripsi": "Mahasiswa mampu memahami dan menjelaskan konsep dasar algoritma",

  "bobot": 25.0,

## License  "urutan": 1

}

MIT License```



## Author---



Backend Kurikulum Apps Team#### Update CPMK


**Endpoint:** `PUT /rps/:rps_id/cpmk/:cpmk_id`

**Request:**
```json
{
  "kode": "CPMK-01",
  "deskripsi": "Updated description...",
  "bobot": 30.0,
  "urutan": 1
}
```

---

#### Delete CPMK

**Endpoint:** `DELETE /rps/:rps_id/cpmk/:cpmk_id`

---

### 6.10 Rencana Pembelajaran

#### Get Rencana Pembelajaran

**Endpoint:** `GET /rps/:id/rencana-pembelajaran`

**Response (200):**
```json
{
  "success": true,
  "data": [
    {
      "id": "rp-uuid",
      "rps_id": "rps-uuid",
      "pertemuan": 1,
      "kemampuan_akhir": "Mahasiswa memahami konsep algoritma",
      "indikator": "Dapat menjelaskan definisi algoritma",
      "materi": "Pengantar Algoritma",
      "metode_pembelajaran": "Ceramah, Diskusi",
      "waktu_menit": 150,
      "pengalaman_belajar": "Mahasiswa mengerjakan latihan...",
      "kriteria_penilaian": "Quiz",
      "bobot_nilai": 5.0,
      "referensi": "Buku: Introduction to Algorithms"
    }
  ]
}
```

---

#### Create Rencana Pembelajaran

**Endpoint:** `POST /rps/:id/rencana-pembelajaran`

**Request:**
```json
{
  "pertemuan": 1,
  "kemampuan_akhir": "Mahasiswa memahami konsep dasar algoritma",
  "indikator": "Dapat menjelaskan definisi dan karakteristik algoritma",
  "materi": "Pengantar Algoritma",
  "metode_pembelajaran": "Ceramah, Diskusi",
  "waktu_menit": 150,
  "pengalaman_belajar": "Mahasiswa mengerjakan latihan pseudocode",
  "kriteria_penilaian": "Quiz",
  "bobot_nilai": 5.0,
  "referensi": "Cormen, Introduction to Algorithms"
}
```

---

#### Update Rencana Pembelajaran

**Endpoint:** `PUT /rps/:rps_id/rencana-pembelajaran/:rp_id`

---

#### Delete Rencana Pembelajaran

**Endpoint:** `DELETE /rps/:rps_id/rencana-pembelajaran/:rp_id`

---

### 6.11 Bahan Bacaan

#### Get Bahan Bacaan

**Endpoint:** `GET /rps/:id/bahan-bacaan`

---

#### Create Bahan Bacaan

**Endpoint:** `POST /rps/:id/bahan-bacaan`

**Request:**
```json
{
  "jenis": "utama",
  "judul": "Introduction to Algorithms",
  "penulis": "Thomas H. Cormen",
  "penerbit": "MIT Press",
  "tahun": 2022,
  "isbn": "978-0262046305",
  "url": "https://mitpress.mit.edu/books/introduction-algorithms",
  "urutan": 1
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| jenis | string | ✅ | `utama`, `pendukung` |
| judul | string | ✅ | Title |
| penulis | string | ❌ | Author |
| penerbit | string | ❌ | Publisher |
| tahun | int | ❌ | Year |
| isbn | string | ❌ | ISBN |
| url | string | ❌ | URL |
| urutan | int | ❌ | Order |

---

#### Update Bahan Bacaan

**Endpoint:** `PUT /rps/:rps_id/bahan-bacaan/:bb_id`

---

#### Delete Bahan Bacaan

**Endpoint:** `DELETE /rps/:rps_id/bahan-bacaan/:bb_id`

---

### 6.12 Evaluasi

#### Get Evaluasi

**Endpoint:** `GET /rps/:id/evaluasi`

**Response (200):**
```json
{
  "success": true,
  "data": [
    {
      "id": "eval-uuid",
      "rps_id": "rps-uuid",
      "komponen": "UTS",
      "teknik_penilaian": "Tes Tertulis",
      "instrumen": "Soal Essay",
      "bobot": 25.0,
      "kriteria_penilaian": "Ketepatan jawaban 40%, Analisis 30%, Penulisan 30%",
      "urutan": 1
    }
  ]
}
```

---

#### Create Evaluasi

**Endpoint:** `POST /rps/:id/evaluasi`

**Request:**
```json
{
  "komponen": "UTS",
  "teknik_penilaian": "Tes Tertulis",
  "instrumen": "Soal Essay dan Pilihan Ganda",
  "bobot": 25.0,
  "kriteria_penilaian": "Ketepatan jawaban 40%, Analisis 30%, Penulisan 30%",
  "urutan": 1
}
```

---

#### Update Evaluasi

**Endpoint:** `PUT /rps/:rps_id/evaluasi/:eval_id`

---

#### Delete Evaluasi

**Endpoint:** `DELETE /rps/:rps_id/evaluasi/:eval_id`

---

## 7. Notifications

### 7.1 Get My Notifications

**Endpoint:** `GET /notifications`

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page | int | Page number |
| limit | int | Items per page |
| type | string | `assignment`, `approval`, `rejection`, `document`, `info`, `deadline`, `system` |
| is_read | bool | Filter by read status |
| sort_order | string | `asc`, `desc` |

**Response (200):**
```json
{
  "success": true,
  "data": {
    "data": [
      {
        "id": "notif-uuid",
        "user_id": "user-uuid",
        "title": "RPS Approved",
        "message": "Your RPS for Algoritma dan Pemrograman has been approved.",
        "type": "approval",
        "is_read": false,
        "action_url": "/rps/rps-uuid",
        "related_id": "rps-uuid",
        "related_type": "rps",
        "priority": "normal",
        "created_at": "2025-12-05T10:00:00Z",
        "read_at": null
      }
    ],
    "page": 1,
    "limit": 10,
    "total_items": 5,
    "total_pages": 1
  }
}
```

---

### 7.2 Get Unread Count

**Endpoint:** `GET /notifications/unread-count`

**Response (200):**
```json
{
  "success": true,
  "data": {
    "count": 5
  }
}
```

---

### 7.3 Get Notification by ID

**Endpoint:** `GET /notifications/:id`

---

### 7.4 Create Notification (Kaprodi Only)

**Endpoint:** `POST /notifications`

**Request:**
```json
{
  "user_id": "dosen-uuid",
  "title": "New CPL Assignment",
  "message": "You have been assigned to implement CPL-01",
  "type": "assignment",
  "action_url": "/cpl-assignments/assignment-uuid",
  "related_id": "assignment-uuid",
  "related_type": "cpl_assignment",
  "priority": "high"
}
```

---

### 7.5 Mark as Read

**Endpoint:** `PATCH /notifications/:id/read`

**Response (200):**
```json
{
  "success": true,
  "message": "Notification marked as read"
}
```

---

### 7.6 Mark All as Read

**Endpoint:** `PATCH /notifications/read-all`

---

### 7.7 Delete Notification

**Endpoint:** `DELETE /notifications/:id`

---

## 8. Dashboard

### 8.1 Get Dashboard (Role-based)

**Endpoint:** `GET /dashboard`

**Response for Kaprodi (200):**
```json
{
  "success": true,
  "data": {
    "total_cpl": 10,
    "published_cpl": 7,
    "draft_cpl": 3,
    "total_rps": 50,
    "approved_rps": 40,
    "pending_review": 5,
    "rejected_rps": 5,
    "active_dosen": 45,
    "active_assignments": 30,
    "completed_assignments": 100,
    "documents_generated": 25
  }
}
```

**Response for Dosen (200):**
```json
{
  "success": true,
  "data": {
    "total_assignments": 10,
    "accepted_assignments": 8,
    "pending_assignments": 2,
    "completed_assignments": 6,
    "total_rps": 5,
    "approved_rps": 3,
    "draft_rps": 1,
    "submitted_rps": 1,
    "rejected_rps": 0
  }
}
```

---

### 8.2 Get Kaprodi Dashboard

**Endpoint:** `GET /dashboard/kaprodi`

---

### 8.3 Get Dosen Dashboard

**Endpoint:** `GET /dashboard/dosen`

---

## 9. Documents

### 9.1 Get All Documents

**Endpoint:** `GET /documents`

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page | int | Page number |
| limit | int | Items per page |
| template_id | string | Filter by template |
| status | string | `processing`, `ready`, `failed`, `archived` |
| tahun | string | Filter by year |
| file_type | string | `docx`, `pdf`, `xlsx` |

---

### 9.2 Get Document by ID

**Endpoint:** `GET /documents/:id`

---

### 9.3 Generate Document

**Endpoint:** `POST /documents/generate`

**Request:**
```json
{
  "template_id": "template-uuid",
  "tahun": "2024/2025",
  "sections": ["cover", "content", "appendix"],
  "file_type": "pdf",
  "generation_data": {
    "rps_id": "rps-uuid"
  }
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "Document generation started",
  "data": {
    "id": "doc-uuid",
    "template_id": "template-uuid",
    "template_name": "Template RPS",
    "tahun": "2024/2025",
    "status": "processing",
    "file_type": "pdf",
    "progress": 0,
    "created_at": "2025-12-05T10:00:00Z"
  }
}
```

---

### 9.4 Download Document

**Endpoint:** `GET /documents/:id/download`

---

### 9.5 Delete Document

**Endpoint:** `DELETE /documents/:id`

---

## 10. File Upload

### 10.1 Upload File

**Endpoint:** `POST /files/upload`

**Headers:**
```
Authorization: Bearer <access_token>
Content-Type: multipart/form-data
```

**Form Data:**
| Field | Type | Description |
|-------|------|-------------|
| file | file | The file to upload |

**Response (201):**
```json
{
  "success": true,
  "message": "File uploaded",
  "data": {
    "id": "file-uuid",
    "original_name": "document.pdf",
    "stored_name": "uuid-document.pdf",
    "file_path": "/uploads/uuid-document.pdf",
    "file_size": 1024000,
    "mime_type": "application/pdf",
    "file_type": "pdf",
    "url": "/uploads/uuid-document.pdf",
    "created_at": "2025-12-05T10:00:00Z"
  }
}
```

---

### 10.2 Get File

**Endpoint:** `GET /files/:filename`

**Description:** Serves the uploaded file.

---

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| APP_NAME | Application name | Kurikulum Management System |
| APP_ENV | Environment | development |
| APP_PORT | Server port | 8080 |
| DB_HOST | Database host | localhost |
| DB_PORT | Database port | 3306 |
| DB_NAME | Database name | kurikulum_db |
| DB_USER | Database user | root |
| DB_PASSWORD | Database password | |
| JWT_SECRET | JWT secret key | your-secret-key |
| JWT_EXPIRES_IN | Access token expiry | 24h |
| JWT_REFRESH_EXPIRES_IN | Refresh token expiry | 168h |
| UPLOAD_PATH | Upload directory | ./uploads |
| MAX_FILE_SIZE | Max upload size | 10485760 |
| CORS_ALLOWED_ORIGINS | CORS origins | http://localhost:3000 |

---

## Default Credentials

| Field | Value |
|-------|-------|
| Email | kaprodi@university.ac.id |
| Password | password123 |
| Role | kaprodi |

---

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 409 | Conflict |
| 422 | Validation Error |
| 500 | Internal Server Error |

---

## License

MIT License

---

## Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin (HTTP Router)
- **ORM**: GORM
- **Database**: MySQL 8.0+
- **Authentication**: JWT (Access & Refresh Token)
- **Password Hashing**: bcrypt
- **UUID**: google/uuid

## Features

- 🔐 **Authentication & Authorization**
  - JWT-based authentication with refresh tokens
  - Role-based access control (Kaprodi & Dosen)
  - Password change functionality

- 📚 **CPL Management**
  - CRUD Capaian Pembelajaran Lulusan
  - Status management (draft, published, archived)
  - CPL assignment to Dosen

- 📖 **Mata Kuliah Management**
  - CRUD Mata Kuliah
  - Assign Dosen Pengampu
  - Filter by semester

- 📝 **RPS Management**
  - Full RPS document management
  - CPMK (Capaian Pembelajaran Mata Kuliah)
  - Rencana Pembelajaran (Weekly Plan)
  - Bahan Bacaan (References)
  - Evaluasi (Assessments)
  - Approval workflow (Submit, Approve, Reject, Request Revision)

- 🔔 **Notifications**
  - Real-time notifications
  - Mark as read/unread
  - Filter by type

- 📄 **Document Generation**
  - Template management
  - Document generation from templates

## Project Structure

```
backend-kurikulum-apps/
├── config/             # Configuration files (database, jwt)
├── controller/         # HTTP handlers
├── dto/                # Data Transfer Objects (Request/Response)
├── helper/             # Helper functions
├── middleware/         # HTTP middlewares (auth, cors, logger, rate_limiter, recovery)
├── model/              # Database models (GORM)
├── repository/         # Database repositories (data access layer)
├── routes/             # Route definitions
├── service/            # Business logic layer
├── uploads/            # File uploads directory
├── util/               # Utility functions
├── migrations/         # SQL migrations
├── main.go             # Entry point
├── go.mod              # Go modules
├── Makefile            # Build commands
├── .env.example        # Environment variables template
└── .air.toml           # Hot reload config
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- MySQL 8.0 or higher
- Make (optional, for Makefile commands)

### Installation

1. Clone the repository
```bash
git clone <repository-url>
cd backend-kurikulum-apps
```

2. Copy environment file
```bash
cp .env.example .env
```

3. Configure your `.env` file with your database credentials and JWT secret

4. Install dependencies
```bash
go mod download
```

5. Create database and run migrations
```bash
mysql -u root -p -e "CREATE DATABASE kurikulum_db"
mysql -u root -p kurikulum_db < migrations/001_init_schema.sql
```

6. Run the application
```bash
go run main.go
```

Or use Makefile:
```bash
make init    # Initialize project (create .env, download deps, create folders)
make run     # Build and run
make dev     # Run with hot reload (requires air)
```

### Using Air for Hot Reload

Install Air:
```bash
go install github.com/cosmtrek/air@latest
```

Run with hot reload:
```bash
air
```

---

## API Documentation

Base URL: `http://localhost:8080/api/v1`

### Response Format

Semua response menggunakan format JSON standar:

**Success Response:**
```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

**Error Response:**
```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error description"
}
```

**Paginated Response:**
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "items": [ ... ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 100,
      "total_pages": 10
    }
  }
}
```

---

## 1. Authentication

### 1.1 Login

**Endpoint:** `POST /auth/login`

**Request Body:**
```json
{
  "email": "kaprodi@university.ac.id",
  "password": "password123"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "user": {
      "id": "uuid",
      "nip": "1234567890",
      "name": "Dr. John Doe",
      "email": "kaprodi@university.ac.id",
      "role": "kaprodi"
    },
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_at": "2024-01-16T10:00:00Z"
  }
}
```

---

### 1.2 Register

**Endpoint:** `POST /auth/register`

**Request Body:**
```json
{
  "nip": "1234567890",
  "name": "Dr. Jane Doe",
  "email": "jane.doe@university.ac.id",
  "password": "securepassword123",
  "role": "dosen"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Registration successful",
  "data": {
    "user": {
      "id": "uuid",
      "nip": "1234567890",
      "name": "Dr. Jane Doe",
      "email": "jane.doe@university.ac.id",
      "role": "dosen"
    },
    "access_token": "...",
    "refresh_token": "...",
    "expires_at": "2024-01-16T10:00:00Z"
  }
}
```

---

### 1.3 Refresh Token

**Endpoint:** `POST /auth/refresh`

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response:**
```json
{
  "success": true,
  "message": "Token refreshed successfully",
  "data": {
    "access_token": "new_access_token...",
    "refresh_token": "new_refresh_token...",
    "expires_at": "2024-01-16T10:00:00Z"
  }
}
```

---

### 1.4 Logout

**Endpoint:** `POST /auth/logout`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Response:**
```json
{
  "success": true,
  "message": "Logout successful"
}
```

---

### 1.5 Get Profile

**Endpoint:** `GET /auth/profile`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Response:**
```json
{
  "success": true,
  "message": "Profile retrieved successfully",
  "data": {
    "id": "uuid",
    "nip": "1234567890",
    "name": "Dr. John Doe",
    "email": "kaprodi@university.ac.id",
    "role": "kaprodi",
    "is_active": true,
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

---

### 1.6 Update Profile

**Endpoint:** `PUT /auth/profile`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "name": "Dr. John Doe Updated",
  "email": "newemail@university.ac.id"
}
```

---

### 1.7 Change Password

**Endpoint:** `POST /auth/change-password`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "old_password": "currentpassword",
  "new_password": "newsecurepassword123"
}
```

---

## 2. Users Management

> **Note:** Most user management endpoints require **Kaprodi** role.

### 2.1 Get All Users

**Endpoint:** `GET /users`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Query Parameters:**
| Parameter | Type | Description | Default |
|-----------|------|-------------|---------|
| page | int | Page number | 1 |
| limit | int | Items per page | 10 |
| search | string | Search by name/email/nip | - |
| role | string | Filter by role (kaprodi/dosen) | - |

**Example:** `GET /users?page=1&limit=10&search=john&role=dosen`

**Response:**
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": {
    "items": [
      {
        "id": "uuid",
        "nip": "1234567890",
        "name": "Dr. John Doe",
        "email": "john@university.ac.id",
        "role": "dosen",
        "is_active": true,
        "created_at": "2024-01-15T10:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 50,
      "total_pages": 5
    }
  }
}
```

---

### 2.2 Get User by ID

**Endpoint:** `GET /users/:id`

**Headers:**
```
Authorization: Bearer <access_token>
```

---

### 2.3 Get All Dosen

**Endpoint:** `GET /users/dosen`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Response:**
```json
{
  "success": true,
  "message": "Dosen list retrieved successfully",
  "data": [
    {
      "id": "uuid",
      "nip": "1234567890",
      "name": "Dr. John Doe",
      "email": "john@university.ac.id"
    }
  ]
}
```

---

### 2.4 Create User

**Endpoint:** `POST /users`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "nip": "0987654321",
  "name": "Dr. New User",
  "email": "newuser@university.ac.id",
  "password": "password123",
  "role": "dosen"
}
```

---

### 2.5 Update User

**Endpoint:** `PUT /users/:id`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "name": "Dr. Updated Name",
  "email": "updated@university.ac.id",
  "role": "dosen"
}
```

---

### 2.6 Delete User

**Endpoint:** `DELETE /users/:id`

**Headers:**
```
Authorization: Bearer <access_token>
```

---

### 2.7 Toggle User Status

**Endpoint:** `PATCH /users/:id/toggle-status`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Description:** Activates or deactivates a user account.

---

## 3. CPL (Capaian Pembelajaran Lulusan)

### 3.1 Get All CPL

**Endpoint:** `GET /cpl`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Query Parameters:**
| Parameter | Type | Description | Default |
|-----------|------|-------------|---------|
| page | int | Page number | 1 |
| limit | int | Items per page | 10 |
| status | string | Filter by status (draft/published/archived) | - |
| search | string | Search by kode/deskripsi | - |

**Response:**
```json
{
  "success": true,
  "message": "CPL retrieved successfully",
  "data": {
    "items": [
      {
        "id": "uuid",
        "kode": "CPL-01",
        "deskripsi": "Mampu menerapkan pemikiran logis, kritis, sistematis...",
        "status": "published",
        "created_at": "2024-01-15T10:00:00Z"
      }
    ],
    "pagination": { ... }
  }
}
```

---

### 3.2 Get CPL by ID

**Endpoint:** `GET /cpl/:id`

---

### 3.3 Get Active CPL

**Endpoint:** `GET /cpl/active`

**Description:** Returns all CPL with status "published".

---

### 3.4 Get CPL Statistics

**Endpoint:** `GET /cpl/statistics`

**Response:**
```json
{
  "success": true,
  "message": "CPL statistics retrieved successfully",
  "data": {
    "total": 10,
    "draft": 2,
    "published": 7,
    "archived": 1
  }
}
```

---

### 3.5 Create CPL (Kaprodi Only)

**Endpoint:** `POST /cpl`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "kode": "CPL-01",
  "deskripsi": "Mampu menerapkan pemikiran logis, kritis, sistematis, dan inovatif dalam konteks pengembangan atau implementasi ilmu pengetahuan dan teknologi yang memperhatikan dan menerapkan nilai humaniora yang sesuai dengan bidang keahliannya.",
  "status": "draft"
}
```

---

### 3.6 Update CPL (Kaprodi Only)

**Endpoint:** `PUT /cpl/:id`

**Request Body:**
```json
{
  "kode": "CPL-01-Updated",
  "deskripsi": "Updated description...",
  "status": "published"
}
```

---

### 3.7 Delete CPL (Kaprodi Only)

**Endpoint:** `DELETE /cpl/:id`

---

### 3.8 Update CPL Status

**Endpoint:** `PATCH /cpl/:id/status`

**Request Body:**
```json
{
  "status": "published"
}
```

**Valid Status Values:** `draft`, `published`, `archived`

---

## 4. Mata Kuliah

### 4.1 Get All Mata Kuliah

**Endpoint:** `GET /mata-kuliah`

**Query Parameters:**
| Parameter | Type | Description | Default |
|-----------|------|-------------|---------|
| page | int | Page number | 1 |
| limit | int | Items per page | 10 |
| semester | int | Filter by semester (1-8) | - |
| search | string | Search by kode/nama | - |

**Response:**
```json
{
  "success": true,
  "message": "Mata kuliah retrieved successfully",
  "data": {
    "items": [
      {
        "id": "uuid",
        "kode": "IF101",
        "nama": "Algoritma dan Pemrograman",
        "sks": 3,
        "semester": 1,
        "deskripsi": "Mata kuliah dasar pemrograman...",
        "dosen_pengampu": {
          "id": "uuid",
          "name": "Dr. John Doe"
        },
        "is_active": true,
        "created_at": "2024-01-15T10:00:00Z"
      }
    ],
    "pagination": { ... }
  }
}
```

---

### 4.2 Get Mata Kuliah by ID

**Endpoint:** `GET /mata-kuliah/:id`

---

### 4.3 Get My Mata Kuliah (Dosen)

**Endpoint:** `GET /mata-kuliah/my`

**Description:** Returns mata kuliah assigned to the current logged-in dosen.

---

### 4.4 Get Mata Kuliah by Semester

**Endpoint:** `GET /mata-kuliah/semester/:semester`

**Example:** `GET /mata-kuliah/semester/3`

---

### 4.5 Get Mata Kuliah by Dosen

**Endpoint:** `GET /mata-kuliah/dosen/:dosen_id`

---

### 4.6 Create Mata Kuliah (Kaprodi Only)

**Endpoint:** `POST /mata-kuliah`

**Request Body:**
```json
{
  "kode": "IF101",
  "nama": "Algoritma dan Pemrograman",
  "sks": 3,
  "semester": 1,
  "deskripsi": "Mata kuliah dasar pemrograman...",
  "dosen_pengampu_id": "dosen-uuid"
}
```

---

### 4.7 Update Mata Kuliah (Kaprodi Only)

**Endpoint:** `PUT /mata-kuliah/:id`

**Request Body:**
```json
{
  "kode": "IF101",
  "nama": "Algoritma dan Pemrograman Dasar",
  "sks": 4,
  "semester": 1,
  "deskripsi": "Updated description...",
  "dosen_pengampu_id": "new-dosen-uuid"
}
```

---

### 4.8 Delete Mata Kuliah (Kaprodi Only)

**Endpoint:** `DELETE /mata-kuliah/:id`

---

### 4.9 Toggle Mata Kuliah Status

**Endpoint:** `PATCH /mata-kuliah/:id/toggle-status`

**Description:** Activates or deactivates a mata kuliah.

---

## 5. CPL Assignment

### 5.1 Get All Assignments

**Endpoint:** `GET /cpl-assignments`

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page | int | Page number |
| limit | int | Items per page |
| cpl_id | string | Filter by CPL |
| dosen_id | string | Filter by Dosen |
| status | string | Filter by status |

**Response:**
```json
{
  "success": true,
  "message": "Assignments retrieved successfully",
  "data": {
    "items": [
      {
        "id": "uuid",
        "cpl": {
          "id": "uuid",
          "kode": "CPL-01",
          "deskripsi": "..."
        },
        "dosen": {
          "id": "uuid",
          "name": "Dr. John Doe"
        },
        "mata_kuliah": {
          "id": "uuid",
          "kode": "IF101",
          "nama": "Algoritma dan Pemrograman"
        },
        "status": "active",
        "assigned_at": "2024-01-15T10:00:00Z"
      }
    ],
    "pagination": { ... }
  }
}
```

---

### 5.2 Get Assignment by ID

**Endpoint:** `GET /cpl-assignments/:id`

---

### 5.3 Get My Assignments (Dosen)

**Endpoint:** `GET /cpl-assignments/my`

**Description:** Returns CPL assignments for the current logged-in dosen.

---

### 5.4 Get Assignments by CPL

**Endpoint:** `GET /cpl-assignments/cpl/:cpl_id`

---

### 5.5 Create Assignment (Kaprodi Only)

**Endpoint:** `POST /cpl-assignments`

**Request Body:**
```json
{
  "cpl_id": "cpl-uuid",
  "dosen_id": "dosen-uuid",
  "mata_kuliah_id": "mk-uuid",
  "catatan": "Notes for this assignment..."
}
```

---

### 5.6 Update Assignment Status

**Endpoint:** `PATCH /cpl-assignments/:id/status`

**Request Body:**
```json
{
  "status": "completed"
}
```

**Valid Status Values:** `active`, `completed`, `cancelled`

---

### 5.7 Delete Assignment (Kaprodi Only)

**Endpoint:** `DELETE /cpl-assignments/:id`

---

## 6. RPS (Rencana Pembelajaran Semester)

### 6.1 Get All RPS

**Endpoint:** `GET /rps`

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page | int | Page number |
| limit | int | Items per page |
| status | string | Filter by status |
| mata_kuliah_id | string | Filter by mata kuliah |
| dosen_id | string | Filter by dosen |
| tahun_ajaran | string | Filter by academic year |

**Response:**
```json
{
  "success": true,
  "message": "RPS retrieved successfully",
  "data": {
    "items": [
      {
        "id": "uuid",
        "mata_kuliah": {
          "id": "uuid",
          "kode": "IF101",
          "nama": "Algoritma dan Pemrograman"
        },
        "dosen": {
          "id": "uuid",
          "name": "Dr. John Doe"
        },
        "tahun_ajaran": "2024/2025",
        "semester": "Ganjil",
        "status": "draft",
        "created_at": "2024-01-15T10:00:00Z"
      }
    ],
    "pagination": { ... }
  }
}
```

---

### 6.2 Get RPS by ID

**Endpoint:** `GET /rps/:id`

**Response (Detail):**
```json
{
  "success": true,
  "message": "RPS retrieved successfully",
  "data": {
    "id": "uuid",
    "mata_kuliah": { ... },
    "dosen": { ... },
    "tahun_ajaran": "2024/2025",
    "semester": "Ganjil",
    "deskripsi_singkat": "Short description...",
    "pustaka_utama": "Main references...",
    "pustaka_pendukung": "Supporting references...",
    "status": "draft",
    "cpmk": [ ... ],
    "rencana_pembelajaran": [ ... ],
    "bahan_bacaan": [ ... ],
    "evaluasi": [ ... ]
  }
}
```

---

### 6.3 Get My RPS (Dosen)

**Endpoint:** `GET /rps/my`

---

### 6.4 Get RPS by Mata Kuliah

**Endpoint:** `GET /rps/mata-kuliah/:mata_kuliah_id`

---

### 6.5 Create RPS

**Endpoint:** `POST /rps`

**Request Body:**
```json
{
  "mata_kuliah_id": "mk-uuid",
  "tahun_ajaran": "2024/2025",
  "semester": "Ganjil",
  "deskripsi_singkat": "Mata kuliah ini membahas...",
  "pustaka_utama": "1. Introduction to Algorithms\n2. Clean Code",
  "pustaka_pendukung": "1. Design Patterns\n2. SICP"
}
```

---

### 6.6 Update RPS

**Endpoint:** `PUT /rps/:id`

**Request Body:**
```json
{
  "tahun_ajaran": "2024/2025",
  "semester": "Ganjil",
  "deskripsi_singkat": "Updated description...",
  "pustaka_utama": "Updated references...",
  "pustaka_pendukung": "Updated supporting references..."
}
```

---

### 6.7 Delete RPS

**Endpoint:** `DELETE /rps/:id`

---

### 6.8 RPS Workflow

#### Submit RPS for Review

**Endpoint:** `PATCH /rps/:id/submit`

**Description:** Changes RPS status from `draft` to `pending`.

---

#### Approve RPS (Kaprodi Only)

**Endpoint:** `PATCH /rps/:id/approve`

**Request Body:**
```json
{
  "catatan": "RPS has been approved. Good work!"
}
```

---

#### Reject RPS (Kaprodi Only)

**Endpoint:** `PATCH /rps/:id/reject`

**Request Body:**
```json
{
  "catatan": "Please revise the evaluation section."
}
```

---

#### Request Revision (Kaprodi Only)

**Endpoint:** `PATCH /rps/:id/request-revision`

**Request Body:**
```json
{
  "catatan": "Please add more details to CPMK section."
}
```

---

### 6.9 CPMK (Capaian Pembelajaran Mata Kuliah)

#### Get CPMK by RPS

**Endpoint:** `GET /rps/:id/cpmk`

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": "uuid",
      "kode": "CPMK-01",
      "deskripsi": "Mahasiswa mampu memahami konsep dasar algoritma",
      "cpl_id": "cpl-uuid",
      "cpl": {
        "kode": "CPL-01"
      },
      "bobot": 25.0
    }
  ]
}
```

---

#### Create CPMK

**Endpoint:** `POST /rps/:id/cpmk`

**Request Body:**
```json
{
  "kode": "CPMK-01",
  "deskripsi": "Mahasiswa mampu memahami konsep dasar algoritma",
  "cpl_id": "cpl-uuid",
  "bobot": 25.0
}
```

---

#### Update CPMK

**Endpoint:** `PUT /rps/:rps_id/cpmk/:cpmk_id`

---

#### Delete CPMK

**Endpoint:** `DELETE /rps/:rps_id/cpmk/:cpmk_id`

---

### 6.10 Rencana Pembelajaran (Weekly Learning Plan)

#### Get Rencana Pembelajaran

**Endpoint:** `GET /rps/:id/rencana-pembelajaran`

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": "uuid",
      "minggu_ke": 1,
      "topik": "Introduction to Algorithms",
      "sub_topik": "What is an algorithm, complexity analysis",
      "metode_pembelajaran": "Lecture, Discussion",
      "pengalaman_belajar": "Students will understand basic concepts...",
      "indikator_penilaian": "Quiz, Assignment",
      "bobot_penilaian": 5.0,
      "waktu_menit": 100
    }
  ]
}
```

---

#### Create Rencana Pembelajaran

**Endpoint:** `POST /rps/:id/rencana-pembelajaran`

**Request Body:**
```json
{
  "minggu_ke": 1,
  "topik": "Introduction to Algorithms",
  "sub_topik": "What is an algorithm, complexity analysis",
  "metode_pembelajaran": "Lecture, Discussion",
  "pengalaman_belajar": "Students will understand basic concepts...",
  "indikator_penilaian": "Quiz, Assignment",
  "bobot_penilaian": 5.0,
  "waktu_menit": 100
}
```

---

#### Update Rencana Pembelajaran

**Endpoint:** `PUT /rps/:rps_id/rencana-pembelajaran/:rp_id`

---

#### Delete Rencana Pembelajaran

**Endpoint:** `DELETE /rps/:rps_id/rencana-pembelajaran/:rp_id`

---

### 6.11 Bahan Bacaan (References)

#### Get Bahan Bacaan

**Endpoint:** `GET /rps/:id/bahan-bacaan`

---

#### Create Bahan Bacaan

**Endpoint:** `POST /rps/:id/bahan-bacaan`

**Request Body:**
```json
{
  "jenis": "buku",
  "judul": "Introduction to Algorithms",
  "penulis": "Thomas H. Cormen",
  "penerbit": "MIT Press",
  "tahun": 2009,
  "url": "https://mitpress.mit.edu/books/introduction-algorithms",
  "is_utama": true
}
```

**Valid Jenis Values:** `buku`, `jurnal`, `artikel`, `website`, `lainnya`

---

#### Update Bahan Bacaan

**Endpoint:** `PUT /rps/:rps_id/bahan-bacaan/:bb_id`

---

#### Delete Bahan Bacaan

**Endpoint:** `DELETE /rps/:rps_id/bahan-bacaan/:bb_id`

---

### 6.12 Evaluasi (Assessments)

#### Get Evaluasi

**Endpoint:** `GET /rps/:id/evaluasi`

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": "uuid",
      "jenis": "UTS",
      "deskripsi": "Mid-term examination covering weeks 1-7",
      "bobot": 30.0,
      "cpmk_terkait": ["CPMK-01", "CPMK-02"],
      "kriteria_penilaian": "Accuracy, Problem solving approach"
    }
  ]
}
```

---

#### Create Evaluasi

**Endpoint:** `POST /rps/:id/evaluasi`

**Request Body:**
```json
{
  "jenis": "UTS",
  "deskripsi": "Mid-term examination covering weeks 1-7",
  "bobot": 30.0,
  "cpmk_terkait": ["CPMK-01", "CPMK-02"],
  "kriteria_penilaian": "Accuracy, Problem solving approach"
}
```

**Valid Jenis Values:** `tugas`, `quiz`, `UTS`, `UAS`, `praktikum`, `proyek`, `lainnya`

---

#### Update Evaluasi

**Endpoint:** `PUT /rps/:rps_id/evaluasi/:eval_id`

---

#### Delete Evaluasi

**Endpoint:** `DELETE /rps/:rps_id/evaluasi/:eval_id`

---

## 7. Notifications

### 7.1 Get My Notifications

**Endpoint:** `GET /notifications`

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page | int | Page number |
| limit | int | Items per page |
| is_read | bool | Filter by read status |
| type | string | Filter by notification type |

**Response:**
```json
{
  "success": true,
  "data": {
    "items": [
      {
        "id": "uuid",
        "title": "RPS Approved",
        "message": "Your RPS for Algoritma dan Pemrograman has been approved.",
        "type": "rps_approved",
        "is_read": false,
        "data": {
          "rps_id": "rps-uuid",
          "mata_kuliah": "IF101"
        },
        "created_at": "2024-01-15T10:00:00Z"
      }
    ],
    "pagination": { ... }
  }
}
```

---

### 7.2 Get Unread Count

**Endpoint:** `GET /notifications/unread-count`

**Response:**
```json
{
  "success": true,
  "data": {
    "count": 5
  }
}
```

---

### 7.3 Get Notification by ID

**Endpoint:** `GET /notifications/:id`

---

### 7.4 Create Notification (Kaprodi Only)

**Endpoint:** `POST /notifications`

**Request Body:**
```json
{
  "user_id": "user-uuid",
  "title": "New Assignment",
  "message": "You have been assigned to teach IF101",
  "type": "assignment"
}
```

---

### 7.5 Mark as Read

**Endpoint:** `PATCH /notifications/:id/read`

---

### 7.6 Mark All as Read

**Endpoint:** `PATCH /notifications/read-all`

---

### 7.7 Delete Notification

**Endpoint:** `DELETE /notifications/:id`

---

## 8. Dashboard

### 8.1 Get Dashboard Statistics

**Endpoint:** `GET /dashboard/stats`

**Description:** Returns role-based dashboard statistics.

**Response (Kaprodi):**
```json
{
  "success": true,
  "data": {
    "total_users": 50,
    "total_dosen": 45,
    "total_cpl": 10,
    "total_mata_kuliah": 30,
    "total_rps": 100,
    "rps_pending": 15,
    "rps_approved": 80,
    "rps_revision": 5
  }
}
```

**Response (Dosen):**
```json
{
  "success": true,
  "data": {
    "my_mata_kuliah": 5,
    "my_rps": 5,
    "rps_draft": 2,
    "rps_pending": 1,
    "rps_approved": 2,
    "my_cpl_assignments": 8
  }
}
```

---

### 8.2 Get Dosen Statistics

**Endpoint:** `GET /dashboard/dosen-stats`

**Description:** Returns statistics for the logged-in dosen.

---

## 9. Documents

### 9.1 Get All Documents

**Endpoint:** `GET /documents`

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| page | int | Page number |
| limit | int | Items per page |
| type | string | Filter by document type |

---

### 9.2 Get Document by ID

**Endpoint:** `GET /documents/:id`

---

### 9.3 Generate Document

**Endpoint:** `POST /documents/generate/:type`

**Request Body:**
```json
{
  "rps_id": "rps-uuid",
  "format": "pdf"
}
```

**Valid Types:** `rps`, `silabus`, `matriks_cpl`

---

### 9.4 Download Document

**Endpoint:** `GET /documents/:id/download`

---

### 9.5 Delete Document

**Endpoint:** `DELETE /documents/:id`

---

## 10. File Upload

### 10.1 Upload File

**Endpoint:** `POST /files/upload`

**Headers:**
```
Authorization: Bearer <access_token>
Content-Type: multipart/form-data
```

**Form Data:**
| Field | Type | Description |
|-------|------|-------------|
| file | file | The file to upload |

**Response:**
```json
{
  "success": true,
  "message": "File uploaded successfully",
  "data": {
    "filename": "generated-uuid-filename.pdf",
    "original_name": "document.pdf",
    "url": "/uploads/generated-uuid-filename.pdf",
    "size": 1024000,
    "mime_type": "application/pdf"
  }
}
```

---

### 10.2 Get File

**Endpoint:** `GET /files/:filename`

**Description:** Serves the uploaded file.

---

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| APP_NAME | Application name | Kurikulum Management System |
| APP_ENV | Environment (development/production) | development |
| APP_PORT | Server port | 8080 |
| APP_DEBUG | Debug mode | true |
| DB_HOST | Database host | localhost |
| DB_PORT | Database port | 3306 |
| DB_NAME | Database name | kurikulum_db |
| DB_USER | Database user | root |
| DB_PASSWORD | Database password | |
| JWT_SECRET | JWT secret key | your-secret-key |
| JWT_EXPIRES_IN | Access token expiry | 24h |
| JWT_REFRESH_EXPIRES_IN | Refresh token expiry | 168h |
| UPLOAD_PATH | Upload directory | ./uploads |
| MAX_FILE_SIZE | Max upload size (bytes) | 10485760 |
| CORS_ALLOWED_ORIGINS | Allowed CORS origins | http://localhost:3000 |

---

## Default Credentials

After running migrations, a default Kaprodi user is created:

| Field | Value |
|-------|-------|
| Email | kaprodi@university.ac.id |
| Password | password123 |
| Role | kaprodi |

---

## Error Codes

| HTTP Code | Description |
|-----------|-------------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request - Invalid input |
| 401 | Unauthorized - Invalid or missing token |
| 403 | Forbidden - Insufficient permissions |
| 404 | Not Found - Resource not found |
| 409 | Conflict - Resource already exists |
| 422 | Unprocessable Entity - Validation error |
| 500 | Internal Server Error |

---

## Rate Limiting

API memiliki rate limiting untuk mencegah abuse:
- **Default**: 100 requests per minute per IP
- **Authentication endpoints**: 10 requests per minute per IP

---

## CORS Configuration

CORS dikonfigurasi untuk mengizinkan request dari frontend:

```go
Allowed Origins: http://localhost:3000
Allowed Methods: GET, POST, PUT, PATCH, DELETE, OPTIONS
Allowed Headers: Origin, Content-Type, Accept, Authorization
```

---

## License

MIT License

---

## Author

Kurikulum Management Team
