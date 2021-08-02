package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Az3z3l/GQL-SERVER/graph/generated"
	"github.com/Az3z3l/GQL-SERVER/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, input *model.RegisterInput) (string, error) {
	// return register(ctx, input)
	return "Who are you and why are you so wise in the ways of science", nil
}

func (r *mutationResolver) AddAdmin(ctx context.Context, input *model.RegisterAdminData) (string, error) {
	// return create_admin(ctx, input)
	return "Who are you and why are you so wise in the ways of science", nil
}

func (r *mutationResolver) AddNotifications(ctx context.Context, input model.Notificationinp) (string, error) {
	return add_notification(ctx, input)
}

func (r *mutationResolver) EditNotification(ctx context.Context, input model.Notifiedit) (string, error) {
	return editNoti(ctx, input)
}

func (r *mutationResolver) DeleteNotification(ctx context.Context, id string) (string, error) {
	return delNoti(ctx, id)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.Useredit) (string, error) {
	return useredit(ctx, input)
}

func (r *mutationResolver) ResetPwd(ctx context.Context, input *model.Resetpwd) (string, error) {
	return ressetpwd(ctx, input)
}

func (r *mutationResolver) Forcepassreset(ctx context.Context, input *model.AdminPreset) (string, error) {
	return adminpassreset(ctx, input)
}

func (r *mutationResolver) FlagSubmit(ctx context.Context, input string) (string, error) {
	return submitFlag(ctx, input)
}

func (r *mutationResolver) AddChallenge(ctx context.Context, input *model.AddChallengeData) (string, error) {
	return addchallenge(ctx, input)
}

func (r *mutationResolver) EditChallenge(ctx context.Context, input *model.EditChallengeData) (string, error) {
	return editchallenge(ctx, input)
}

func (r *mutationResolver) Deletechallenge(ctx context.Context, id string) (string, error) {
	return DelChall(ctx, id)
}

func (r *mutationResolver) Deletefile(ctx context.Context, input *model.Delfile) (string, error) {
	return removefilechall(ctx, input)
}

func (r *mutationResolver) Challvisibility(ctx context.Context, input *model.Public) (string, error) {
	return changevisibility(ctx, input)
}

func (r *mutationResolver) Adminmanagement(ctx context.Context, input *model.Manager) (string, error) {
	return Updaterules(ctx, input)
}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	return user(ctx)
}

func (r *queryResolver) Oneuser(ctx context.Context, username string) (*model.User, error) {
	return auser(ctx, username)
}

func (r *queryResolver) Userid(ctx context.Context) (*model.User, error) {
	return UserbyID(ctx)
}

func (r *queryResolver) Challenge(ctx context.Context) ([]*model.Challenge, error) {
	return challenge(ctx)
}

func (r *queryResolver) Onechallenge(ctx context.Context, id string) (*model.Challenge, error) {
	return achallenge(ctx, id)
}

func (r *queryResolver) Login(ctx context.Context, input *model.LoginInput) (string, error) {
	return "Who are you and why are you so wise in the ways of science", nil
	// return loginuser(ctx, input)
}

func (r *queryResolver) Ping(ctx context.Context, name string) (string, error) {
	return "pong", nil
}

func (r *queryResolver) Scoreboard(ctx context.Context) ([]*model.User, error) {
	return rankin(ctx)
}

func (r *queryResolver) Userdata(ctx context.Context, id string) (*model.Fulluser, error) {
	return fullyuser(ctx, id)
}

func (r *queryResolver) Me(ctx context.Context) (*model.Fulluser, error) {
	return meeseeks(ctx)
}

func (r *queryResolver) Notify(ctx context.Context) ([]*model.Notification, error) {
	return get_notifications(ctx)
}

func (r *queryResolver) Onenotify(ctx context.Context, id string) (*model.Notification, error) {
	return OneNoti(ctx, id)
}

func (r *queryResolver) Frontendmanagement(ctx context.Context) (*model.Managerial, error) {
	return viewRules(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
