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

// POST: /douyin/user/register
func (serv *CoreService) PostUserRegister(ctx *context.Context, req *core.DouyinUserRegisterRequest) *core.DouyinUserRegisterResponse {
	msg := "tsf"
	fmt.Printf("req: %v\n", req)
	ctx.Application().Logger().Error("芜湖起飞")
	return &core.DouyinUserRegisterResponse{
		StatusCode: 123,
		StatusMsg:  &msg,
		UserId:     321,
		Token:      "1233211234567",
	}
}
