package poems

type Poem struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Year    string `json:"year"`
	Content string `json:"content"`
}
