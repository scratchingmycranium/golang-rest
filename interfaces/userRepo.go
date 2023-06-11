package interfaces

import "github.com/scratchingmycranium/golang-rest/model"

type IUserRepo interface {
	Get() (model.User, error)
}
