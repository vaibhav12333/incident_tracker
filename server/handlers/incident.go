package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	InDB "kafkaP/server/db"
	"kafkaP/server/models"
	"kafkaP/server/services"
	"net/http"
)

func IncidentHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method == http.MethodGet {
			GetHandler(db, w, r)
		} else if r.Method == http.MethodPost {
			PostHandler(db)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	}
}

func GetHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	incidents, err := InDB.GetIncidents(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to fetch incidents: " + err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incidents)
}

func PostHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var incidentReq models.IncidentReq
		err := json.NewDecoder(r.Body).Decode(&incidentReq)
		if err != nil {
			fmt.Println("Error decoding incident request:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		severity, category, err := services.GetAIInsights(incidentReq.Title, incidentReq.Description, incidentReq.AffectedService)
		if err != nil {
			fmt.Println("Error getting AI insights:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		incident := models.Incident{
			Title:           incidentReq.Title,
			Description:     incidentReq.Description,
			AffectedService: incidentReq.AffectedService,
			AISeverity:      severity,
			AICategory:      category,
		}

		exists, err := InDB.IncidentExists(db, incident)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error checking for existing incident: " + err.Error()))
			return
		}
		if exists {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("Incident already exists"))
			return
		}

		success, err := InDB.InsertIncidents(db, incident)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("failed to save incidents in db err: ", err)
		}
		if success {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(severity + " " + category))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
