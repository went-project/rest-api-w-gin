# Went Gin API Template

## Kurulum

```bash
go mod download
```

## Konfigürasyon

Bu proje farklı ortamlar için ayrı `.env` dosyaları kullanır. `APP_ENV` ortam değişkeni ile hangi dosyanın kullanılacağını belirleyebilirsiniz.

### Desteklenen Ortamlar

| Ortam | ENV Dosyası | Komut |
|-------|-------------|-------|
| Local  | `.env.local` | `go run main.go` |
| Development | `.env.development` | `APP_ENV=development go run main.go` |
| Test | `.env.test` | `APP_ENV=test go run main.go` |
| Production (varsayılan) | `.env` | `APP_ENV=production go run main.go` |

### Örnek .env Dosyası

```env
DB_USER=myuser
DB_PASS=mypassword
DB_HOST=localhost
DB_PORT=5432
DB_NAME=mydb
```

### Kullanım Örnekleri

```bash
# Local ortamda çalıştır (.env.local kullanır)
go run main.go

# Development ortamında çalıştır
APP_ENV=development go run main.go

# Test ortamında çalıştır
APP_ENV=test go run main.go

# Production ortamında çalıştır
APP_ENV=production go run main.go
```

### Docker ile Kullanım

```bash
# Development ortamı
docker run -e APP_ENV=development myapp

# Production ortamı
docker run -e APP_ENV=production myapp
```

## Çalıştırma

```bash
go run main.go
```