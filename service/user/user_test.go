package service

import (
	"reflect"
	"testing"

	"github.com/akash-scalent/shield-system/entity"
	service "github.com/akash-scalent/shield-system/service/mission"
)

func TestUserService_GetStatusOfAvengers(t *testing.T) {
	tests := []struct {
		name string
		u    *UserService
		want []entity.Avenger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{}
			if got := u.GetStatusOfAvengers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetStatusOfAvengers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_AddAvenger(t *testing.T) {
	type args struct {
		avenger *entity.Avenger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"HeroName arg is empty", args{&entity.Avenger{User: entity.User{Name: "akash"}}}, true},
		{"Name arg is empty", args{&entity.Avenger{HeroName: "test"}}, true},
		{"Name and HeroName are supplied", args{&entity.Avenger{HeroName: "Akash", User: entity.User{Name: "Akash"}}}, false},
		{"Adding avenger", args{&entity.Avenger{
			User:              entity.User{Name: "akash"},
			HeroName:          "Sky",
			AbilityIds:        []int{1},
			CurrentMissionIds: []int{1},
			MissionsCompleted: 0,
		}}, false},
		{"Adding avenger with same Name", args{&entity.Avenger{
			User:              entity.User{Name: "akash"},
			HeroName:          "Sky2",
			AbilityIds:        []int{1},
			CurrentMissionIds: []int{1},
			MissionsCompleted: 0,
		}}, true},
		{"Adding avenger with same hero name", args{&entity.Avenger{
			User:              entity.User{Name: "akash1"},
			HeroName:          "Sky",
			AbilityIds:        []int{1},
			CurrentMissionIds: []int{1},
			MissionsCompleted: 0,
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{}
			if err := u.AddAvenger(tt.args.avenger); (err != nil) != tt.wantErr {
				t.Errorf("UserService.AddAvenger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_AssignMissionToAvenger(t *testing.T) {
	type args struct {
		missionId  int
		avengerIds []int
	}

	u := &UserService{}
	testAvenger := &entity.Avenger{User: entity.User{Name: "test3"},HeroName: "heroTest"}
	u.AddAvenger(testAvenger)

	m := &service.MissionService{}
	mission := &entity.Mission{Name: "Test",Detail: "Test"}
	m.AddMission(mission)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Mission does not exists",args{-9999,[]int{testAvenger.ID}},true,
			
		},{
			"Avenger does not exists",args{mission.ID,[]int{-9999}},true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := u.AssignMissionToAvenger(tt.args.missionId, tt.args.avengerIds); (err != nil) != tt.wantErr {
				t.Errorf("UserService.AssignMissionToAvenger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
