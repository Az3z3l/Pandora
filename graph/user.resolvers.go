package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Az3z3l/GQL-SERVER/graph/generated"
	"github.com/Az3z3l/GQL-SERVER/graph/model"
)

func (r *userwhenResolver) Challenge(ctx context.Context, obj *model.Userwhen) (*model.Challenge, error) {
	achall, err := achallenge(ctx, obj.Challenge)
	if err != nil {
		return nil, err
	}
	return achall, nil
}

// Userwhen returns generated.UserwhenResolver implementation.
func (r *Resolver) Userwhen() generated.UserwhenResolver { return &userwhenResolver{r} }

type userwhenResolver struct{ *Resolver }
