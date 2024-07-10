package dto

type UserRequest struct {
	Name    string `json:"name" validate:"required,min=4"`
	Email   string `json:"email" validate:"required,email"`
	Age     uint64 `json:"age" validate:"gte=18,lte=120"`
	Address string `json:"address" validate:"required,min=10"`
	Phone   string `json:"phone" validate:"required,e164"` // https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-E_164_Phone_Number_String
}
