package dto

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     uint64 `json:"age"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
