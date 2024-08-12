package main

import (
	"handlingcsv/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/report", controller.ReportHandler)
	http.HandleFunc("/report/csv", controller.ReportCSVHandler)
	http.HandleFunc("/report/excel", controller.ReportExcelHandler)

	port := ":8080"
	log.Printf("starting server on port %s", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
