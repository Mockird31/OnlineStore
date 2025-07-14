package ctxWorker

import "context"

type UserIDKey struct{}

func UserIDToContext(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, UserIDKey{}, userID)
}

func UserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey{}).(int64)
	if !ok {
		return 0, false
	}
	return userID, true
}
