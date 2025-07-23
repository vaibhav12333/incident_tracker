package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	InDB "kafkaP/server/db"
	"kafkaP/server/models"

	_ "github.com/mattn/go-sqlite3" // Use SQLite for testing
)

// Helper to create a test DB and table
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE incidents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		affected_service TEXT,
		ai_severity TEXT,
		ai_category TEXT,
		created_at DATETIME
	)`)
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestPostHandler_NewIncident(t *testing.T) {
	db := setupTestDB(t)
	handler := http.HandlerFunc(PostHandler(db))

	incident := models.IncidentReq{
		Title:           "Test Incident",
		Description:     "Test Description",
		AffectedService: "Test Service",
	}
	body, _ := json.Marshal(incident)
	req := httptest.NewRequest("POST", "/incidents", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rr.Code)
	}
}

func TestPostHandler_DuplicateIncident(t *testing.T) {
	db := setupTestDB(t)
	incident := models.Incident{
		Title:           "Dup Incident",
		Description:     "Dup Desc",
		AffectedService: "Dup Service",
		AISeverity:      "Low",
		AICategory:      "Software",
	}
	_, _ = InDB.InsertIncidents(db, incident)

	handler := http.HandlerFunc(PostHandler(db))
	incidentReq := models.IncidentReq{
		Title:           "Dup Incident",
		Description:     "Dup Desc",
		AffectedService: "Dup Service",
	}
	body, _ := json.Marshal(incidentReq)
	req := httptest.NewRequest("POST", "/incidents", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusConflict {
		t.Errorf("expected status 409, got %d", rr.Code)
	}
}

func TestGetHandler(t *testing.T) {
	db := setupTestDB(t)
	// Insert a sample incident
	incident := models.Incident{
		Title:           "Get Incident",
		Description:     "Get Desc",
		AffectedService: "Get Service",
		AISeverity:      "Medium",
		AICategory:      "Network",
	}
	_, _ = InDB.InsertIncidents(db, incident)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/incidents", nil)
	GetHandler(db, rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
	var incidents []models.Incident
	if err := json.Unmarshal(rr.Body.Bytes(), &incidents); err != nil {
		t.Errorf("invalid JSON response: %v", err)
	}
	if len(incidents) == 0 {
		t.Errorf("expected at least one incident")
	}
}
