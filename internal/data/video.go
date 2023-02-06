package data

import (
	"context"
	"douyin-simple/internal/biz/repo"
)

type videoRepo struct {
	data *Data
}

func NewVideoRepo(data *Data) repo.VideoRepo {
	return &videoRepo{
		data: data,
	}
}

// 创建视频
func (r *videoRepo) Create(ctx context.Context, video *repo.Video) (*repo.Video, error) {
	// todo
	return nil, nil
}

// 通过 id 获取视频
// 此处需要带出 Author IsFavorite CommentCount FavoriteCount
// 但是 ThumbBy Comments 留空
// 若需要则通过 GetCommentsById GetThumbById 进行二次查询
func (r *videoRepo) GetById(ctx context.Context, id int64) (*repo.Video, error) {
	// todo
	return nil, nil
}

// 通过 id 获取评论列表
func (r *videoRepo) ListCommentsById(ctx context.Context, id int64) ([]*repo.Comment, error) {
	// todo
	return nil, nil
}

// 通过 id 获取点赞用户列表
func (r *videoRepo) ListThumbById(ctx context.Context, id int64) ([]*repo.User, error) {
	// todo
	return nil, nil
}
