package scryfall

import (
	"encoding/json"
	"os"
)

type Database struct {
	cards []Card
}

func (d *Database) LoadJSON(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &(d.cards))
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) CardByName(name string) *Card {
	for _, card := range d.cards {
		if card.Name == name {
			return &card
		}
	}
	return nil
}
