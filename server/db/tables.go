package InDB

import (
	"database/sql"
	"kafkaP/server/models"
	"time"
)

func InsertIncidents(db *sql.DB, inci models.Incident) (bool, error) {
	query := `INSERT INTO incidents( title, description, affected_service, ai_severity, ai_category, created_at) values ($1,$2,$3,$4,$5,$6)`
	_, err := db.Exec(query, inci.Title, inci.Description, inci.AffectedService, inci.AISeverity, inci.AICategory, time.Now())
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetIncidents(db *sql.DB) ([]*models.Incident, error) {
	query := `SELECT id,title, description, affected_service, ai_severity, ai_category, created_at FROM incidents`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incidents []*models.Incident
	for rows.Next() {
		incident := new(models.Incident)
		err := rows.Scan(&incident.ID, &incident.Title, &incident.Description, &incident.AffectedService, &incident.AISeverity, &incident.AICategory, &incident.CreatedAt)
		if err != nil {
			return nil, err
		}
		incidents = append(incidents, incident)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return incidents, nil
}

func IncidentExists(db *sql.DB, inci models.Incident) (bool, error) {
	query := `SELECT COUNT(1) FROM incidents WHERE title=$1 AND description=$2 AND affected_service=$3`
	var count int
	err := db.QueryRow(query, inci.Title, inci.Description, inci.AffectedService).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
