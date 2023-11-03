package main

import (
	"database/sql"
	"github.com/hallgren/eventsourcing"
	eventstore "github.com/hallgren/eventsourcing/eventstore/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nbari/violetear"
	"log"
	"net/http"
)

var repo *eventsourcing.Repository

func main() {
	db, err := sql.Open("sqlite3", "./eventstore.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	es := eventstore.Open(db)
	es.Migrate()

	repo = eventsourcing.NewRepository(es)
	repo.Register(&TestKit{})

	router := violetear.New()
	router.LogRequests = true
	router.RequestID = "Request-ID"

	router.AddRegex(":testID", `DE\d{2}[A-Z0-9]{12}`)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	router.HandleFunc("/testkit", CreateTestsHandler, http.MethodPost)
	router.HandleFunc("/testkit/:testID", GetTestHandler, http.MethodGet)
	router.HandleFunc("/testkit/:testID/sync", TestSyncHandler, http.MethodPut)
	router.HandleFunc("/testkit/:testID/lab", TestArrivedAtLabHandler, http.MethodPut)

	srv := &http.Server{
		Addr:    ":1337",
		Handler: router,
	}
	log.Fatal(srv.ListenAndServe())
}
