package todo

type GetUserTodosRequest struct {
	UserId string `json:"userId"`
}

type TodoDto struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   int64   `json:"createdAt"`
	IsCompleted bool    `json:"isCompleted"`
}

type GetUserTodosResponse struct {
	Data []TodoDto `json:"data"`
}
