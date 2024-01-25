package Models

import "time"

type Device_Log struct {
	ID          int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	DeviceID    int       `json:"deviceId"`
	Value       int       `json:"value"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}
