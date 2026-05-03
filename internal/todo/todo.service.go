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
	priority database.Priority,
) (uuid.UUID, error) {
	queries := database.New(t.Pool)
	return queries.CreateTodo(ctx, database.CreateTodoParams{
		AccountID: accountID,
		Title:     title,
		Content:   content,
		Priority:  priority,
	})
}

func (t *TodoService) Filter(
	ctx context.Context,
	accountID uuid.UUID,
	query *string,
	priority *database.Priority,
	isDone *bool,
) ([]database.Todo, error) {
	queries := database.New(t.Pool)
	return queries.FilterTodo(ctx, database.FilterTodoParams{
		AccountID: accountID,
		Query:     query,
		IsDone:    isDone,
		Priority:  priority,
	})
}

func (t *TodoService) Get(
	ctx context.Context,
	accountID uuid.UUID,
	id uuid.UUID,
) (*database.Todo, error) {
	queries := database.New(t.Pool)
	todo, err := queries.GetTodo(ctx, database.GetTodoParams{
		AccountID: accountID,
		ID:        id,
	})
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *TodoService) Update(
	ctx context.Context,
	accountID uuid.UUID,
	id uuid.UUID,
	title *string,
	content *string,
	priority *database.Priority,
) error {
	queries := database.New(t.Pool)
	return queries.UpdateTodo(ctx, database.UpdateTodoParams{
		AccountID: accountID,
		ID:        id,
		Title:     title,
		Content:   content,
		Priority:  priority,
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
