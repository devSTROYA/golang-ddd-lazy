package todo

type AddTodoRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
}
