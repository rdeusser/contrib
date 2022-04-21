package todoplugin

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"entgo.io/contrib/entgql/internal/todoplugin/ent"
)

func (r *masterUserResolver) Age(ctx context.Context, obj *ent.User) (float64, error) {
	return float64(obj.Age), nil
}

func (r *masterUserResolver) Amount(ctx context.Context, obj *ent.User) (float64, error) {
	return float64(obj.Amount), nil
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
		)
}

// MasterUser returns MasterUserResolver implementation.
func (r *Resolver) MasterUser() MasterUserResolver { return &masterUserResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type masterUserResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
