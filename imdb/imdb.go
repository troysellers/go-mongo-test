package imdb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/troysellers/mongotest/util"
)

type MovieList struct {
	Items []*Movie `json:"items"`
}

type Movies struct {
	SearchType string  `json:"searchType"`
	Expression string  `json:"expression"`
	Results    []Movie `json:"results"`
}
type Movie struct {
	Id              string `json:"id"`
	ResultType      string `json:"resultType"`
	Title           string `json:"title"`
	Image           string `'json:"image"`
	Description     string `json:"description"`
	Rank            string `json:"rank"`
	FullTitle       string `json:"fullTitle"`
	Year            string `json:"year"`
	Crew            string `json:"crew"`
	ImdbRating      string `json:"imDbRating"`
	ImdbRatingCount string `json:"imDbRatingCount"`
}

func GetList(id string) ([]*Movie, error) {
	url := fmt.Sprintf("%s/%s/API/IMDbList/%s/%s", os.Getenv("IMDB_API"), os.Getenv("IMDB_API_LANG"), os.Getenv("IMDB_API_KEY"), id)

	headers := make(map[string]string)

	_, bytes, err := util.DoHttp(url, nil, "GET", headers)
	if err != nil {
		return nil, err
	}
	movieList := MovieList{}
	if err := json.Unmarshal(bytes, &movieList); err != nil {
		return nil, err
	}
	return movieList.Items, nil
}

func GetTop250() ([]*Movie, error) {
	url := fmt.Sprintf("%s/%s/API/Top250Movies/%s", os.Getenv("IMDB_API"), os.Getenv("IMDB_API_LANG"), os.Getenv("IMDB_API_KEY"))
	headers := make(map[string]string)

	_, bytes, err := util.DoHttp(url, nil, "GET", headers)
	if err != nil {
		return nil, err
	}
	movieList := MovieList{}
	if err := json.Unmarshal(bytes, &movieList); err != nil {
		return nil, err
	}
	return movieList.Items, nil
}

func SearchByTitle(title string) (*Movies, error) {
	//https://imdb-api.com/en/API/Search/k_psqhqimi/inception%202010
	url := fmt.Sprintf("%s/%s/API/SearchMovie/%s/%s", os.Getenv("IMDB_API"), os.Getenv("IMDB_API_LANG"), os.Getenv("IMDB_API_KEY"), title)
	headers := make(map[string]string)

	_, bytes, err := util.DoHttp(url, nil, "GET", headers)
	if err != nil {
		return nil, err
	}
	movies := Movies{}
	if err := json.Unmarshal(bytes, &movies); err != nil {
		return nil, err
	}

	return &movies, nil
}

func (m *Movie) GetTitle() error {
	url := fmt.Sprintf("%s/%s/API/Title/%s/%s", os.Getenv("IMDB_API"), os.Getenv("IMDB_API_LANG"), os.Getenv("IMDB_API_KEY"), m.Id)
	headers := make(map[string]string)

	_, bytes, err := util.DoHttp(url, nil, "GET", headers)
	if err != nil {
		return err
	}
	n := Movie{}
	if err := json.Unmarshal(bytes, &n); err != nil {
		return err
	}
	log.Printf("%v", n)
	println(fmt.Sprintf("%v", n))
	return nil
}
