package entities

type Rarity struct {
	Id            string
	Localizations []RarityLocalization
	Color         string
}

type RarityLocalization struct {
	Name     string
	Language string
}
