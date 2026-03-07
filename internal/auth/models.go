package auth

type LoginPayload struct {
	Email    string `json:"email,required"`
	Password string `json:"password,required"`
}

type RegisterPayload struct {
	Name     string `json:"name,required"`
	Email    string `json:"email,required"`
	Password string `json:"password,required"`
}
