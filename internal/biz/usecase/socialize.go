package usecase

import "douyin-simple/internal/biz/repo"

// 社交用例
type SocializeUsecase struct {
	userRepo repo.UserRepo
	msgRepo  repo.MessageRepo
}

func NewSocializeUsecase(
	userRepo repo.UserRepo,
	msgRepo repo.MessageRepo,
) *SocializeUsecase {
	return &SocializeUsecase{
		userRepo: userRepo,
		msgRepo:  msgRepo,
	}
}
