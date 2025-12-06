# API Dokumentasi: Assign/Unassign Dosen ke Mata Kuliah

## Overview

Endpoint untuk mengelola penugasan dosen pengampu dan koordinator pada mata kuliah. Hanya dapat diakses oleh user dengan role **kaprodi**.

---

## 1. Assign Dosen ke Mata Kuliah

Menugaskan dosen pengampu dan/atau koordinator ke mata kuliah tertentu.

### Endpoint

```
PATCH /api/v1/mata-kuliah/:id/assign-dosen
```

### Headers

| Header        | Value              | Required |
|---------------|-------------------|----------|
| Authorization | Bearer {token}    | Yes      |
| Content-Type  | application/json  | Yes      |

### Path Parameters

| Parameter | Type   | Description                    |
|-----------|--------|--------------------------------|
| id        | string | UUID dari mata kuliah          |

### Request Body

```json
{
  "dosen_pengampu_id": "uuid-dosen-pengampu",
  "koordinator_id": "uuid-koordinator"
}
```

| Field             | Type   | Required | Description                              |
|-------------------|--------|----------|------------------------------------------|
| dosen_pengampu_id | string | No       | UUID dosen yang akan ditugaskan sebagai pengampu |
| koordinator_id    | string | No       | UUID dosen yang akan ditugaskan sebagai koordinator |

> **Note:** Minimal salah satu field harus diisi. Jika keduanya kosong, akan mengembalikan error.

### Response

#### Success (200 OK)

```json
{
  "status": "success",
  "message": "Dosen berhasil ditugaskan ke mata kuliah",
  "data": {
    "id": "mk-uuid",
    "kode": "IF101",
    "nama": "Algoritma dan Pemrograman",
    "sks": 3,
    "semester": 1,
    "jenis": "wajib",
    "status": "aktif",
    "deskripsi": "Mata kuliah dasar pemrograman",
    "dosen_pengampu": {
      "id": "dosen-uuid",
      "nama": "Dr. John Doe",
      "email": "john.doe@unismuh.ac.id",
      "nidn": "1234567890"
    },
    "koordinator": {
      "id": "koordinator-uuid",
      "nama": "Prof. Jane Smith",
      "email": "jane.smith@unismuh.ac.id",
      "nidn": "0987654321"
    },
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-12-06T08:00:00Z"
  }
}
```

#### Error Responses

**400 Bad Request** - Tidak ada dosen yang ditugaskan
```json
{
  "status": "error",
  "message": "Minimal satu dosen harus ditugaskan (pengampu atau koordinator)"
}
```

**404 Not Found** - Mata kuliah tidak ditemukan
```json
{
  "status": "error",
  "message": "Mata kuliah tidak ditemukan"
}
```

**404 Not Found** - Dosen tidak ditemukan
```json
{
  "status": "error",
  "message": "Dosen pengampu tidak ditemukan"
}
```

**401 Unauthorized** - Token tidak valid
```json
{
  "status": "error",
  "message": "Unauthorized"
}
```

**403 Forbidden** - Bukan kaprodi
```json
{
  "status": "error",
  "message": "Forbidden: only kaprodi can access this resource"
}
```

### Example Request

```bash
curl -X PATCH "http://localhost:8080/api/v1/mata-kuliah/mk-uuid-here/assign-dosen" \
  -H "Authorization: Bearer your-jwt-token" \
  -H "Content-Type: application/json" \
  -d '{
    "dosen_pengampu_id": "dosen-uuid-here",
    "koordinator_id": "koordinator-uuid-here"
  }'
```

---

## 2. Unassign Dosen dari Mata Kuliah

Membatalkan penugasan dosen pengampu dan/atau koordinator dari mata kuliah.

### Endpoint

```
PATCH /api/v1/mata-kuliah/:id/unassign-dosen
```

### Headers

| Header        | Value              | Required |
|---------------|-------------------|----------|
| Authorization | Bearer {token}    | Yes      |
| Content-Type  | application/json  | Yes      |

### Path Parameters

| Parameter | Type   | Description                    |
|-----------|--------|--------------------------------|
| id        | string | UUID dari mata kuliah          |

### Request Body

```json
{
  "type": "pengampu"
}
```

| Field | Type   | Required | Description                                           |
|-------|--------|----------|-------------------------------------------------------|
| type  | string | Yes      | Jenis dosen yang akan di-unassign. Nilai yang valid: `pengampu`, `koordinator`, `both` |

### Type Values

| Value       | Description                                    |
|-------------|------------------------------------------------|
| pengampu    | Hanya unassign dosen pengampu                  |
| koordinator | Hanya unassign koordinator                     |
| both        | Unassign kedua dosen (pengampu dan koordinator)|

