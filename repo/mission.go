package repo

import "github.com/akash-scalent/shield-system/entity"

type MissionRepository interface {
	AddMission(*entity.Mission) error
	GetAllMissions() []*entity.Mission
}
