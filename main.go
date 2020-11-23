package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Started app")
	actor := getFromEnv()
	log.Printf("Using env value: %s\n",actor)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(actor))
	})
	log.Fatal(http.ListenAndServe(":80",nil))
}

func getFromEnv() string {
	actor := "Steven Seagal"
	if a := os.Getenv("ACTOR"); a != "" {
		actor = a
	}
	return actor
}
