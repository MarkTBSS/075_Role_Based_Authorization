```
go get github.com/gofiber/fiber/v2
go get github.com/joho/godotenv
go get github.com/jmoiron/sqlx
go get github.com/jackc/pgx/v5/stdlib
go get github.com/golang-jwt/jwt/v5

console
25/05/2024 [127.0.0.1] 200 - GET /v1
25/05/2024 [127.0.0.1] 404 - GET /v1/ee
201 - POST /v1/users/signup
200 - POST /v1/users/signin

postman
{
    "user": {
        "id": "U000003",
        "email": "test@test.com",
        "username": "test",
        "role_id": 1
    },
    "token": {
        "id": "serial",
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.XXX.XXX",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.XXX.XXX"
    }
}
```