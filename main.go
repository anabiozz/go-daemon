package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello from Go")
}

var pc [256]byte

func main() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Println(byte(i & 1))
	}
	// port := os.Getenv("PORT")
	// if len(port) == 0 {
	// 	port = "8080"
	// }

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", hello)

	// n := negroni.Classic()
	// n.UseHandler(mux)
	// hostString := fmt.Sprintf(":%s", port)
	// n.Run(hostString)

}
