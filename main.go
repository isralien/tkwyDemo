package main

import (
	"log"
	"net/http"
	"tkwyDemo/healthCheck"
	"tkwyDemo/typiCode"
)


func main()  {
	http.HandleFunc("/health", healthCheck.HealthCheck)
	http.HandleFunc("/main", typiCode.HandleMain)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

