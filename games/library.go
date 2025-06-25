package games

import (
	"fmt"
	"strings"
	"study/scaner"
	"time"
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
	if len(g) == 0 {
		fmt.Println("Ваша библиотека пуста :(")
		time.Sleep(time.Second * 2)
	} else {

		for _, v := range g {
			status := "Установлена"
			if !v.Installed {
				status = "Не установлена"
			}
			fmt.Printf("\nНазвание: %v\nразработчик: %v\nгод выпуска: %v\nЖанр: %v\nУстановленна: %v\nЧасов в игре: %v\n", v.Title, v.Developer, v.Year, v.Genre, status, v.PlayHours)
			fmt.Println("____________________________________________________________")
		}

		for {
			fmt.Print("Выйти в главное меню (да/нет): ")
			if scaner.ScannerText() == "да" {
				break
			}
		}
	}

}

func FindByDev(arr []*Game) {
	fmt.Print("Введите разработчика:")
	dev := scaner.ScannerText()
	var devArr []*Game

	for _, v := range arr {
		if strings.EqualFold(dev, v.Developer) {
			devArr = append(devArr, v)
		} else {
			continue
		}
	}

	if len(devArr) == 0 {
		fmt.Printf("Игр от разработчика %s не найдено\n", dev)
		time.Sleep(time.Second * 2)
	} else {
		ShowUserLibrary(devArr)
	}

}

func CountByGenre(arr []*Game) {
	gameGenres := map[string]int{
		"action":           0,
		"action-adventury": 0,
		"adventury":        0,
		"role-playing":     0,
		"strategy":         0,
		"simulation":       0,
		"RPG":              0,
	}

	for _, v := range arr {
		if _, ok := gameGenres[v.Genre]; ok {
			gameGenres[v.Genre] += 1
		}
	}
	for k, v := range gameGenres {
		if v > 0 {
			fmt.Println(k, v)
		}
	}
	time.Sleep(time.Second * 2)

}
