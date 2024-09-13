package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"userName"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
