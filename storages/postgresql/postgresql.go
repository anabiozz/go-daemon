package postgresql

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Schema table
var Schema = `
CREATE TABLE IF NOT EXISTS data (
    id  serial primary key not null,
    pid int not null,
    data text
);`

var InsertRequest = `INSERT INTO data (pid, data) VALUES (:pid, :data) RETURNING id`

type namedPreparer interface {
	PrepareNamed(string) (*sqlx.NamedStmt, error)
}

// Connection (connection to postgres)
func Connection() *sqlx.DB {
	postgresUser := os.Getenv("POSTGRES_ENV_POSTGRESS_USER")
	postgresPass := os.Getenv("POSTGRES_ENV_POSTGRESS_PASSWORD")
	postgresAddr := os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
	var postgresqlConnection *sqlx.DB
	if len(postgresAddr) > 0 {
		postgresqlConnection = connect(postgresAddr, postgresPass, postgresUser)
	} else {
		postgresqlConnection = connect("localhost", "Almera103", "almex")
	}
	postgresqlConnection.MustExec(Schema)

	return postgresqlConnection
}

func connect(postgresAddr string, postgresPass string, postgresUser string) *sqlx.DB {
	dsn := fmt.Sprintf("postgres://%v:%v@"+postgresAddr+":5432/postgres?sslmode=disable", postgresUser, postgresPass)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Println("postgresql error: ", err)
	}
	return db
}

// NamedGet (custom function PrepareNamed + returning last id)
func NamedGet(db namedPreparer, dest *int64, query string, arg interface{}) error {
	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	return stmt.Get(dest, arg)
}
