package service

import (
	"testing"

	"github.com/akash-scalent/shield-system/constants"
	"github.com/akash-scalent/shield-system/entity"
)

func TestMissionService_AddMission(t *testing.T) {
	type args struct {
		mission *entity.Mission
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Add mission with missing name",args{&entity.Mission{Detail: ""}},true},
		{"Add mission with missing detail",args{&entity.Mission{Name: ""}},true},
		{"Mission detail is status is Assigned",args{&entity.Mission{Name: "Test",Detail: "test detail"}},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MissionService{}
			if err := m.AddMission(tt.args.mission); (err != nil) != tt.wantErr {
				t.Errorf("MissionService.AddMission() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.args.mission.Status != "" && tt.args.mission.Status != constants.Assigned{
				t.Errorf("Mission status should be %v, not %v",constants.Assigned,tt.args.mission.Status)
			}
		})
	}
}
