package postgres

import (
	"catalog/internals/domain/aggregates"
	"catalog/internals/domain/readModels"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func CreateManga(id string, name string) error {
	db, err := CreateConnection()
	manga := aggregates.NewManga(id, name)

	sql := "INSERT INTO \"Mangas\" (\"Id\",\"Name\", \"CreationDate\") VALUES ($1,$2,$3)"

	result, err := db.Exec(sql, manga.Id, manga.Name, manga.CreationDate)

	CloseConnection(db)
	if err != nil {
		log.Fatal(result)
		log.Fatal(err)
		return nil
	}

	return nil
}

func UpdateManga(id string, name string) error {
	db, err := CreateConnection()

	manga := getMangaByIdManga(id)

	if manga == nil {
		return errors.New("Manga with not found")
	}

	manga.Update(name)

	sql := "UPDATE \"Mangas\"  SET \"Name\" = $1, \"LastUpdate\"= $2 WHERE \"Id\" = $3"
	result, err := db.Exec(sql, manga.Name, manga.LastUpdate, manga.Id)

	CloseConnection(db)
	if err != nil {
		log.Fatal(result)
		log.Fatal(err)
		return nil
	}

	return nil
}

func DeleteManga(id string) error {
	db, err := CreateConnection()

	sql := "DELETE FROM \"Mangas\" WHERE \"Id\" = $1"
	result, err := db.Exec(sql, id)

	CloseConnection(db)
	if err != nil {
		log.Fatal(result)
		log.Fatal(err)
		return nil
	}

	return nil
}

func SearchManga(filter string) readModels.SearchMangaReadModel {
	db, err := CreateConnection()
	mangas := []SearchMangaQueryModel{}
	if err != nil {
		log.Fatal(err)
		return readModels.SearchMangaReadModel{
			Items: []readModels.SearchMangaReadModelDetail{},
		}
	}

	sql := ""
	if filter == "" {
		sql = "SELECT \"Id\",\"Name\", \"CreationDate\", \"LastUpdate\"  FROM \"Mangas\""
		err = db.Select(&mangas, sql)
	} else {
		sql = "SELECT \"Id\",\"Name\", \"CreationDate\" FROM \"Mangas\" WHERE \"Name\" ILIKE '%$1%'"
		err = db.Select(&mangas, sql, filter)
	}

	CloseConnection(db)
	if err != nil {
		log.Fatal(err)
		return readModels.SearchMangaReadModel{
			Items: []readModels.SearchMangaReadModelDetail{},
		}
	}
	var readModelDetails = make([]readModels.SearchMangaReadModelDetail, len(mangas))
	for i := 0; i < len(mangas); i++ {
		readModelDetails[i] = readModels.SearchMangaReadModelDetail{
			Id:           mangas[i].Id,
			Name:         mangas[i].Name,
			CreationDate: mangas[i].CreationDate,
			LastUpdate:   mangas[i].LastUpdate,
		}
	}
	return readModels.SearchMangaReadModel{
		Items: readModelDetails,
	}
}

func GetMangaByIdManga(id string) readModels.GetMangaByIdReadModel {
	var manga = getMangaByIdManga(id)
	return readModels.GetMangaByIdReadModel{
		Id:           manga.Id,
		Name:         manga.Name,
		CreationDate: manga.CreationDate,
		LastUpdate:   manga.LastUpdate,
	}
}

func getMangaByIdManga(id string) *aggregates.Manga {
	db, err := CreateConnection()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	manga := GetMangaByIdQueryModel{}
	sql := "SELECT \"Id\",\"Name\", \"CreationDate\", \"LastUpdate\"  FROM \"Mangas\" WHERE \"Id\" = $1 LIMIT 1"
	err = db.Get(&manga, sql, id)
	CloseConnection(db)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &aggregates.Manga{
		Id:           manga.Id,
		Name:         manga.Name,
		CreationDate: manga.CreationDate,
		LastUpdate:   manga.LastUpdate,
	}
}

type SearchMangaQueryModel struct {
	Id           string     `json:"Id" db:"Id"`
	Name         string     `json:"Name" db:"Name"`
	CreationDate time.Time  `json:"CreationDate" db:"CreationDate"`
	LastUpdate   *time.Time `json:"LastUpdate" db:"LastUpdate"`
}

type GetMangaByIdQueryModel struct {
	Id           string     `json:"Id" db:"Id"`
	Name         string     `json:"Name" db:"Name"`
	CreationDate time.Time  `json:"CreationDate" db:"CreationDate"`
	LastUpdate   *time.Time `json:"LastUpdate" db:"LastUpdate"`
}
