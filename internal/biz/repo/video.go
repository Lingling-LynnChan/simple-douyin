package repo

import "context"

type Video struct {
	// id
	ID int64 `json:"id,omitempty"`
	// 视频播放地址
	PlayURL string `json:"play_url,omitempty"`
	// 视频封面地址
	CoverURL string `json:"cover_url,omitempty"`
	// 视频标题
	Title string `json:"title,omitempty"`
	// author: 作者;
	Author *User `json:"author,omitempty"`
	// thumb_by: 视频被哪些人点赞
	ThumbBy []*User `json:"thumb_by,omitempty"`
	// comments: 视频的评论
	Comments []*Comment `json:"comments,omitempty"`

	/*额外的*/
	// 是否被当前登录用户点赞
	IsFavorite bool `json:"is_favorite,omitempty"`
	// 视频的评论总数
	CommentCount int64 `json:"comment_count,omitempty"`
	// 视频的点赞总数
	FavoriteCount int64 `json:"favorite_count,omitempty"`
}

type VideoRepo interface {
	// 创建视频
	Create(ctx context.Context, video *Video) (*Video, error)
	// 通过 id 获取视频
	// 此处需要带出 Author IsFavorite CommentCount FavoriteCount
	// 但是 ThumbBy Comments 留空
	// 若需要则通过 GetCommentsById GetThumbById 进行二次查询
	GetById(ctx context.Context, id int64) (*Video, error)
	// 通过 id 获取评论列表
	ListCommentsById(ctx context.Context, id int64) ([]*Comment, error)
	// 通过 id 获取点赞用户列表
	ListThumbById(ctx context.Context, id int64) ([]*User, error)
}
