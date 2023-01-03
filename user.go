package shop

type User struct {
	Id       int    `json:"-"`
	Login    string `json:"login" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Password string `json:"password" binding:"required"`
}
