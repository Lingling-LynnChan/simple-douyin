package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("id"),
		field.String("play_url").Comment("视频播放地址"),
		field.String("cover_url").Comment("视频封面地址"),
		field.String("title").Comment("视频标题"),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Comment("author: 作者;").
			Ref("videos").Unique().Required(),
		edge.To("thumb_by", User.Type).Comment("thumb_by: 视频被哪些人点赞"),

		edge.To("comments", Comment.Type).Comment("comments: 视频的评论"),
	}
}
