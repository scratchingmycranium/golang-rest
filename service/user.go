package service

import (
	"github.com/scratchingmycranium/golang-rest/interfaces"
	"github.com/scratchingmycranium/golang-rest/model"
)

type UserSvc struct {
	userRepo interfaces.IUserRepo
}

func InitUserService(userRepo interfaces.IUserRepo) interfaces.IUserService {
	return &UserSvc{userRepo: userRepo}
}

func (svc *UserSvc) Get() (model.User, error) {
	return svc.userRepo.Get()
}
