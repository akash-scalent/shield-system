package entity

type Ability struct {
	Name string
}

type Status string

const (
	Assigned  Status = "Assigned"
	Completed Status = "Completed"
)

type Missions struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Status
}

type Avenger struct {
	User
	HeroName          string    `json:"heroName"`
	Abilities         []Ability `json:"abilities"`
	CurrentMissions   Missions  `json:"missions"`
	MissionsCompleted int
}
