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

func TestDatabaseImport(t *testing.T) {
	d := Database{}

	err := d.LoadJSON("oracle-cards-20220321090421.json")

	assert.Nil(t, err)
}

func BenchmarkSearch(b *testing.B) {
	d := Database{}

	err := d.LoadJSON("oracle-cards-20220321090421.json")

	assert.Nil(b, err)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c := d.CardByName("")
		assert.Nil(b, c)
	}
}
