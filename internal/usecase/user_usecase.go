package usecase

import "github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUseCase {
	return &userUsecase{
		userRepo: ur,
	}
}

// sync from auth
func (u *userUsecase) SyncUserFromAuth(user *domain.User) error {
	return u.userRepo.CreateOrUpdate(user)
}
