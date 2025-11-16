# Unit Testing Documentation

Folder ini berisi unit test untuk aplikasi Sistem Rekam Medis.

## Struktur

```
tests/
├── <nama_file>_test.go  # Unit test untuk berbagai komponen
└── README.md            # Dokumentasi testing
```

## Menjalankan Test

### Menjalankan Semua Test

```bash
cd backend
go test ./tests/... -v
```

### Menjalankan Test Spesifik

```bash
cd backend
go test ./tests/<nama_file>_test.go -v
```

### Menjalankan Test dengan Coverage

```bash
cd backend
go test ./tests/... -cover
```

### Menjalankan Test dengan Coverage Detail

```bash
cd backend
go test ./tests/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Coverage

Unit test mencakup berbagai skenario validasi dan logika bisnis untuk memastikan aplikasi berjalan sesuai dengan yang diharapkan.

### Contoh Test yang Dicakup

#### Validasi Request

- ✅ Valid request dengan semua field
- ✅ Missing required field
- ✅ Field melebihi panjang maksimal
- ✅ Invalid value untuk field tertentu

#### Logika Bisnis

- ✅ Perhitungan nilai total
- ✅ Penanganan kondisi edge case
- ✅ Validasi hubungan antar entitas

### Total Tests: <jumlah_test> test cases

## Dependencies

- `github.com/stretchr/testify` - Assertion library untuk testing
- `github.com/go-playground/validator/v10` - Validation library

## Menambahkan Test Baru

1. Buat file test baru di folder `tests/` dengan format `*_test.go`
2. Import package yang diperlukan
3. Tulis test function dengan prefix `Test`
4. Gunakan `assert` dari testify untuk assertion
5. Jalankan test dengan `go test`

### Contoh Test Function

```go
func TestYourFunction_Valid(t *testing.T) {
    // Arrange
    req := validators.YourRequest{
        Field: "value",
    }

    // Act
    err := validate.Struct(req)

    // Assert
    assert.NoError(t, err, "Description of expected behavior")
}
```

## Best Practices

1. **Naming Convention**: `Test<FunctionName>_<Scenario>`
2. **Test Independence**: Setiap test harus independen dan tidak bergantung pada test lain
3. **Clear Assertions**: Gunakan message yang jelas pada assertion
4. **Positive & Negative Tests**: Test untuk kasus sukses dan gagal
5. **Edge Cases**: Test untuk boundary conditions dan edge cases
6. **AAA Pattern**: Arrange, Act, Assert

## Menjalankan Server dan Test Secara Terpisah

### Menjalankan Server

```bash
cd backend
go run cmd/server/main.go
```

### Menjalankan Test (di terminal berbeda)

```bash
cd backend
go test ./tests/... -v
```

### Menjalankan Test Tanpa Server

Test unit tidak memerlukan server berjalan karena hanya menguji validasi struct.

## CI/CD Integration

Untuk integrasi dengan CI/CD, tambahkan command berikut di pipeline:

```yaml
# Example for GitHub Actions
- name: Run Tests
  run: |
    cd backend
    go test ./tests/... -v -cover
```

## Troubleshooting

### Test Gagal

1. Pastikan semua dependencies sudah terinstall: `go mod tidy`
2. Periksa import path sudah benar
3. Pastikan validator sudah diinisialisasi dengan benar

### Coverage Rendah

1. Identifikasi area yang belum tercover
2. Tambahkan test untuk edge cases
3. Test untuk error scenarios

## Kontribusi

Saat menambahkan fitur baru:

1. Tambahkan unit test untuk validator baru
2. Pastikan coverage minimal 80%
3. Update dokumentasi ini jika diperlukan
