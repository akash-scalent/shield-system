package memory

import (
	"github.com/akash-scalent/shield-system/entity"
	"github.com/akash-scalent/shield-system/repo"
)

type repoImpl struct {
}

var memoryDatabase = struct {
	avengers []*entity.Avenger
	missions []*entity.Mission
	abilities []*entity.Ability
}{}

func NewUserMemoryRepository() repo.UserRepository {
	return &repoImpl{}
}

func NewMissionMemoryRepository() repo.MissionRepository {
	return &repoImpl{}
}
