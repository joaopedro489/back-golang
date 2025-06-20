package auth

import (
	"context"
	"strconv"
)

type AuthContext interface {
	GetUserID() int
}

const userContextKey = contextKey("authContext")

type UserCtx struct {
	UserID int
}

func (u *UserCtx) GetUserID() int {
	return u.UserID
}

func WithAuthContext(ctx context.Context, userID string) context.Context {
	id, err := strconv.Atoi(userID)
	if err != nil {
		panic(err)
	}

	return context.WithValue(ctx, userContextKey, &UserCtx{UserID: id})
}

func FromContext(ctx context.Context) (AuthContext, bool) {
	ac, ok := ctx.Value(userContextKey).(AuthContext)
	return ac, ok
}
