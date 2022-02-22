package entity

import "github.com/akash-scalent/shield-system/constants"

type Ability struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Mission struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
	constants.Status
}

type Avenger struct {
	User
	HeroName          string `json:"heroName"`
	AbilityIds        []int  `json:"abilityIds"`
	CurrentMissionIds []int  `json:"currentMissionIds"`
	MissionsCompleted int
}
