package repo

import (
	"context"
)

type User struct {
	// id
	ID int64 `json:"id,omitempty"`
	// 用户名
	Name string `json:"name,omitempty"`
	// 密码
	Password string `json:"password,omitempty"`
	// likes: 用户点赞过的视频
	Likes []*Video `json:"likes,omitempty"`
	// videos: 用户发布的视频
	Videos []*Video `json:"videos,omitempty"`
	// friends: 好友;
	Friends []*User `json:"friends,omitempty"`
	// follower: 我的粉丝;
	Followers []*User `json:"followers,omitempty"`
	// following: 我的关注;
	Following []*User `json:"following,omitempty"`

	/*额外的*/
	// true-已关注，false-未关注
	IsFollow bool `json:"is_follow,omitempty"`
	// 关注总数
	FollowCount int64 `json:"follow_count,omitempty"`
	// 粉丝总数
	FollowerCount int64 `json:"follower_count,omitempty"`
}

type UserRepo interface {
	// 创建用户
	Create(ctx context.Context, user *User) (*User, error)
	// 通过 id 获取 user 信息
	// 此处需要带出 IsFollow FollowCount FollowerCount
	// 但是 Following Followers Friends 留空
	GetById(ctx context.Context, id int64) (*User, error)
	// 根据账户和密码拿到 user 信息
	GetByAcctPwd(ctx context.Context, name string, pwd string) (*User, error)
	// 通过 id 获取关注列表
	ListFollowingById(ctx context.Context, id int64) ([]*User, error)
	// 通过 id 获取粉丝列表
	ListFollowerById(ctx context.Context, id int64) ([]*User, error)
	// 通过 id 获取好友列表
	ListFriendsById(ctx context.Context, id int64) ([]*User, error)
	// 通过 id 获取用户点赞过的视频
	ListLikesById(ctx context.Context, id int64) ([]*Video, error)
	// 通过 id 获取用户发布的视频
	ListVideosById(ctx context.Context, id int64) ([]*Video, error)

	// 点赞操作
	// action ture 点赞 false 取消点赞
	Thumb(ctx context.Context, userId, videioId int64, action bool) (bool, error)
}
