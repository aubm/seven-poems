package poems

import (
	"encoding/json"
	"log"
)

type Poem struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Author  string `json:"author"`
	Year    string `json:"year"`
	Content string `json:"content"`
}

func (poem Poem) ToJson() string {
	bytes, err := json.Marshal(poem)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes[:])
}
