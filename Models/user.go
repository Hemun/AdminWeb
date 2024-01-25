package Models

import "time"

type User struct {
	ID          int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Age         int       `json:"age"`
	Phone       int       `json:"phone"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Gender      string    `json:"gender"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}
