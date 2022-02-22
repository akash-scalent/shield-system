package console

import (
	"bufio"
	"fmt"
	"os"

	"github.com/akash-scalent/shield-system/entity"
	service "github.com/akash-scalent/shield-system/service/mission"
)

var missionService = service.MissionService{}

func (*consoleController) AddMission(mission *entity.Mission) error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nEnter name of mission:")
	scanner.Scan()
	mission.Name = scanner.Text()
	fmt.Print("\nEnter detail of mission:")
	scanner.Scan()
	mission.Detail = scanner.Text()
	fmt.Println()
	err := missionService.AddMission(mission)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Mission Created !")
	return nil
}

func (*consoleController) GetAllMissions() []*entity.Mission {
	return missionService.GetAllMissions()
}