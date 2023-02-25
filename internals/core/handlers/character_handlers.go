package handlers

import (
	"catalog/internals/domain/commands"
	"catalog/internals/domain/handlers"
	persistence "catalog/internals/infrastructure/persistence/postgres"

	"github.com/google/uuid"
)

func CreateCharacter(command commands.CreateCharacter) handlers.HandlerResponse {
	var id = uuid.New().String()
	var error = persistence.CreateCharacter(id, command.MangaId,
		command.Name, command.ClassId,
		command.PersonalityId, command.LeadershipId,
		command.RarityId)
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
