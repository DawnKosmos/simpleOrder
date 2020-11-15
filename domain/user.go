package domain

type User struct {
	Model     `db:",inline"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	IsActive  *bool  `json:"isActive" db:"is_active"`
	IsAdmin   *bool  `json:"isAdmin" db:"is_admin"`
}
