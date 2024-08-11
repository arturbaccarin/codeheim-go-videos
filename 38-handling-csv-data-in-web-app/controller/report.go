package controller

import (
	"bytes"
	"encoding/csv"
	"handlingcsv/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type PageData struct {
	Title     string
	UserStats *[]models.UserStats
}

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := PageData{
			Title:     "Report",
			UserStats: models.GetUserStats(),
		}

		tmpl, err := template.ParseFiles("views/home.html")
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			log.Printf("error loading template: %v", err)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
			log.Printf("error rendering template: %v", err)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ReportCSVHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write the header
	writer.Write([]string{"UserID", "UserName", "LoginCount", "LastLogin", "Active"})

	userStats := models.GetUserStats()
	// Write the data
	for _, user := range *userStats {
		writer.Write([]string{
			strconv.Itoa(user.UserID),
			user.UserName,
			strconv.Itoa(user.LoginCount),
			user.LastLogin.Format(time.RFC3339),
			strconv.FormatBool(user.Active),
		})
	}

	writer.Flush()

	// Set the response headers
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=user_stats.csv")
	w.Write(buf.Bytes())
}

func ReportCSVDownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		log.Printf("Error reading file: %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Unable to parse CSV", http.StatusBadRequest)
		log.Printf("Error parsing CSV: %v", err)
		return
	}

	err = models.AddRows(&rows)
	if err != nil {
		http.Error(w, "Failed to add rows", http.StatusInternalServerError)
		log.Printf("Failed to add rows: %v", err)
		return
	}

	http.Redirect(w, r, "/report", http.StatusSeeOther)
}
