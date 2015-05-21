package poems

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func GetPoemsList() PoemsList {
	poems := getPoemsListFromJsonMockFile()
	return poems
}

func GetOnePoemBySlug(slug string) (Poem, error) {
	poems := getPoemsListFromJsonMockFile()
	for _, poem := range poems {
		if slug == poem.Slug {
			return poem, nil
		}
	}
	return Poem{}, errors.New("No poem found")
}

func getPoemsListFromJsonMockFile() PoemsList {
	jsonBlob, err := ioutil.ReadFile("./poems_mocks.json")
	if err != nil {
		fmt.Println("error:", err)
	}
	var poems PoemsList
	err = json.Unmarshal(jsonBlob, &poems)
	if err != nil {
		fmt.Println("error:", err)
	}
	return poems
}
