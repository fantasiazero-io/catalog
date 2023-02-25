package handlers

import (
	"catalog/internals/domain/commands"
	"catalog/internals/domain/handlers"
	"catalog/internals/domain/queries"
	"catalog/internals/domain/readModels"
	persistence "catalog/internals/infrastructure/persistence/postgres"

	"github.com/google/uuid"
)

func CreateManga(command commands.CreateManga) handlers.HandlerResponse {
	var id = uuid.New().String()
	var error = persistence.CreateManga(id, command.Name)
	if error != nil {
		return handlers.HandlerResponse{
			HasError: true,
			Message:  "Manga creation failed",
		}
	}
	return handlers.HandlerResponse{
		HasError: false,
	}
}

func UpdateManga(command commands.UpdateManga) handlers.HandlerResponse {
	var error = persistence.UpdateManga(command.Id, command.Name)
	if error != nil {
		return handlers.HandlerResponse{
			HasError: true,
			Message:  "Manga update failed",
		}
	}
	return handlers.HandlerResponse{
		HasError: false,
	}
}

func SearchManga(query queries.SearchMangaQuery) readModels.SearchMangaReadModel {
	var readModel = persistence.SearchManga(query.Filter)
	return readModel
}
