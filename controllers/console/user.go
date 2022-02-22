package console

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/akash-scalent/shield-system/constants"
	"github.com/akash-scalent/shield-system/entity"
	service "github.com/akash-scalent/shield-system/service/user"
)

var userService = service.UserService{}

func (*consoleController) AddAvenger(avenger *entity.Avenger) error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nEnter name of Avenger:")
	scanner.Scan()
	avenger.Name = scanner.Text()
	fmt.Print("\nEnter hero name of Avenger:")
	scanner.Scan()
	avenger.HeroName = scanner.Text()
	fmt.Println()
	err := userService.AddAvenger(avenger)
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrAvengerNameOrHeroNameEmpty):
			fmt.Println(err.Error())
		case errors.Is(err, constants.ErrDuplicateHeroName):
			fmt.Println(err.Error())
		case errors.Is(err, constants.ErrDuplicateNameOfUser):
			fmt.Println(err.Error())
		default:
			fmt.Println("Logger", err)
			fmt.Println("Something went wrong")
		}
		return err
	}
	fmt.Println("Avenger Created !")
	return nil
}
