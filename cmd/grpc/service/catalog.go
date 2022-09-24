package service

import (
	handlers "catalog/internals/core/handlers"
	"catalog/internals/domain/commands"
	"catalog/proto"
	"context"
)

type CatalogServer struct {
	proto.UnimplementedCatalogServer
}

// func (s *CatalogServer) GetCharacters(ctx context.Context, in *proto.GetCharactersRequest) (*proto.GetCharactersResponse, error) {
// 	characters := handlers.GetCharacters()

// 	var result = make([]*proto.GetCharactersResponse, len(characters))
// 	for i := 0; i < len(characters); i++ {
// 		result[i] = &proto.GetCharactersResponse_Character{
// 			Id: characters[i].Id,
// 		}
// 	}
// 	return &proto.GetCharactersResponse{
// 		Characters: result,
// 	}, nil
// }

func (s *CatalogServer) CreateProduct(ctx context.Context, in *proto.CreateCharacterRequest) (*proto.MutationResponse, error) {
	var localizations = make([]commands.CharacterLocalization, len(in.Localizations))
	for index, element := range in.Localizations {
		localizations[index] = commands.CharacterLocalization{
			Name: element.Name,
			Description: commands.Description{
				Short: element.Description.Short,
				Full:  element.Description.Full,
			},
			Language: element.Language,
		}
	}
	handlers.CreateCharacter(commands.CreateCharacter{
		MangaId:       in.MangaId,
		Localizations: localizations,
		ClassId:       int(in.ClassId),
		PersonalityId: in.PersonalityId,
		LeadershipId:  in.LeadershipId,
		RarityId:      int(in.RarityId),
		TeamImage:     in.TeamImage,
		CardImage:     in.CardImage,
	})
	return &proto.MutationResponse{}, nil
}
