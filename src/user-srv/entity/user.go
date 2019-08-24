package entity

type User struct {
	UserId	int64 `json:"user_id" db:"user_id"`
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
	Email string `json:"email" db:"email"`
	Phone string `json:"phone" db:"phone"`
}