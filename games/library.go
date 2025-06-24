package games

import (
	"bufio"
	"fmt"
	"os"
)

type Game struct {
	Title     string
	Developer string
	Year      int
	Genre     string
	Installed bool
	PlayHours float64
}

func ShowUserLibrary(g []*Game) {
	scaner := bufio.NewScanner(os.Stdin)
	for _, v := range g {
		fmt.Printf("\nНазвание: %v\nразработчик: %v\nгод выпуска: %v\nЖанр: %v\nУстановленна: %v\nЧасов в игре: %v\n", v.Title, v.Developer, v.Year, v.Genre, v.Installed, v.PlayHours)
		fmt.Println("____________________________________________________________")
	}

	for {
		fmt.Print("Выйти в главное меню (да/нет): ")
		scaner.Scan()
		exit := scaner.Text()
		if exit == "да" {
			break
		} else {
			continue
		}
	}

}

func FindByDev() {

}
