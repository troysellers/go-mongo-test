package imdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	//https://imdb-api.com/en/API/Search/k_psqhqimi/inception%202010
	os.Setenv("IMDB_API", "https://imdb-api.com")
	os.Setenv("IMDB_API_LANG", "en")
	os.Setenv("IMDB_API_KEY", "k_psqhqimi")
}
func TestSearchByString(t *testing.T) {

	movies, err := SearchByTitle("2022")
	assert.Nil(t, err)
	assert.True(t, len(movies.Results) > 0)
	for _, m := range movies.Results {
		err := m.GetTitle()
		assert.Nil(t, err)
	}
}
