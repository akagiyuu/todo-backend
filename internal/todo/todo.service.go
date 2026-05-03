package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/akagiyuu/todo-api/internal/database"
)

type TodoService struct {
	Pool *pgxpool.Pool
}

func (t *TodoService) Create(
	ctx context.Context,
	accountID uuid.UUID,
	title string,
	content string,
	priority Priority,
) (uuid.UUID, error) {
	queries := database.New(t.Pool)
	return queries.CreateTodo(ctx, database.CreateTodoParams{
		AccountID: accountID,
		Title:     title,
		Content:   content,
		Priority:  (database.Priority)(priority),
	})
}

func (t *TodoService) Filter(
	ctx context.Context,
	accountID uuid.UUID,
	query *string,
	priority *Priority,
	isDone *bool,
) ([]Todo, error) {
	queries := database.New(t.Pool)
	raws, err := queries.FilterTodo(ctx, database.FilterTodoParams{
		AccountID: accountID,
		Query:     query,
		IsDone:    isDone,
		Priority:  (*database.Priority)(priority),
	})
	if err != nil {
		return nil, err
	}

	todos := make([]Todo, len(raws))
	for i, raw := range raws {
		todos[i] = ParseTodo(raw)
	}
	return todos, nil
}

func (t *TodoService) Get(
	ctx context.Context,
	accountID uuid.UUID,
	id uuid.UUID,
) (*Todo, error) {
	queries := database.New(t.Pool)
	raw, err := queries.GetTodo(ctx, database.GetTodoParams{
		AccountID: accountID,
		ID:        id,
	})
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	todo := ParseTodo(raw)

	return &todo, nil
}

func (t *TodoService) Update(
	ctx context.Context,
	accountID uuid.UUID,
	id uuid.UUID,
	title *string,
	content *string,
	priority *Priority,
) error {
	queries := database.New(t.Pool)
	return queries.UpdateTodo(ctx, database.UpdateTodoParams{
		AccountID: accountID,
		ID:        id,
		Title:     title,
		Content:   content,
		Priority:  (*database.Priority)(priority),
	})
}

func (t *TodoService) Delete(
	ctx context.Context,
	accountID uuid.UUID,
	id uuid.UUID,
) error {
	queries := database.New(t.Pool)
	return queries.DeleteTodo(ctx, database.DeleteTodoParams{
		AccountID: accountID,
		ID:        id,
	})
}
