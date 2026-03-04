package auth

type LoginPayload struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}
