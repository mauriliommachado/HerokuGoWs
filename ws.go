package main

import (
	"github.com/bmizerany/pat"
	"io"
	"log"
	"net/http"
	"os"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	m := pat.New()
	m.Get("/", http.HandlerFunc(HelloWorld))
	m.Get("/hello/:name", http.HandlerFunc(HelloServer))
	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	err := http.ListenAndServe(os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

