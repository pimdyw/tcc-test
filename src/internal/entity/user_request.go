package entity

type UserRequest struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Address   string `json:"address"`
	BirthDate string `json:"birthDate"`
	Age       int64  `json:"age"`
}
