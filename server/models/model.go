package models

import "time"

type Incident struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	AffectedService string    `json:"affected_service"`
	AISeverity      string    `json:"ai_severity"`
	AICategory      string    `json:"ai_category"`
	CreatedAt       time.Time `json:"created_at"`
}

type IncidentReq struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	AffectedService string `json:"affected_service"`
}
