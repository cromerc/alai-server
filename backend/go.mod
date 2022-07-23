module backend

go 1.18

require github.com/julienschmidt/httprouter v1.3.0

require golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa

require github.com/joho/godotenv v1.4.0

require github.com/golang-jwt/jwt/v4 v4.4.2

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/handlers v1.5.1
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)
