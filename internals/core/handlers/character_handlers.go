package handlers

import (
	"catalog/internals/domain/aggregates"
	commands "catalog/internals/domain/commands"
	handlers "catalog/internals/domain/handlers"
	"catalog/internals/domain/valueObjects"
	persistence "catalog/internals/infrastructure/persistence/mongo"

	"github.com/google/uuid"
)

func CreateCharacter(command commands.CreateCharacter) handlers.HandlerResponse {
	var id = uuid.New().String()
	var localizations = make([]aggregates.CharacterLocalization, len(command.Localizations))
	for index, element := range command.Localizations {
		localizations[index] = aggregates.CharacterLocalization{
			Name:        element.Name,
			Language:    element.Language,
			Description: valueObjects.NewDescription(element.Description.Short, element.Description.Full),
		}
	}
	var error = persistence.CreateCharacter(id, command.MangaId,
		localizations, aggregates.Class(command.ClassId),
		command.PersonalityId, command.LeadershipId,
		valueObjects.Rarity(command.RarityId), 1, true)
	if error != nil {
		return handlers.HandlerResponse{
			HasError: true,
			Message:  "Character creation failed",
		}
	}
	return handlers.HandlerResponse{
		HasError: false,
	}

}
