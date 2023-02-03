package repo

import (
	"context"
	"douyin-simple/internal/data/ent"
)

type User struct {
	ent.User
}

type UserRepo interface {
	Create(ctx context.Context) *User
	GetUserByLogin(ctx context.Context) *User
}
