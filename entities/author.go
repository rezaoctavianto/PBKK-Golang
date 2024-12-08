package entities

import "time"

type Author struct {
	Id   int
	Name string
	DoB  time.Time
	Updated_At time.Time
}