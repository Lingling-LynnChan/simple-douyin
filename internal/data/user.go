package data

import (
	"context"
	"douyin-simple/internal/biz/repo"
)

type userRepo struct {
}

func NewUserRepo() repo.UserRepo {
	return &userRepo{}
}

func (r *userRepo) Create(ctx context.Context) *repo.User {
	return &repo.User{}
}
func (r *userRepo) GetUserByLogin(ctx context.Context) *repo.User {
	return &repo.User{}
}
