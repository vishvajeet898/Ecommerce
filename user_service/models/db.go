package models

type Users struct {
	User_ID  string `json:"user_ID" gorm:"primaryKey,column:users"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	IsAdmin  bool   `json:"isAdmin" gorm:"column:is_admin"`
}
