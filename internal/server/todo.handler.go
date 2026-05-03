package server

import (
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"

	"github.com/akagiyuu/todo-api/internal/todo"
)

func (s *Server) CreateTodo(c fuego.ContextWithBody[CreateTodoRequest]) (uuid.UUID, error) {
	accountID := c.Value(AuthKey).(uuid.UUID)

	req, err := c.Body()
	if err != nil {
		return uuid.Nil, err
	}

	id, err := s.TodoService.Create(
		c.Context(),
		accountID,
		req.Title,
		req.Content,
		req.Priority,
	)
	if err != nil {
		return uuid.Nil, fuego.BadRequestError{
			Err:    err,
			Detail: "todo with given title already existed",
		}
	}

	return id, nil
}

func (s *Server) FilterTodo(c fuego.ContextNoBody) ([]todo.Todo, error) {
	accountID := c.Value(AuthKey).(uuid.UUID)

	var filter FilterTodoParams
	err := s.Decoder.Decode(&filter, c.QueryParams())
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid query params",
		}
	}

	todos, err := s.TodoService.Filter(
		c.Context(),
		accountID,
		filter.Query,
		filter.Priority,
		filter.IsDone,
	)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "failed to query todo with given params",
		}
	}

	return todos, nil
}

func (s *Server) GetTodo(c fuego.ContextNoBody) (*todo.Todo, error) {
	accountID := c.Value(AuthKey).(uuid.UUID)

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Required UUID v4",
		}
	}

	todo, err := s.TodoService.Get(c.Context(), accountID, id)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "no todo with given id",
		}
	}

	return todo, nil
}

func (s *Server) UpdateTodo(c fuego.ContextWithBody[UpdateTodoRequest]) (any, error) {
	accountID := c.Value(AuthKey).(uuid.UUID)

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Required UUID v4",
		}
	}

	req, err := c.Body()
	if err != nil {
		return nil, err
	}

	s.TodoService.Update(
		c.Context(),
		accountID,
		id,
		req.Title,
		req.Content,
		req.Priority,
	)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to update todo",
		}
	}

	return nil, nil
}

func (s *Server) DeleteTodo(c fuego.ContextNoBody) (any, error) {
	accountID := c.Value(AuthKey).(uuid.UUID)

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Required UUID v4",
		}
	}

	s.TodoService.Delete(c.Context(), accountID, id)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to update todo",
		}
	}

	return nil, nil
}
