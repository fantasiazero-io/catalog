package aggregates

import (
	"catalog/internals/domain/valueObjects"
)

type Booster struct {
	Id            string
	MangaId       string
	Localizations []BoosterLocalization
	Price         valueObjects.Price
	Rarity        string
}

type BoosterLocalization struct {
	Name        string
	Description valueObjects.Description
	Language    string
}

func NewBooster(id string, mangaId string, localizations []BoosterLocalization, price valueObjects.Price,
	rarity string) Booster {
	return Booster{
		Id:            id,
		MangaId:       mangaId,
		Localizations: localizations,
		Price:         price,
		Rarity:        rarity,
	}
}
