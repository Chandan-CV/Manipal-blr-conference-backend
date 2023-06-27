package models

import "gorm.io/gorm"

type User struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Name          string `json:"name" gorm:"not null"`
	Email         string `json:"email" gorm:"not null; unique"`
	Institution   string `json:"institution"`
	Address       string `json:"address"`
	Phone         string `json:"phone" gorm:"unique"`
	Date          string `json:"date"`
	Accomodation  bool   `json:"accomodation"`
	MaheStaff     bool   `json:"mahe_staff"`
	MaheRollNo    string `json:"mahe_roll_no"`
	Category      string `json:"category"`
	AmountPaid    int    `json:"amount_paid"`
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
	PaymentStatus string `json:"payment_status"`
}

type Auth struct {
	ID            uint   `json:"id" gorm:"primary_key; not null; unique"`
	Email         string `json:"email" gorm:"not null; unique"`
	Password      string `json:"password" gorm:"not null; unique"`
	EmailVerified bool   `json:"email_verified" gorm:"not null"`
}

type Event struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	IsFood      bool   `json:"is_food" gorm:"not null"`
}

type Attendance struct {
	gorm.Model
	ID      uint `json:"id" gorm:"primary_key"`
	EventID uint `json:"event_id" gorm:"not null"`
	UserID  uint `json:"user_id" gorm:"not null"`
	IsFood  bool `json:"is_food"`
}

type EventJWT struct {
	ID     uint   `json:"id"`
	IsFood bool   `json:"is_food"`
	Name   string `json:"name"`
	JWT    string `json:"jwt"`
}

type AttendanceBody struct {
	JWT string `json:"jwt"`
}

type GetAttendancesResponse struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	IsFood    bool   `json:"is_food"`
	Name      string `json:"name"`
}
