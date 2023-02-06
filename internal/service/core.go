package service

import (
	"douyin-simple/api/core"
	"douyin-simple/internal/biz/usecase"
	"fmt"

	"github.com/kataras/iris/v12/context"
)

type CoreService struct {
	uc *usecase.CoreUsecase
}

func NewCoreService(uc *usecase.CoreUsecase) *CoreService {
	cs := new(CoreService)
	cs.uc = uc
	return cs
}

// GET: /douyin/feed
func (serv *CoreService) GetFeed(
	ctx *context.Context,
	req *core.DouyinFeedRequest,
) *core.DouyinFeedResponse {
	// todo
	return nil
}

// POST: /douyin/user/register
func (serv *CoreService) PostUserRegister(
	ctx *context.Context,
	req *core.DouyinUserRegisterRequest,
) *core.DouyinUserRegisterResponse {
	// todo

	// test
	msg := "tsf"
	fmt.Printf("Path: %v\n", ctx.Request().URL.Path)
	fmt.Printf("req: %v\n", req)
	ctx.Application().Logger().Error("芜湖起飞")
	return &core.DouyinUserRegisterResponse{
		StatusCode: 123,
		StatusMsg:  &msg,
		UserId:     321,
		Token:      "1233211234567",
	}
}

// POST: /douyin/user/login
func (serv *CoreService) PostUserLogin(
	ctx *context.Context,
	req *core.DouyinUserLoginRequest,
) *core.DouyinUserLoginResponse {
	// todo
	return nil
}

// GET: /douyin/user
func (serv *CoreService) GetUser(
	ctx *context.Context,
	req *core.DouyinUserRequest,
) *core.DouyinUserResponse {
	// todo
	return nil
}

// POST: /douyin/publish/action
func (serv *CoreService) PostPublishAction(
	ctx *context.Context,
	req *core.DouyinPublishActionRequest,
) *core.DouyinPublishActionResponse {
	// todo
	return nil
}

// GET: /douyin/publish/list
func (serv *CoreService) GetPublishList(
	ctx *context.Context,
	req *core.DouyinPublishListRequest,
) *core.DouyinPublishListResponse {
	// todo
	return nil
}
