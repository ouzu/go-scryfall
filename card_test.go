package scryfall

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardImport(t *testing.T) {
	data, err := os.ReadFile("oracle-cards-20220321090421.json")
	assert.Nil(t, err)

	var cards []Card

	err = json.Unmarshal(data, &cards)
	assert.Nil(t, err)
}
