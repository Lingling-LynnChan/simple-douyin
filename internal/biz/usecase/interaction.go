package usecase

import "douyin-simple/internal/biz/repo"

// 互动用例
type InteractionUsecase struct {
	userRepo    repo.UserRepo
	videoRepo   repo.VideoRepo
	commentRepo repo.CommentRepo
}

func NewInteractionUsecase(
	userRepo repo.UserRepo,
	videoRepo repo.VideoRepo,
	commentRepo repo.CommentRepo,
) *InteractionUsecase {
	return &InteractionUsecase{
		userRepo:    userRepo,
		videoRepo:   videoRepo,
		commentRepo: commentRepo,
	}
}

func (uc *InteractionUsecase) Thumb() {

}
