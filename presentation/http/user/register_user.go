package user

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	AccessToken string `json:"accessToken"`
}
