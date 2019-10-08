package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func roll(dice string) int {
	diceNum, _ := getDiceNum(dice)
	diceSize, _ := getDiceSize(dice)

	var totalRoll int
	for i := 0; i < int(diceNum); i++ {
		totalRoll += rollDice(diceSize)
	}

	return totalRoll
}

func getDiceNum(dice string) (int64, error) {
	return strconv.ParseInt(dice[:strings.IndexByte(dice, 'd')], 10, 16)
}

func getDiceSize(dice string) (int64, error) {
	return strconv.ParseInt(dice[strings.IndexByte(dice, 'd')+1:], 10, 16)
}

func rollDice(diceSize int64) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(int(diceSize+1) - 1) + 1
}