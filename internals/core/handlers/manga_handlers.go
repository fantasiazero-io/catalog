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

func DeleteManga(id string) handlers.HandlerResponse {
	var error = persistence.DeleteManga(id)
	if error != nil {
		return handlers.HandlerResponse{
			HasError: true,
			Message:  "Manga deletion failed",
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

func GetMangaById(id string) readModels.GetMangaByIdReadModel {
	var manga = persistence.GetMangaByIdManga(id)
	return manga
}
