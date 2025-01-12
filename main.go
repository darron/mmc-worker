package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/syumai/workers"

	_ "github.com/syumai/workers/cloudflare/d1" // register driver
)

func main() {
	http.HandleFunc("/sqltest", func(w http.ResponseWriter, req *http.Request) {
		id := "9b68b10e-94b3-46c6-8a90-18fbef43c9e6"
		db, err := sql.Open("d1", "DB")
		if err != nil {
			log.Fatalf("error opening DB: %s", err.Error())
		}
		rows, err := db.Query("SELECT * from news_stories WHERE ID = ?", id)
		if err != nil {
			log.Fatalf("error querying DB: %s", err.Error())
		}
		rowsText := fmt.Sprintf("%#v\n", rows)
		msg := "Hello!" + rowsText
		w.Write([]byte(msg))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello!"
		w.Write([]byte(msg))
	})
	http.HandleFunc("/echo", func(w http.ResponseWriter, req *http.Request) {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		io.Copy(w, bytes.NewReader(b))
	})
	workers.Serve(nil) // use http.DefaultServeMux
}
