module auth-service

go 1.23.0

toolchain go1.23.8

require (
	github.com/go-sql-driver/mysql v1.9.2
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.37.0
)

require filippo.io/edwards25519 v1.1.0 // indirect
