package todo

import (
	"time"

	"github.com/akagiyuu/todo-api/internal/database"
	"github.com/google/uuid"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type Todo struct {
	ID        uuid.UUID `json:"id"`
	AccountID uuid.UUID `json:"account_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Priority  Priority  `json:"priority"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}

func ParseTodo(raw database.Todo) Todo {
	return Todo{
		ID:        raw.ID,
		AccountID: raw.AccountID,
		Title:     raw.Title,
		Content:   raw.Content,
		Priority:  Priority(raw.Priority),
		IsDone:    raw.IsDone,
		CreatedAt: raw.CreatedAt.Time,
	}
}
