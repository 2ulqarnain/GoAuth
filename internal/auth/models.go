package auth

type loginPayload struct {
	Email    string `json:"email,required"`
	Password string `json:"password,required"`
}

type registerUserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
