package aggregates

import (
	"catalog/internals/domain/valueObjects"
)

type Equipment struct {
	Id            string
	MangaId       string
	Localizations []BoosterLocalization
}

type EquipmentLocalization struct {
	Name        string
	Description valueObjects.Description
	Language    string
}

func NewEquipment(id string, mangaId string, localizations []BoosterLocalization) Equipment {
	return Equipment{
		Id:            id,
		MangaId:       mangaId,
		Localizations: localizations,
	}
}
