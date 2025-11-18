package entity

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birthDate"`
	Age       int64     `json:"age"`
}
