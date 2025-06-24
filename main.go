package main

import (
	"fmt"
	"study/games"
	"study/scaner"
)

// Необходимые функции:
// 1. Установка/удаление игры (изменение флага Installed)
// 2. Поиск игр по разработчику
// 3. Подсчет игр по жанру
// 4. Вывод самых играемых игр (сортировка по PlayHours)
// 5. Рекомендации по жанрам (на основе мапы)

func main() {
	userLibrary := []*games.Game{}
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
			games.DeleteGame(userLibrary)
		}
	}

}
