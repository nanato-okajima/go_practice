package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNumber() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func omikuji() {
	uranai := []string{"大凶", "中吉", "小吉", "大吉", "凶"}
	fmt.Println("1 ~ 4までの数値を入力してください")
	num, err := getNumber()
	if err != nil {
		log.Fatal(uranai[num])
	}

	fmt.Println(uranai[num])
}
