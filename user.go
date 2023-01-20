package shop

type User struct {
	Id       int    `json:"-" db:"id"`
	Login    string `json:"login" binding:"required" db:"login"`
	Name     string `json:"name" binding:"required" db:"first_name"`
	Surname  string `json:"surname" binding:"required" db:"surname"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}
