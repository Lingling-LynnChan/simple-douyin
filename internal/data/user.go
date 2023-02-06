package data

import (
	"context"
	"douyin-simple/internal/biz/repo"
	"douyin-simple/internal/data/ent/user"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) repo.UserRepo {
	return &userRepo{
		data: data,
	}
}

// 创建用户
func (r *userRepo) Create(ctx context.Context, user *repo.User) (*repo.User, error) {
	// todo
	// hash pwd

	ent, err := r.data.db.User.Create().
		SetName(user.Name).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	// r.data.db.Video.Query
	return transUserEntity(ent), nil
}

// 通过 id 获取 user 信息
// 此处需要带出 IsFollow FollowCount FollowerCount
// 但是 Following Followers Friends 留空
func (r *userRepo) GetById(ctx context.Context, id int64) (*repo.User, error) {
	// todo
	r.data.db.User.Query().Where(user.ID(id))
	// r.data.db.User.QueryFriends().Count()
	return nil, nil
}

// 根据账户和密码拿到 user 信息
func (r *userRepo) GetByAcctPwd(ctx context.Context, name string, pwd string) (*repo.User, error) {
	// todo
	return nil, nil
}

// 通过 id 获取关注列表
func (r *userRepo) ListFollowingById(ctx context.Context, id int64) ([]*repo.User, error) {
	// todo
	return nil, nil
}

// 通过 id 获取粉丝列表
func (r *userRepo) ListFollowerById(ctx context.Context, id int64) ([]*repo.User, error) {
	// todo
	return nil, nil
}

// 通过 id 获取好友列表
func (r *userRepo) ListFriendsById(ctx context.Context, id int64) ([]*repo.User, error) {
	// todo
	return nil, nil
}

// 通过 id 获取用户点赞过的视频
func (r *userRepo) ListLikesById(ctx context.Context, id int64) ([]*repo.Video, error) {
	// todo
	return nil, nil
}

// 通过 id 获取用户发布的视频
func (r *userRepo) ListVideosById(ctx context.Context, id int64) ([]*repo.Video, error) {
	// todo
	return nil, nil
}

// 点赞操作
// action ture 点赞 false 取消点赞
func (r *userRepo) Thumb(ctx context.Context, userId, videioId int64, action bool) (bool, error) {
	// todo
	return true, nil
}
