package console

import (
	"bufio"
	"fmt"
	"os"

	"github.com/akash-scalent/shield-system/entity"
)

// TODO: Add interface support for controllers
type consoleController struct {
	
}

func RunConsole()  {
	fmt.Println(`				=======------S.H.I.E.L.D ------=========
					Welcome to Captain Fury.`)
	console := consoleController{}
	scanner := bufio.NewScanner(os.Stdin)
	for  {
		fmt.Println(`
1.	Add Avenger
2.	Add mission
3.	Check the missions
4.	Assign mission to Avengers
5.	Check mission’s details
6.	Check Avenger’s details
7.	Update Mission’s status
8.	List Avengers`)
fmt.Print("\nEnter the option:")
		scanner.Scan()
		// option := scanner.Text()
		switch scanner.Text() {
		case "1":
			avenger := &entity.Avenger{}
			err := console.AddAvenger(avenger)
			// TODO: How to minimize this repetitive code
			if err != nil {
				continue
			}
		case "2":
			// TODO: Should we handle error here or inside specific controller
			mission := &entity.Mission{}
			err := console.AddMission(mission)
			if err != nil {
				continue
			}
		case "3":
			missons := console.GetAllMissions()
			if len(missons)==0 {
				fmt.Println("No Mission has been assigned to an Avenger.")
			}
			fmt.Println("Mission Name		Status		Avengers")
			for _, mission := range missons {
				fmt.Printf("%s	%s Avengers",mission.Name,mission.Status)
			}
			fmt.Println("")
		case "4":
			
		default:
			fmt.Println("Invalid Input")
		}
	}
}