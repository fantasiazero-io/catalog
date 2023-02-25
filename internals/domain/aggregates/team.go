package aggregates

import "errors"

type Team struct {
	Id         string
	CaptainId  string
	Characters []string
	Boosters   []string
}

func NewTeam(id string, captainId string, characters []string, boosters []string) (*Team, error) {
	if len(characters) > 7 || len(boosters) > 3 {
		return nil, errors.New("characters cannot be more than 7 and boosters cannot be more than 3")
	}
	return &Team{
		Id:         id,
		CaptainId:  captainId,
		Characters: characters,
		Boosters:   boosters,
	}, nil

}
