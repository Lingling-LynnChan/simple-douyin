package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("id"),
		field.String("name").Unique().Comment("用户名"),
		field.String("password").Comment("密码"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// relation about
		edge.To("followers", User.Type).Comment("follower: 我的粉丝;").
			From("following").Comment("following: 我的关注;"),
		edge.To("friends", User.Type).Comment("friends: 好友;"),
		// video about
		edge.To("videos", Video.Type).Comment("videos: 用户发布的视频"),
		edge.From("likes", Video.Type).Ref("thumb_by").Comment("likes: 用户点赞过的视频"),
		edge.To("comments", Comment.Type).Comment("用户的评论"),
		// message about
		edge.To("recvs", User.Type).Comment("发送信息或接收信息").
			Through("messages", Message.Type),
		// edge.To("send_msgs", Message.Type).Comment("发送的消息"),
		// edge.To("recv_msgs", Message.Type).Comment("接收的消息"),
	}
}
