package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Helloy igor"))
	});

	port := "6060"
	log.Printf("App run in port %v \n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
