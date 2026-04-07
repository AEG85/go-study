package models

import "time"

type Book struct {
	ID              int
	Title           string
	Author          string
	Review          string
	PublicationYear int
	IsRead          bool
	DateAdded       time.Time
	DateReaded      *time.Time
}
