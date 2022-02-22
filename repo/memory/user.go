package memory

import (
	"errors"
	"fmt"
	"strings"

	"github.com/akash-scalent/shield-system/constants"
	"github.com/akash-scalent/shield-system/entity"
)

func (*repoImpl) GetAllAvengers() []entity.Avenger {

	status := []entity.Avenger{{}}

	return status
}

func (*repoImpl) AddAvenger(avenger *entity.Avenger) error {
	for i := range memoryDatabase.avengers {
		if memoryDatabase.avengers[i].HeroName == avenger.HeroName {
			return constants.ErrDuplicateHeroName
		} else if memoryDatabase.avengers[i].Name == avenger.Name {
			return constants.ErrDuplicateNameOfUser
		}
	}
	avenger.ID = len(memoryDatabase.avengers) + 1
	memoryDatabase.avengers = append(memoryDatabase.avengers, avenger)
	return nil
}

func (*repoImpl) AssignMissionToAvenger(missionId int, avengersId []int) error {
	var mission *entity.Mission
	for i := range memoryDatabase.missions {
		if memoryDatabase.missions[i].ID == missionId {
			mission = memoryDatabase.missions[i]
		}
	}
	if mission == nil {
		return errors.New("mission not found")
	}
	for i := range memoryDatabase.avengers {
		for j, k := range avengersId {
			avenger := memoryDatabase.avengers[i]
			if avenger.ID == k {
				avenger.CurrentMissionIds = append(avenger.CurrentMissionIds, missionId)
				avengersId = append(avengersId[:j], avengersId[j+1:]...)
			}
		}
	}
	if len(avengersId) != 0 {
		idsOfAvengersNotFound := []string{}
		for _, id := range avengersId {
			idsOfAvengersNotFound = append(idsOfAvengersNotFound, string(id))
		}
		return fmt.Errorf("avengers with id %s not found", strings.Join(idsOfAvengersNotFound, ","))
	}
	return nil
}

func (*repoImpl) AddAbility(ability *entity.Ability) error {
	for i := range memoryDatabase.abilities {
		if memoryDatabase.abilities[i].ID == ability.ID {
			return constants.ErrDuplicateEntity
		}
	}
	memoryDatabase.abilities = append(memoryDatabase.abilities, ability)
	ability.ID = len(memoryDatabase.abilities)

	return nil
}

func (*repoImpl) AssignAbilityToAvenger(avengerId int, abilityId int) error {
	avenger := &entity.Avenger{}
	for i := range memoryDatabase.avengers {
		if memoryDatabase.avengers[i].ID == avengerId {
			avenger = memoryDatabase.avengers[i]
			break
		}
	}
	if avenger== nil {
		return constants.ErrEntityNotFound
	}
	ability := &entity.Ability{}
	for i := range memoryDatabase.abilities {
		if memoryDatabase.abilities[i].ID == abilityId {
			ability = memoryDatabase.abilities[i]
			break
		}
	}
	if ability== nil {
		return constants.ErrEntityNotFound
	}

	avenger.AbilityIds = append(avenger.AbilityIds, abilityId)
	return nil
}