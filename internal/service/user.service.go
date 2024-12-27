package service

import "github.com/thienchuong/golang-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// us user Service
func (us *UserService) GetInFoUserService() string {
	return us.userRepo.GetInFoUser()
}
