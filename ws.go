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
	"time"
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

func GetAlunos(w http.ResponseWriter, e *http.Request) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM escola.alunos")
	if err != nil {
		log.Fatal(err)
	}
	var respBody string
	for rows.Next() {
		var id int
		var cpf string
		var nome string
		var email string
		var fone string
		var dataNascimento time.Time
		err = rows.Scan(&id, &cpf, &nome, &email, &fone, &dataNascimento)
		respBody +=string(id)+cpf+nome+email+fone+dataNascimento.String()+"\n"
	}
	io.WriteString(w,respBody)
}

func main() {
	m := pat.New()
	m.Get("/aluno", http.HandlerFunc(GetAlunos))
	m.Get("/", http.HandlerFunc(HelloWorld))
	m.Get("/hello/:name", http.HandlerFunc(HelloServer))
	m.Get("/poti", http.HandlerFunc(Poti))
	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	fmt.Println("listening..." + GetPort())
	err := http.ListenAndServe(GetPort(), nil)

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
