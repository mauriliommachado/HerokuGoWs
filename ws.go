package main

import (
	"github.com/bmizerany/pat"
	"io"
	"log"
	"net/http"
	"os"
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}

func Poti(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>TE AMO MEU AMOR!!<h1>")
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	m := pat.New()
	m.Get("/", http.HandlerFunc(HelloWorld))
	m.Get("/hello/:name", http.HandlerFunc(HelloServer))
	m.Get("/poti", http.HandlerFunc(Poti))
	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	fmt.Println("listening..."+GetPort())
	err = http.ListenAndServe(GetPort(), nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
