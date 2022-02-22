package service

import (
	"github.com/akash-scalent/shield-system/constants"
	"github.com/akash-scalent/shield-system/entity"
	"github.com/akash-scalent/shield-system/repo"
	"github.com/akash-scalent/shield-system/repo/memory"
	"github.com/akash-scalent/shield-system/utils"
)

type UserService struct {
}

var rep repo.UserRepository = memory.NewUserMemoryRepository()

func (*UserService) GetStatusOfAvengers() []entity.Avenger {
	return rep.GetAllAvengers()
}

func (*UserService) AddAvenger(avenger *entity.Avenger)  error {
	if avenger.Name == "" || avenger.HeroName == "" {
		return  constants.ErrAvengerNameOrHeroNameEmpty
	}
  err:=	rep.AddAvenger(avenger)
	if err != nil {
		return  err
	}
	return  nil
}

func (*UserService) AssignMissionToAvenger(missionIds int, avengerIds []int ) error {
	if err := utils.CheckIfEmpty("missionIds",missionIds); err != nil {
		return err
	} else if err := utils.CheckIfEmpty("avengerIds",avengerIds); err != nil {
		return err
	} 
	err := rep.AssignMissionToAvenger(missionIds,avengerIds)
	if err != nil {
		return err
	}
	return nil
}