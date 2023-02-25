package readModels

import (
	"time"
)

type SearchMangaReadModel struct {
	Items []SearchMangaReadModelDetail
}

type SearchMangaReadModelDetail struct {
	Id           string
	Name         string
	CreationDate time.Time
	LastUpdate   *time.Time
}

type GetMangaByIdReadModel struct {
	Id           string
	Name         string
	CreationDate time.Time
	LastUpdate   *time.Time
}
