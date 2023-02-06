package service

import (
	"douyin-simple/api/interaction"
	"douyin-simple/internal/biz/usecase"

	"github.com/kataras/iris/v12/context"
)

type InteractionService struct {
	uc *usecase.InteractionUsecase
}

func NewInteractionService(uc *usecase.InteractionUsecase) *InteractionService {
	cs := new(InteractionService)
	cs.uc = uc
	return cs
}

// POST: /douyin/favorite/action
func (serv *InteractionService) PostFavoriteAction(
	ctx *context.Context,
	req *interaction.DouyinFavoriteActionRequest,
) *interaction.DouyinFavoriteActionResponse {
	// todo
	return nil
}

// GET: /douyin/favorite/list
func (serv *InteractionService) GetFavoriteList(
	ctx *context.Context,
	req *interaction.DouyinFavoriteListRequest,
) *interaction.DouyinFavoriteListResponse {
	// todo
	return nil
}

// POST: /douyin/comment/action
func (serv *InteractionService) GetCommentAction(
	ctx *context.Context,
	req *interaction.DouyinCommentActionRequest,
) *interaction.DouyinCommentActionRequest {
	// todo
	return nil
}

// GET: /douyin/comment/list
func (serv *InteractionService) GetCommentList(
	ctx *context.Context,
	req *interaction.DouyinCommentListRequest,
) *interaction.DouyinCommentListResponse {
	// todo
	return nil
}
