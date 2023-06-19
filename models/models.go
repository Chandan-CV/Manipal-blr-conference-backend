package models

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
	Email         string `json:"email" gorm:"not null; unique"`
	Password      string `json:"password" gorm:"not null; unique"`
	EmailVerified bool   `json:"email_verified" gorm:"not null"`
}
