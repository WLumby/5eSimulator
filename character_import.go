package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

func getCharacter(filename string) (Character, error) {
	jsonFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return Character{}, errors.New(fmt.Sprintf("failed to read file %v, %v", filename, err))
	}

	var character Character
	err = json.Unmarshal(jsonFile, &character)

	if err != nil {
		return Character{}, errors.New(fmt.Sprintf("failed to unmarshal JSON, %v", err))
	}

	return character, nil
}
