package schemas


type LoginSchema struct {
	Email string `json:"email" validate:"required,email,min=12" `
}

type OTPSchema struct {
	OTP string `json:"token" validate:"required"`
	Email string `json:"email" validate:"required,email,min=12" `
}