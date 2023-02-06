package repo

import "context"

type Message struct {
	// id
	ID int64 `json:"id,omitempty"`
	// 评论内容
	Content string `json:"content,omitempty"`
	// 评论发布日期 mm-dd
	CreateDate string `json:"create_date,omitempty"`
	// 发送者
	SendUser *User `json:"user_id,omitempty"`
	// 接收者
	RecvUser *User `json:"recv_id,omitempty"`
}

type MessageRepo interface {
	// 发送消息
	Send(ctx context.Context, msg Message) (bool, error)
	// 获取聊天记录
	List(ctx context.Context, user1, user2 int64) ([]*Message, error)
}
