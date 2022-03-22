package scryfall

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const file = "oracle-cards-20220322090414.json"

func TestCardImport(t *testing.T) {
	data, err := os.ReadFile(file)
	assert.Nil(t, err)

	var cards []Card

	err = json.Unmarshal(data, &cards)
	assert.Nil(t, err)
}

func TestDatabaseImport(t *testing.T) {
	_, err := LoadJSON(file)

	assert.Nil(t, err)
}

func BenchmarkSearch(b *testing.B) {
	d, err := LoadJSON(file)

	assert.Nil(b, err)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c := d.CardByName("")
		assert.Nil(b, c)
	}
}
