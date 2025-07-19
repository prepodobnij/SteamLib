package games

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"study/scaner"
)

func AddNewGame(arr []*Game) []*Game {
	var game Game
	scaner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите название игры: ")
	if scaner.Scan() {
		game.Title = scaner.Text()
	}
	if game.Title == "" {
		fmt.Println("название игры не может быть пустым")
		return arr
	}
	fmt.Print("Введите разработчика игры: ")
	if scaner.Scan() {
		game.Developer = scaner.Text()
	}
	if game.Developer == "" {
		fmt.Println("разработчик игры не может быть пустым")
		return arr
	}
	fmt.Print("Введите год выпуска игры: ")
	if scaner.Scan() {
		text := scaner.Text()
		year, err := strconv.Atoi(text)
		if err != nil || year < 1958 || year > 2025 {
			fmt.Println("Ошибка: неверный год выпуска. Введите число от 1958 до 2025.")
			return arr
		}
		game.Year = year
	}

	fmt.Print("Введите жанр игры: ")
	if scaner.Scan() {
		game.Genre = scaner.Text()
	}
	if game.Genre == "" {
		fmt.Println("Жанр не может быть пустым")
		return arr
	}
	game.Installed = true
	game.PlayHours = 0.0

	newGame := Game{
		Title:     game.Title,
		Developer: game.Developer,
		Year:      game.Year,
		Genre:     game.Genre,
		Installed: game.Installed,
		PlayHours: game.PlayHours,
	}

	arr = append(arr, &newGame)
	return arr
}

func DeleteGame(arr []*Game) {
	fmt.Print("Введите название игры которой хотите удалить: ")
	text := scaner.ScannerText() 

	found := false
	for _, v := range arr {
		if strings.EqualFold(text, v.Title) {
			v.Installed = false
			found = true
		}
	}

	if !found {
		fmt.Printf("Игра '%s' не найдена\n", text)
	} else {
		fmt.Printf("Игра '%s' успешно удалена\n", text)
	}

}
