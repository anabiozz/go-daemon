package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anabiozz/go-daemon/storages/postgresql"
	"github.com/anabiozz/go-daemon/storages/redis"
	"github.com/codegangsta/negroni"
)

// ReturnValueStruct (last saved id)
type ReturnValueStruct struct {
	ID  int64 `json:"id"`
	PID int64 `json:"pid"`
}

func handler(res http.ResponseWriter, req *http.Request) {

	postgresqlConnection := postgresql.Connection()
	redisConnection := redis.Connection()

	defer postgresqlConnection.Close()
	defer redisConnection.Close()

	parameters := req.URL.Query()
	if len(parameters) < 2 {
		log.Println("Must be two url parameters (pid, data)")
		return
	}

	args := make(map[string]interface{})
	for key := range parameters {
		args[key] = parameters[key][0]
	}

	returnValue := ReturnValueStruct{}

	err := postgresql.NamedGet(postgresqlConnection, &returnValue.ID, postgresql.InsertRequest, args)
	if err != nil {
		log.Println("postgresql error ", err)
	}

	js, err := json.Marshal(returnValue)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(js)

	redis.Save(string(returnValue.PID), js, redisConnection)
}

func main() {
	mux := http.NewServeMux()
	go mux.HandleFunc("/", handler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":8080", n)

}
