package aggregates

import (
	"catalog/internals/domain/valueObjects"
)

type Character struct {
	Id            string
	MangaId       string
	Localizations []CharacterLocalization
	Class         Class
	PersonalityId string
	LeadershipId  string
	Rarity        valueObjects.Rarity
	Level         int
	IsPackable    bool
}

type CharacterLocalization struct {
	Name        string
	Description valueObjects.Description
	Language    string
}

func NewCharacter(id string, mangaId string, localizations []CharacterLocalization,
	class Class, personalityId string, leadershipId string,
	rarity valueObjects.Rarity, level int, isPackable bool) Character {
	return Character{
		Id:            id,
		MangaId:       mangaId,
		Localizations: localizations,
		Class:         class,
		PersonalityId: personalityId,
		LeadershipId:  leadershipId,
		Rarity:        rarity,
		Level:         level,
		IsPackable:    isPackable,
	}
}

type Class int

const (
	Fighter Class = iota + 1
	Leader
	Sapient
	Slasher
	Sniper
	Striker
	Support
	Tactician
)
