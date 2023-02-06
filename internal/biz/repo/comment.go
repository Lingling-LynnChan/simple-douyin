package repo

import "context"

type Comment struct {
	// id
	ID int64 `json:"id,omitempty"`
	// 评论内容
	Content string `json:"content,omitempty"`
	// 评论发布日期 mm-dd
	CreateDate string `json:"create_date,omitempty"`
	// 该条评论属于哪个用户
	OwnedUser *User `json:"owned_user,omitempty"`
	// 该条评论属于哪个视频
	OwnedVideo *Video `json:"owned_video,omitempty"`
}

type CommentRepo interface {
	// 创建评论 && 删除评论
	// action ture 发布评论 false 删除评论
	Action(ctx context.Context, comment *Comment, action bool) (*Comment, error)
}
