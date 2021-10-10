package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Uint("p", 9876, "port number")
	flag.Parse()

	var srv *server
	http.HandleFunc("/create", srv.ReceiveData)
	http.HandleFunc("/read", srv.ServeData)
	http.HandleFunc("/update", srv.UpdateData)
	http.HandleFunc("/delete", srv.DeleteData)
	log.Printf("Listening on port %v\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), nil))
}
