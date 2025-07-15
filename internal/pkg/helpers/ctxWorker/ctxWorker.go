package ctxWorker

import (
	"context"

	"github.com/Mockird31/OnlineStore/internal/pkg/model"
)

type UserKey struct{}

func UserToContext(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, UserKey{}, user)
}

func UserFromContext(ctx context.Context) (*model.User, bool) {
	user, ok := ctx.Value(UserKey{}).(*model.User)
	if !ok {
		return nil, false
	}
	return user, true
}
