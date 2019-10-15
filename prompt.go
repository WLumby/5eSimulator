package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
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

	prompt = promptui.Select{
		Label: "Select Enemy",
		Items: items,
	}

	_, enemyPath, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	playerPath = baseSheetsDirectory + playerPath
	enemyPath = baseSheetsDirectory + enemyPath

	runSim(playerPath, enemyPath)
}
