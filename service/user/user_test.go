package service

import (
	"reflect"
	"testing"

	"github.com/akash-scalent/shield-system/entity"
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
		// TODO: Add test cases.
		{"HeroName arg is empty", args{&entity.Avenger{User: entity.User{Name: "akash"}}}, true},
		{"Name arg is empty", args{&entity.Avenger{HeroName: "Akash"}}, true},
		{"Name and HeroName are supplied", args{&entity.Avenger{HeroName: "Akash", User: entity.User{Name: "akash"}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{}
			if _,err := u.AddAvenger(tt.args.avenger); (err != nil) != tt.wantErr {
				t.Errorf("UserService.AddAvenger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
