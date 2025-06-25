package main

import (
	"fmt"
	"study/games"
	"study/scaner"
	"time"
)

// Необходимые функции:
// 1. Установка/удаление игры (изменение флага Installed)
// 2. Поиск игр по разработчику
// 3. Подсчет игр по жанру
// 4. Вывод самых играемых игр (сортировка по PlayHours)
// 5. Рекомендации по жанрам (на основе мапы)

func main() {
	userLibrary := []*games.Game{}

	/* game1 := games.Game{
		Title:     "GTA 5",
		Developer: "Rockstar",
		Year:      2013,
		Genre:     "Open world",
		Installed: true,
		PlayHours: 222.6,
	}
	game2 := games.Game{
		Title:     "Read Dead Redemption 2",
		Developer: "Rockstar",
		Year:      2018,
		Genre:     "Open world",
		Installed: true,
		PlayHours: 117.8,
	}
	game3 := games.Game{
		Title:     "The Witcher 3",
		Developer: "CDPR",
		Year:      2015,
		Genre:     "RPG",
		Installed: true,
		PlayHours: 240.4,
	}

	userLibrary = append(userLibrary, &game1)
	userLibrary = append(userLibrary, &game2)
	userLibrary = append(userLibrary, &game3) */

	for {
		fmt.Println("Добро пожаловать в Steam!")
		fmt.Println("--------Меню--------")
		commands := [7]string{
			"Моя библиотека",
			"Установить игру",
			"Удалить игру",
			"Поиск игры по разработчику",
			"Подсчет игр по жанрам",
			"Вывод самых играемых игр",
			"Рекомендации по жанрам",
		}

		for i, v := range commands {
			fmt.Printf("%v) %v\n", i+1, v)
		}

		fmt.Print("Введите номер нужной команды:")
		userChoice := scaner.ScannerStart()
		if userChoice == -1 {
			continue
		}

		switch userChoice {
		case 1:
			games.ShowUserLibrary(userLibrary)
		case 2:
			userLibrary = games.AddNewGame(userLibrary)
		case 3:
			if len(userLibrary) == 0 {
				fmt.Println("Ваша библиотека пуста")
				time.Sleep(time.Second * 2)
			} else {
				games.DeleteGame(userLibrary)
			}

		case 4:
			games.FindByDev(userLibrary)
		}
	}

}
