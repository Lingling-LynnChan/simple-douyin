package service

import (
	"douyin-simple/api/socialize"
	"douyin-simple/internal/biz/usecase"
	"fmt"

	"github.com/kataras/iris/v12/context"
)

type SocializeService struct {
	uc *usecase.SocializeUsecase
}

func NewSocializeService(uc *usecase.SocializeUsecase) *SocializeService {
	cs := new(SocializeService)
	cs.uc = uc
	return cs
}

// POST: /douyin/relation/action
func (serv *SocializeService) PostRelationAction(
	ctx *context.Context,
	req *socialize.DouyinRelationActionRequest,
) *socialize.DouyinRelationActionResponse {
	// todo
	return nil
}

// GET: /douyin/relation/follow/list
func (serv *SocializeService) GetRelationFollowList(
	ctx *context.Context,
	req *socialize.DouyinRelationFollowListRequest,
) *socialize.DouyinRelationFollowListResponse {
	// todo
	return nil
}

// GET: /douyin/relation/follower/list
func (serv *SocializeService) GetRelationFollowerList(
	ctx *context.Context,
	req *socialize.DouyinRelationFollowerListRequest,
) *socialize.DouyinRelationFollowerListResponse {
	// todo
	return nil
}

// GET: /douyin/relation/friend/list
func (serv *SocializeService) GetRelationFriendList(
	ctx *context.Context,
	req *socialize.DouyinRelationFriendListRequest,
) *socialize.DouyinRelationFriendListResponse {
	// todo
	return nil
}

// POST: /douyin/message/action
func (serv *SocializeService) PostMessageAction(
	ctx *context.Context,
	req *socialize.DouyinMassageActionRequest,
) *socialize.DouyinMassageActionResponse {
	// todo
	return nil
}

// GET: /douyin/message/chat
func (serv *SocializeService) GetMessageChat(
	ctx *context.Context,
	req *socialize.DouyinMessageChatRequest,
) *socialize.DouyinMessageChatResponse {
	// todo

	// test
	mag := "芜湖"
	fmt.Printf("req: %v\n", req)
	return &socialize.DouyinMessageChatResponse{
		StatusCode: 0,
		StatusMsg: &mag,
	}
}
