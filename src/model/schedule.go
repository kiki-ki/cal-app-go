package model

import "time"

type Schedule struct {
	ID        int
	Title     string     `json:"title"`
	SubTitle  string     `json:"sub-title"`
	Place     string     `json:"place"`
	FromDate  *time.Time `json:"from-date"`
	ToDate    *time.Time `json:"to-date"`
	UpdatedAt *time.Time
	CreatedAt *time.Time
}
