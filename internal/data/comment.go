package data

import (
	"context"
	"douyin-simple/internal/biz/repo"
)

type commentRepo struct {
	data *Data
}

func NewCommentRepo(data *Data) repo.CommentRepo {
	return &commentRepo{
		data: data,
	}
}

// 创建评论 && 删除评论
// action ture 发布评论 false 删除评论
func (r *commentRepo) Action(ctx context.Context, comment *repo.Comment, action bool) (*repo.Comment, error) {
	// todo
	return nil, nil
}
