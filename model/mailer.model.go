package model

type SendOtp struct {
	Email  string `json:"email"`
	UserId int    `json:"user_id"`
}

type VerifyEmail struct {
	Email string `json:"email" validate:"required"`
	Pin   string `json:"pin"  validate:"required"`
}

type MailerRelationResponse struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Email  string `json:"email" validate:"required"`
	Status string `json:"status"`
}
