package middleware

import (
	"net/http"

	authProto "github.com/Mockird31/OnlineStore/gen/auth"
	ctxWorker "github.com/Mockird31/OnlineStore/internal/pkg/helpers/ctxWorker"
	model "github.com/Mockird31/OnlineStore/internal/pkg/model"
)

func IsAuth(authClient authProto.AuthServiceClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			sessionCookie, err := r.Cookie("session_id")
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			userProto, err := authClient.GetUserBySessionID(ctx, model.StringToSessionIDProto(sessionCookie.Value))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			user := model.AuthProtoUserToUser(userProto)

			newCtx := ctxWorker.UserToContext(ctx, user)

			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
