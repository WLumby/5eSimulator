package main

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"strconv"
)

var (
	baseSheetsDirectory = "./sheets/"
	items []string
)

func main() {
	files, _ := ioutil.ReadDir(baseSheetsDirectory)

	for _, f := range files {
		items = append(items, f.Name())
	}

	prompt := promptui.Select{
		Label: "Select Player",
		Items: items,
	}

	_, playerPath, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	prompt = promptui.Select{
		Label: "Select Enemy",
		Items: items,
	}

	_, enemyPath, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	promptP := promptui.Prompt{
		Label:    "Sim Iterations",
		Validate: validate,
	}

	iterations, err := promptP.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	playerPath = baseSheetsDirectory + playerPath
	enemyPath = baseSheetsDirectory + enemyPath

	iterationsVal, err := strconv.ParseInt(iterations, 10,64)
	runSim(playerPath, enemyPath, int(iterationsVal))
}

func validate(input string) error {
	_, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return errors.New("Invalid number")
	}
	return nil
}
