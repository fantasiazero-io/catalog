package commands

type CreateCharacter struct {
	MangaId       string
	Localizations []CharacterLocalization
	ClassId       int
	PersonalityId string
	LeadershipId  string
	RarityId      int
	TeamImage     string
	CardImage     string
}

type CharacterLocalization struct {
	Name        string
	Description Description
	Language    string
}

type Description struct {
	Short string
	Full  string
}
