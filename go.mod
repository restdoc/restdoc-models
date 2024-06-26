module github.com/restdoc/restdoc-models

go 1.17

replace github.com/restdoc/restdoc-models => ../restdoc-models

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/cockroachdb/cockroach-go/v2 v2.3.7
	github.com/go-sql-driver/mysql v1.6.0
	gorm.io/driver/mysql v1.3.2
	gorm.io/driver/postgres v1.5.7
	gorm.io/gorm v1.25.7-0.20240204074919-46816ad31dde
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
