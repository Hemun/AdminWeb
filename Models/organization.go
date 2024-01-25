package Models

import "time"

type Organization struct {
	ID          int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	OrgName     string    `json:"orgName"`
	OrgRegister string    `json:"orgRegister"`
	Phone       int       `json:"phone"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}
