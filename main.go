package main

import (
	"fmt"
	"study/games"
	"study/scaner"
	"time"
)

// Необходимые функции:
// 1. Установка/удаление игры (изменение флага Installed)   done
// 2. Поиск игр по разработчику							    done
// 3. Подсчет игр по жанру									done
// 4. Вывод самых играемых игр (сортировка по PlayHours)	done
// 5. Рекомендации по жанрам (на основе мапы)				not

func main() {
	userLibrary := []*games.Game{}

	game1 := games.Game{
		Title:     "GTA 5",
		Developer: "Rockstar",
		Year:      2013,
		Genre:     "экшн",
		Installed: true,
		PlayHours: 222.6,
	}
	game2 := games.Game{
		Title:     "Read Dead Redemption 2",
		Developer: "Rockstar",
		Year:      2018,
		Genre:     "экшн-приключение",
		Installed: true,
		PlayHours: 117.8,
	}
	game3 := games.Game{
		Title:     "The Witcher 3",
		Developer: "CDPR",
		Year:      2015,
		Genre:     "РПГ",
		Installed: true,
		PlayHours: 240.4,
	}
	game4 := games.Game{
		Title:     "Clair Obscure: Expidition 33",
		Developer: "Sandfall Interactive",
		Year:      2025,
		Genre:     "ЖРПГ",
		Installed: true,
		PlayHours: 22.5,
	}
	game5 := games.Game{
		Title:     "Total war Rome 2",
		Developer: "Creative Assembly",
		Year:      2016,
		Genre:     "стратегия",
		Installed: true,
		PlayHours: 88.7,
	}
	game6 := games.Game{
		Title:     "Total war Atilla",
		Developer: "Creative Assembly",
		Year:      2018,
		Genre:     "стратегия",
		Installed: true,
		PlayHours: 12.7,
	}

	userLibrary = append(userLibrary, &game1)
	userLibrary = append(userLibrary, &game2)
	userLibrary = append(userLibrary, &game3)
	userLibrary = append(userLibrary, &game4)
	userLibrary = append(userLibrary, &game5)
	userLibrary = append(userLibrary, &game6)

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
		case 5:
			games.CountByGenre(userLibrary)
		case 6:
			games.PlayestGames(userLibrary)
		case 7:
			games.RecommendedGenre(userLibrary)
		}
	}

}
