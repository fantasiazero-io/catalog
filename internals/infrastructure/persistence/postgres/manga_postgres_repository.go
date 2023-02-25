package postgres

import (
	"catalog/internals/domain/aggregates"
	"catalog/internals/domain/readModels"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func CreateManga(id string, name string) error {
	db, err := CreateConnection()
	manga := aggregates.NewManga(id, name)

	if err != nil {
		panic(err)
	}

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

	manga := aggregates.NewManga(id, name)

	if err != nil {
		panic(err)
	}

	sql := "UPDATE FROM \"Mangas\"  SET \"Name\" = $1, \"LastUpdate\"= $2 WHERE \"Id\" = $3"

	result, err := db.Exec(sql,  manga.Name, ,manga.Id,)

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
		panic(err)
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

func GetMangaByIdManga(id string)  *readModels.GetMangaByIdReadModel{
	manga:= getMangaByIdManga(id)
	return readModels.GetMangaByIdReadModel{

	}
}

func getMangaByIdManga(id string)  *GetMangaByIdQueryModel{
	db, err := CreateConnection()
	if err != nil {
		panic(err)
	}
	manga := GetMangaByIdQueryModel{}
	sql := "SELECT \"Id\",\"Name\", \"CreationDate\", \"LastUpdate\"  FROM \"Mangas\" WHERE \"Id\" = $1"
	err = db.Select(&manga, sql, id)
	CloseConnection(db)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &manga
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
