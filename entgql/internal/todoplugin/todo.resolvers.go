package todoplugin

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"entgo.io/contrib/entgql/internal/todoplugin/ent"
	"entgo.io/contrib/entgql/internal/todoplugin/ent/category"
	"entgo.io/contrib/entgql/internal/todoplugin/ent/todo"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, todo TodoInput) (*ent.Todo, error) {
	client := ent.FromContext(ctx)
	return client.Todo.
		Create().
		SetStatus(todo.Status).
		SetNillablePriority(todo.Priority).
		SetText(todo.Text).
		SetNillableParentID(todo.Parent).
		Save(ctx)
}

func (r *mutationResolver) ClearTodos(ctx context.Context) (int, error) {
	client := ent.FromContext(ctx)
	return client.Todo.
		Delete().
		Exec(ctx)
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

func (r *todoResolver) Category(ctx context.Context, obj *ent.Todo) (*ent.Category, error) {
	e, err := r.client.Category.
		Query().
		Where(category.HasTodosWith(todo.ID(obj.ID))).
		First(ctx)
	return e, ent.MaskNotFound(err)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
