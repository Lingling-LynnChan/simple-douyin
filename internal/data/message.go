package data

import (
	"context"
	"douyin-simple/internal/biz/repo"
)

type messageRepo struct {
	data *Data
}

func NewMessageRepo(data *Data) repo.MessageRepo {
	return &messageRepo{
		data: data,
	}
}

// 发送消息
func (r *messageRepo) Send(ctx context.Context, msg repo.Message) (bool, error) {
	// todo
	return true, nil
}

// 获取聊天记录
func (r *messageRepo) List(ctx context.Context, user1, user2 int64) ([]*repo.Message, error) {
	// todo
	return nil, nil
}
