package data

import (
	"douyin-simple/internal/biz/repo"
	"douyin-simple/internal/data/ent"
)

// 这个文件用于放置数据转换的相关方法

func transUserEntity(user *ent.User) *repo.User {
	return &repo.User{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}
}

func transUserSlice(users []*ent.User) []*repo.User {
	slice := make([]*repo.User, 0, len(users))
	for _, user := range users {
		slice = append(slice, transUserEntity(user))
	}
	return slice
}

func transVideoEntity(video *ent.Video) *repo.Video {
	return &repo.Video{
		ID:       video.ID,
		Title:    video.Title,
		PlayURL:  video.PlayURL,
		CoverURL: video.CoverURL,
	}
}

func transVideoSlice(videos []*ent.Video) []*repo.Video {
	slice := make([]*repo.Video, 0, len(videos))
	for _, video := range videos {
		slice = append(slice, transVideoEntity(video))
	}
	return slice
}

func transCommentEntity(comment *ent.Comment) *repo.Comment {
	return &repo.Comment{
		ID:         comment.ID,
		Content:    comment.Content,
		CreateDate: comment.CreateDate,
		OwnedUser:  transUserEntity(comment.Edges.OwnedUser),
		OwnedVideo: transVideoEntity(comment.Edges.OwnedVideo),
	}
}

func transCommentSlice(comments []*ent.Comment) []*repo.Comment {
	slice := make([]*repo.Comment, 0, len(comments))
	for _, comment := range comments {
		slice = append(slice, transCommentEntity(comment))
	}
	return slice
}

func transMessageEntity(msg *ent.Message) *repo.Message {
	return &repo.Message{
		ID:         msg.ID,
		Content:    msg.Content,
		CreateDate: msg.Content,
		SendUser:   transUserEntity(msg.Edges.SendBy),
		RecvUser:   transUserEntity(msg.Edges.RecvBy),
	}
}

func transMessageSlice(msgs []*ent.Message) []*repo.Message {
	slice := make([]*repo.Message, 0, len(msgs))
	for _, msg := range msgs {
		slice = append(slice, transMessageEntity(msg))
	}
	return slice
}

// 用户转换 ent->repo 具有外键的转换
// 没有带入 is_follow
func transUserAll(user *ent.User) *repo.User {
	u := transUserEntity(user)
	u.Videos = transVideoSlice(user.Edges.Videos)
	u.Likes = transVideoSlice(user.Edges.Likes)
	u.Followers = transUserSlice(user.Edges.Followers)
	u.Following = transUserSlice(user.Edges.Following)
	u.FollowCount = int64(len(user.Edges.Following))
	u.FollowerCount = int64(len(user.Edges.Followers))
	return u
}
