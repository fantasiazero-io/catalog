package api

import (
	"catalog/internals/core/handlers"
	"catalog/internals/domain/commands"
	"catalog/internals/domain/queries"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var request CreateCharacterRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	response := handlers.CreateCharacter(commands.CreateCharacter{
		MangaId:       mux.Vars(r)["id"],
		Name:          request.Name,
		ClassId:       request.ClassId,
		PersonalityId: request.PersonalityId,
		LeadershipId:  request.LeadershipId,
		RarityId:      request.RarityId,
	})
	if response.HasError {
		respondWithError(w, http.StatusInternalServerError, response.Message)
		return
	}
	respondWithJSON(w, http.StatusCreated, request)
}

func CreateManga(w http.ResponseWriter, r *http.Request) {
	var request CreateMangaRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	response := handlers.CreateManga(commands.CreateManga{
		Name: request.Name,
	})

	if response.HasError {
		respondWithError(w, http.StatusInternalServerError, response.Message)
		return
	}
	respondWithJSON(w, http.StatusOK, request)
}

func UpdateManga(w http.ResponseWriter, r *http.Request) {
	var request CreateMangaRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	response := handlers.UpdateManga(commands.UpdateManga{
		Id:   mux.Vars(r)["id"],
		Name: request.Name,
	})

	if response.HasError {
		respondWithError(w, http.StatusInternalServerError, response.Message)
		return
	}
	respondWithJSON(w, http.StatusOK, request)
}

func DeleteManga(w http.ResponseWriter, r *http.Request) {
	response := handlers.DeleteManga(mux.Vars(r)["id"])

	if response.HasError {
		respondWithError(w, http.StatusInternalServerError, response.Message)
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

func SearchManga(w http.ResponseWriter, r *http.Request) {

	response := handlers.SearchManga(queries.SearchMangaQuery{
		Filter: mux.Vars(r)["filter"],
	})
	respondWithJSON(w, http.StatusOK, response.Items)
}

func GetMangaById(w http.ResponseWriter, r *http.Request) {
	response := handlers.GetMangaById(mux.Vars(r)["id"])
	respondWithJSON(w, http.StatusOK, response)
}

type CreateCharacterRequest struct {
	Name          string
	ClassId       int
	LeadershipId  int
	PersonalityId int
	RarityId      int
}

type CreateMangaRequest struct {
	Name string
}

type UpdateMangaRequest struct {
	Name string
}
