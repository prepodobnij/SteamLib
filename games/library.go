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
	var genreArr []*Game
	finded := false //флаг что все игры найдены
	gameGenres := make(map[string]int)

	for _, game := range arr {
		gameGenres[game.Genre]++
	}

	fmt.Print("Введите жанр:")
	genre := scaner.ScannerText()

	for _, v := range arr { //проходим по массиву
		if genre == v.Genre { //если введенный жанр есть в массиве, то добавляем его в слайс
			genreArr = append(genreArr, v)
		}
	}

	for k, v := range gameGenres {
		if k == genre && len(gameGenres) != 0 { //если значение ключа равно введенному жанру и слайс не пустой то выполняется блок кода
			fmt.Printf("У вас %v игр с жанром %s:\n", v, k)
			ShowUserLibrary(genreArr)
			finded = true
		}
	}
	if !finded {
		fmt.Println("У вас нет игр данного жанра")
	}

	time.Sleep(time.Second * 2)

}

func PlayestGames(arr []*Game) {

	if len(arr) == 0 {
		fmt.Println("Библиотека игр пуста.")
		return
	}

	// Создаём копию слайса, чтобы не изменять оригинальный порядок
	playestArr := make([]*Game, len(arr))
	copy(playestArr, arr)

	// Сортировка пузырьком по убыванию часов
	n := len(playestArr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if playestArr[j].PlayHours < playestArr[j+1].PlayHours {
				// Меняем местами указатели на игры
				playestArr[j], playestArr[j+1] = playestArr[j+1], playestArr[j]
			}
		}
	}

	// Выводим отсортированный список
	ShowUserLibrary(playestArr)

}

func RecommendedGenre(arr []*Game) {

	if len(arr) == 0 {
		fmt.Println("Ваша библиотека пуста")
		time.Sleep(time.Second * 2)
		return
	}

	genreCount := make(map[string]int) //создается пустая мапа

	for _, game := range arr {
		genreCount[game.Genre]++ // Для каждого жанра увеличиваем счетчик
	} // Если жанр встречается впервые - создается новая запись

	maxCount := 0
	recommended := ""
	for genre, count := range genreCount {
		// Если текущий жанр имеет больше игр, чем текущий максимум
		if count > maxCount {
			maxCount = count    // Обновляем максимальное количество
			recommended = genre // Запоминаем жанр-лидер
		}
	}

	fmt.Printf("Рекомендованный жанр на основе ваших предпочтений: %s\n", recommended)
	time.Sleep(time.Second * 2)
}
