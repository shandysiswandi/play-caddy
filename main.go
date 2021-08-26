package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("server running on port:", 8080)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("call '%s' from %s", r.URL.String(), r.RemoteAddr)
		log.Println(str)

		w.Header().Add("X-Api-Key", "lorem_ipsum_dolor_amet")
		w.WriteHeader(200)
		w.Write([]byte(""))
	})
	http.ListenAndServe(":8080", nil)
}
