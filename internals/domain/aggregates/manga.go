package aggregates

import "time"

type Manga struct {
	Id           string
	Name         string
	CreationDate time.Time
	LastUpdate   *time.Time
}

func NewManga(id string, name string) Manga {
	return Manga{
		Id:           id,
		Name:         name,
		CreationDate: time.Now().UTC(),
	}
}

func (manga *Manga) Update(name string) (message string, hasError bool) {
	if name == "" {
		return "Name cannot be empty", true
	}
	manga.Name = name
	manga.LastUpdate = setTime(time.Now().UTC())
	return "", false
}

func setTime(t time.Time) *time.Time {
	return &t
}
