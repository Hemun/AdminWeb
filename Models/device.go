package Models

import "time"

type Device struct {
	ID          int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	DeviceName  string    `json:"deviceName"`
	DeviceNo    string    `json:"deviceNo"`
	DeviceTspe  string    `json:"deviceTspe"`
	DevicePrice int       `json:"devicePrice"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}
