package repo

import "github.com/akash-scalent/shield-system/entity"

type UserRepository interface {
	GetAllAvengers() []entity.Avenger
	AddAvenger( *entity.Avenger) error
	AssignMissionToAvenger(missionId int, avengersId []int) error
	AddAbility(ability *entity.Ability) error
	AssignAbilityToAvenger(avengerId int, abilityId int) error
}
