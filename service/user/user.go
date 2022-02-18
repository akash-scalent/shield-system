package service

import (
	"errors"

	"github.com/akash-scalent/shield-system/entity"
	"github.com/akash-scalent/shield-system/repo"
	"github.com/akash-scalent/shield-system/repo/memory"
)

type UserService struct {
}

var rep repo.UserRepository = memory.NewMemoryRepository()

func (*UserService) GetStatusOfAvengers() []entity.Avenger {
	return rep.GetAllAvengers()
}

func (*UserService) AddAvenger(avenger *entity.Avenger) (*entity.Avenger, error) {
	if avenger.Name == "" || avenger.HeroName == "" {
		return avenger, errors.New("avenger name or heroname cannot be empty")
	}
	return nil, nil
}
