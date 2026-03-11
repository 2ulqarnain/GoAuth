package auth

type loginPayload struct {
	Email    string `json:"email,required"`
	Password string `json:"password,required"`
}

type signupPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type signupResponseData struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"required,email"`
}

type signupResponse struct {
	Ok      bool                `json:"ok"`
	Message string              `json:"message"`
	Data    *signupResponseData `json:"data"`
}
