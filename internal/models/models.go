package models

type Song struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
	Author string `json:"author"`
}
