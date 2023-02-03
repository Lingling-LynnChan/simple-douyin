package usecase

import "douyin-simple/internal/biz/repo"

type CoreUsecase struct {
	userRepo repo.UserRepo
}

func NewCoreUsecase(userRepo repo.UserRepo) *CoreUsecase {
	return &CoreUsecase{
		userRepo: userRepo,
	}
}

func (uc *CoreUsecase) Register() {

}

func (uc *CoreUsecase) Login() {

}
