package postgres

import (
	"catalog/internals/domain/aggregates"
	sq "github.com/Masterminds/squirrel"
	"log"
)

func CreateCharacter(id string, mangaId string, name string,
	classId int, personalityId int, leadershipId int,
	rarityId int) error {
	client, _ := CreateConnection()
	character := aggregates.NewCharacter(id, mangaId, name, classId, personalityId, leadershipId, rarityId)

	sql, args, err := sq.
		Insert("Characters").Columns("Id", "MangaId", "Name", "ClassId", "PersonalityId", "LeadershipId", "RarityId").
		Values(character.Id, character.MangaId, character.Name,
			character.ClassId, character.PersonalityId, character.LeadershipId,
			character.RarityId).
		ToSql()
	_, error := client.Exec(sql, args)

	if error != nil {
		log.Fatal(err)
	}
	return err
}
