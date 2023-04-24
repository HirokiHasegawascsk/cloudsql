package main

import (
	"database/sql"
	"fmt"
	"github.com/goark/gocli/config"
	"github.com/goark/gocli/exitcode"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("defaulting to port %s", port)
	}
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		db  *sql.DB
		err error
	)

	if os.Getenv("INSTANCE_HOST") != "" {
		db, err = connectTCPSocket()
		fmt.Fprint(w, db)
		fmt.Fprint(w, err)
		if err != nil {
			fmt.Fprint(w, "connectTCPSocket: unable to connect")
		}
		fmt.Fprint(w, "Connected to Cloud SQL successfully!")
	}
}

func connectTCPSocket() (*sql.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_tcp.go: %s environment variable not set.", k)
		}
		return v
	}
	var (
		dbUser    = mustGetenv("DB_USER")
		dbPwd     = mustGetenv("DB_PASS")
		dbTCPHost = mustGetenv("INSTANCE_HOST")
		dbPort    = mustGetenv("DB_PORT")
		dbName    = mustGetenv("DB_NAME")
	)
	log.Printf("listening on port %s", dbUser)
	log.Printf("listening on port %s", dbPwd)
	log.Printf("listening on port %s", dbTCPHost)
	log.Printf("listening on port %s", dbPort)
	log.Printf("listening on port %s", dbName)

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s",
		dbTCPHost, dbUser, dbPwd, dbPort, dbName)

	dbPool, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	return dbPool, err
}
