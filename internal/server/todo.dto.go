package server

import "github.com/akagiyuu/todo-api/internal/todo"

type CreateTodoRequest struct {
	Title    string        `json:"title"`
	Content  string        `json:"content"`
	Priority todo.Priority `json:"priority"`
}

type FilterTodoParams struct {
	Query    *string        `form:"query"`
	Priority *todo.Priority `form:"priority"`
	IsDone   *bool          `form:"is_done"`
}

type UpdateTodoRequest struct {
	Title    *string        `json:"title,omitempty"`
	Content  *string        `json:"content,omitempty"`
	Priority *todo.Priority `json:"priority,omitempty"`
}