### Response

#### Success (200 OK)

```json
{
  "status": "success",
  "message": "Dosen pengampu berhasil di-unassign dari mata kuliah",
  "data": {
    "id": "mk-uuid",
    "kode": "IF101",
    "nama": "Algoritma dan Pemrograman",
    "sks": 3,
    "semester": 1,
    "jenis": "wajib",
    "status": "aktif",
    "deskripsi": "Mata kuliah dasar pemrograman",
    "dosen_pengampu": null,
    "koordinator": {
      "id": "koordinator-uuid",
      "nama": "Prof. Jane Smith",
      "email": "jane.smith@unismuh.ac.id",
      "nidn": "0987654321"
    },
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-12-06T08:30:00Z"
  }
}
```

#### Success Messages by Type

| Type        | Message                                               |
|-------------|-------------------------------------------------------|
| pengampu    | "Dosen pengampu berhasil di-unassign dari mata kuliah"|
| koordinator | "Koordinator berhasil di-unassign dari mata kuliah"   |
| both        | "Semua dosen berhasil di-unassign dari mata kuliah"   |

#### Error Responses

**400 Bad Request** - Type tidak valid
```json
{
  "status": "error",
  "message": "Type harus salah satu dari: pengampu, koordinator, both"
}
```

**404 Not Found** - Mata kuliah tidak ditemukan
```json
{
  "status": "error",
  "message": "Mata kuliah tidak ditemukan"
}
```

**401 Unauthorized** - Token tidak valid
```json
{
  "status": "error",
  "message": "Unauthorized"
}
```

**403 Forbidden** - Bukan kaprodi
```json
{
  "status": "error",
  "message": "Forbidden: only kaprodi can access this resource"
}
```

### Example Requests

**Unassign Dosen Pengampu:**
```bash
curl -X PATCH "http://localhost:8080/api/v1/mata-kuliah/mk-uuid-here/unassign-dosen" \
  -H "Authorization: Bearer your-jwt-token" \
  -H "Content-Type: application/json" \
  -d '{"type": "pengampu"}'
```

**Unassign Koordinator:**
```bash
curl -X PATCH "http://localhost:8080/api/v1/mata-kuliah/mk-uuid-here/unassign-dosen" \
  -H "Authorization: Bearer your-jwt-token" \
  -H "Content-Type: application/json" \
  -d '{"type": "koordinator"}'
```

**Unassign Keduanya:**
```bash
curl -X PATCH "http://localhost:8080/api/v1/mata-kuliah/mk-uuid-here/unassign-dosen" \
  -H "Authorization: Bearer your-jwt-token" \
  -H "Content-Type: application/json" \
  -d '{"type": "both"}'
```

---

## Flow Penggunaan

### Skenario 1: Assign Dosen Baru ke Mata Kuliah

1. Kaprodi login dan mendapat token
2. Kaprodi melihat daftar dosen via `GET /api/v1/users/dosen`
3. Kaprodi melihat daftar mata kuliah via `GET /api/v1/mata-kuliah`
4. Kaprodi assign dosen ke mata kuliah via `PATCH /api/v1/mata-kuliah/:id/assign-dosen`

### Skenario 2: Ganti Dosen Pengampu

1. Kaprodi assign dosen baru dengan `PATCH /api/v1/mata-kuliah/:id/assign-dosen`
   - Cukup kirim `dosen_pengampu_id` baru, akan otomatis mengganti yang lama

### Skenario 3: Hapus Penugasan Dosen

1. Kaprodi unassign dosen via `PATCH /api/v1/mata-kuliah/:id/unassign-dosen`
   - Gunakan `type: "pengampu"` untuk hapus pengampu saja
   - Gunakan `type: "koordinator"` untuk hapus koordinator saja
   - Gunakan `type: "both"` untuk hapus keduanya

---

## Authorization

| Role    | Assign Dosen | Unassign Dosen |
|---------|--------------|----------------|
| kaprodi | ✅ Yes       | ✅ Yes         |
| dosen   | ❌ No        | ❌ No          |
| admin   | ❌ No        | ❌ No          |

---

## Related Endpoints

| Endpoint                           | Method | Description                    |
|------------------------------------|--------|--------------------------------|
| `/api/v1/mata-kuliah`              | GET    | Daftar semua mata kuliah       |
| `/api/v1/mata-kuliah/:id`          | GET    | Detail mata kuliah             |
| `/api/v1/mata-kuliah/dosen/:id`    | GET    | Mata kuliah by dosen           |
| `/api/v1/users/dosen`              | GET    | Daftar semua dosen             |
