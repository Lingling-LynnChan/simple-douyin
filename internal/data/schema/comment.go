package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("id"),
		field.String("content").Comment("评论内容"),
		field.String("create_date").Comment("评论发布日期 mm-dd"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owned_video", Video.Type).Ref("comments").
			Required().Unique().Comment("该条评论属于哪个视频"),
		edge.From("owned_user", User.Type).Ref("comments").
			Required().Unique().Comment("该条评论属于哪个用户"),
	}
}
