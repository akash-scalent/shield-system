package service

import (
	"github.com/akash-scalent/shield-system/entity"
	"github.com/akash-scalent/shield-system/repo"
	"github.com/akash-scalent/shield-system/repo/memory"
	"github.com/akash-scalent/shield-system/utils"
)

type MissionService struct {
}

var rep repo.MissionRepository = memory.NewMissionMemoryRepository()

func (*MissionService) AddMission(mission *entity.Mission) ( error) {
	if err := utils.CheckIfEmpty("detial",mission.Detail); err != nil {
		return err
	} else if err := utils.CheckIfEmpty("name",mission.Name); err != nil {
		return err
	}
	rep.AddMission(mission)
	return nil
}

func (*MissionService) GetAllMissions() []*entity.Mission {
	return rep.GetAllMissions()
}