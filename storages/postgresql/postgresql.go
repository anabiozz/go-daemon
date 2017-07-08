package postgresql

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Postgresql_connect() *sqlx.DB {
	POSTGRESS_USER := os.Getenv("POSTGRESS_USER")
	POSTGRESS_PASSWORD := "Almera103" //os.Getenv("POSTGRESS_PASSWORD")
	dsn := fmt.Sprintf("postgres://%v:%v@localhost:5432/go-daemon?sslmode=disable", POSTGRESS_USER, POSTGRESS_PASSWORD)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Println("postgresql error: ", err)
	}

	return db
}
