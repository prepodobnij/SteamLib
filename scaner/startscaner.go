package scaner

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ScannerStart() int {
	scaner := bufio.NewScanner(os.Stdin)

	if ok := scaner.Scan(); !ok {
		fmt.Println("Ошибка ввода")
		return -1 // Можно вернуть -1 как индикатор ошибки
	}

	text := scaner.Text()
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Ошибка преобразования, введите число!")

	}

	return num
}
