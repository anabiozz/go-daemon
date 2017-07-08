package main

import (
	"fmt"
	"log"
	"net/http"

	postgresql "github.com/anabiozz/go-daemon/storages/postgresql"
)

type data struct {
	id   string `db:"uuid"`
	pid  string `db:"pid"`
	data string `db:"data"`
}

func handler(res http.ResponseWriter, req *http.Request) {
	parametrs := req.URL.Query()

	if len(parametrs) < 1 {
		log.Println("Url param 'key' is missing")
		return
	}

	record := data{}
	record.id = "1"
	record.pid = parametrs["pid"][0]
	record.data = parametrs["data"][0]

	log.Println(record)

	_, err := postgresql.Postgresql_connect().NamedExec(`insert into data values (:id,:pid,:data)`, record)
	if err != nil {
		fmt.Println("postgresql error ", err)
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
