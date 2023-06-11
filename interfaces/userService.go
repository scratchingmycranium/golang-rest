package interfaces

import "github.com/scratchingmycranium/golang-rest/model"

type IUserService interface {
	Get() (model.User, error)
}
