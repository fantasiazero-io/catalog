package aggregates

type Character struct {
	Id            string
	MangaId       string
	Name          string
	ClassId       int
	PersonalityId int
	LeadershipId  int
	RarityId      int
}

func NewCharacter(id string, mangaId string, name string,
	classId int, personalityId int, leadershipId int,
	rarityId int) Character {
	return Character{
		Id:            id,
		MangaId:       mangaId,
		Name:          name,
		ClassId:       classId,
		PersonalityId: personalityId,
		LeadershipId:  leadershipId,
		RarityId:      rarityId,
	}
}
