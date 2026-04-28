package user

type GetCurrentUserRequest struct{}

type UserDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"createdAt"`
}

type GetCurrentUserResponse struct {
	Data UserDto `json:"data"`
}
