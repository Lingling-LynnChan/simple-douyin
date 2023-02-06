package usecase

import "douyin-simple/internal/biz/repo"

// 核心用例
type CoreUsecase struct {
	userRepo  repo.UserRepo
	videoRepo repo.VideoRepo
}

func NewCoreUsecase(
	userRepo repo.UserRepo,
	videoRepo repo.VideoRepo,
) *CoreUsecase {
	return &CoreUsecase{
		userRepo:  userRepo,
		videoRepo: videoRepo,
	}
}

func (uc *CoreUsecase) Register() {

}

func (uc *CoreUsecase) Login() {

}
