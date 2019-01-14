package entities

import "time"

type User struct {
	Id            uint       `json:"id"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	FirstKanaName string     `json:"firstKanaName"`
	LastKanaName  string     `json:"lastKanaName"`
	Email         string     `json:"email"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     *time.Time `json:"-"`
	DeletedAt     *time.Time `json:"-"`
}
