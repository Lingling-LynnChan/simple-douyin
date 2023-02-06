package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("id"),
		field.String("content").Comment("消息内容"),
		field.String("create_date").Comment("评论发布日期 mm-dd"),

		field.Int64("user_id").Comment("发送者id"),
		field.Int64("recv_id").Comment("接收者id"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("send_by", User.Type).Comment("用于指定发送者,相当于外键约束").
			Required().Unique().Field("user_id"),
		edge.To("recv_by", User.Type).Comment("用于指定接收者,相当于外键约束").
			Required().Unique().Field("recv_id"),
		// edge.From("user_send", User.Type).Required().
		// 	Unique().Field("send_id").Comment("该条消息由谁发出"),
		// edge.From("user_recv", User.Type).Required().
		// 	Unique().Field("recv_id").Comment("该条消息由谁接收"),
	}
}
