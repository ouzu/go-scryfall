package scryfall

import (
	"encoding/json"
	"os"
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Database struct {
	cards []Card
	names []string
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

	names := make([]string, len(cards))

	for i, card := range cards {
		names[i] = card.Name
	}

	return &Database{cards, names}, nil
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
	results := fuzzy.RankFind(name, d.names)

	sort.Sort(results)

	cards := make([]Card, len(results))

	for i, result := range results {
		cards[i] = d.cards[result.OriginalIndex]
	}

	return cards
}
