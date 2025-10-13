package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
)

const (
	aeromexicoEnabled = false
	aircanadaEnabled  = false
	deltaEnabled      = true
	unitedEnabled     = false
	virginEnabled     = false
)

var (
	commitHash = os.Getenv("HEROKU_SLUG_COMMIT")
)

func listenAddress() string {
	if port := os.Getenv("PORT"); port != "" {
		return "0.0.0.0:" + port
	}

	return "0.0.0.0:8080"
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/help", HelpHandler).Methods("GET")


	if deltaEnabled {
		r.HandleFunc("/delta", DeltaHomeHandler).Methods("GET")
		r.HandleFunc("/delta", DeltaRetrieveHandler).Methods("POST")
	}


	srv := &http.Server{
		Handler: r,
		Addr:    listenAddress(),

		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	r.Use(secure.New(secure.Options{
		BrowserXssFilter:     true,
		ContentTypeNosniff:   true,
		FrameDeny:            true,
		STSSeconds:           31536000,
		STSIncludeSubdomains: true,
		STSPreload:           true,
	}).Handler)

	log.Println("Visit", "http://"+listenAddress(), "to use the app!")
	log.Fatal(srv.ListenAndServe())
}
