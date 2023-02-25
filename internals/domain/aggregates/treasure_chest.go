package aggregates

import (
	"catalog/internals/domain/valueObjects"
)

type TreasureChest struct {
	Id            string
	MangaId       string
	Localizations []BoosterLocalization
	Price         valueObjects.Price
	Rarity        string
}

type TreasureChestLocalization struct {
	Name        string
	Description valueObjects.Description
	Language    string
}

func NewTreasureChest(id string, mangaId string, localizations []BoosterLocalization, price valueObjects.Price,
	rarity string) TreasureChest {
	return TreasureChest{
		Id:            id,
		MangaId:       mangaId,
		Localizations: localizations,
		Price:         price,
		Rarity:        rarity,
	}
}
