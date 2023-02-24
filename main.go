package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	handler := func(w http.ResponseWriter, req *http.Request) {

		//  do something
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(300)
		time.Sleep(time.Duration(n) * time.Millisecond) // 0〜300 ミリ秒待つ

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
