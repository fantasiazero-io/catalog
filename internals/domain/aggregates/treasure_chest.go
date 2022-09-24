package aggregates

import (
	"catalog/internals/domain/valueObjects"
)

type TreasureChest struct {
	Id            string
	MangaId       string
	Localizations []BoosterLocalization
	Price         valueObjects.Price
	Rarity        valueObjects.Rarity
}

type TreasureChestLocalization struct {
	Name        string
	Description valueObjects.Description
	Language    string
}

func NewTreasureChest(id string, mangaId string, localizations []BoosterLocalization, price valueObjects.Price,
	rarity valueObjects.Rarity) TreasureChest {
	return TreasureChest{
		Id:            id,
		MangaId:       mangaId,
		Localizations: localizations,
		Price:         price,
		Rarity:        rarity,
	}
}
