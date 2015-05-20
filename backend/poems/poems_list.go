package poems

import (
	"encoding/json"
	"log"
)

type PoemsList []Poem

func (poems PoemsList) ToJson() string {
	bytes, err := json.Marshal(poems)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes[:])
}
