package memory

import "github.com/akash-scalent/shield-system/repo"

type repoImpl struct {
}

var memoryDatabase = struct{}{}

func NewMemoryRepository() repo.UserRepository {
	return &repoImpl{}
}

