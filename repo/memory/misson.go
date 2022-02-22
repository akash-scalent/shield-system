package memory

import (
	"github.com/akash-scalent/shield-system/constants"
	"github.com/akash-scalent/shield-system/entity"
)

func (*repoImpl) AddMission(mission *entity.Mission) error {
	mission.ID = len(memoryDatabase.missions) + 1
	mission.Status = constants.UnAssigned
	memoryDatabase.missions = append(memoryDatabase.missions, mission)
	return nil
}

func (*repoImpl) GetAllMissions() []*entity.Mission {
	return memoryDatabase.missions
}
