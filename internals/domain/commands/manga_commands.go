package commands

type CreateManga struct {
	Name string
}

type UpdateManga struct {
	Id   string
	Name string
}

type DeleteManga struct {
	Id string
}
