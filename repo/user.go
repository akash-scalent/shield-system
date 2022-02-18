package repo

import "github.com/akash-scalent/shield-system/entity"

type UserRepository interface {
	GetAllAvengers() []entity.Avenger
}
