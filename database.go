package scryfall

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Database struct {
	cards []Card
}

func LoadJSON(path string) (*Database, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cards []Card

	err = json.Unmarshal(data, &cards)
	if err != nil {
		return nil, err
	}

	return &Database{cards}, nil
}

func (d *Database) CardByName(name string) *Card {
	for _, card := range d.cards {
		if card.Name == name {
			return &card
		}
	}
	for _, card := range d.cards {
		if strings.HasPrefix(card.Name, name+" // ") {
			return &card
		}
	}
	return nil
}

func (d *Database) CardByFuzzyName(name string) []Card {
	var cards []Card
	for _, card := range d.cards {
		if fuzzy.Match(name, card.Name) {
			cards = append(cards, card)
		}
	}
	return cards
}
