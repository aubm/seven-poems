package poems

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetPoemsList() PoemsList {
	poems := getPoemStructsFromJsonMockFile()
	return poems
}

func getOnePoemById() {
}

func getPoemStructsFromJsonMockFile() PoemsList {
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
