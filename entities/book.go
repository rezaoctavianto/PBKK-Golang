package entities

import "time"

type Book struct {
	Id       uint
	Title    string
	Author   Author
	Genre    string
	Added_At time.Time
}
