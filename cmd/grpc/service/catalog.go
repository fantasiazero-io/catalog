package service

import (
	"catalog/proto"
	"context"
)

type CatalogServer struct {
	proto.UnimplementedCatalogServer
}

func (s *CatalogServer) GetProducts(ctx context.Context, in *proto.GetProductsRequest) (*proto.GetProductsResponse, error) {
	return &proto.GetProductsResponse{
		Products: []*proto.GetProductsResponse_Product{
			{
				Title: "One Piece",
				Image: "onepiece.jpg",
			},
			{
				Title: "Dragon ball",
				Image: "dragonball.jpg",
			},
			{
				Title: "Naruto",
				Image: "naruto.jpg",
			},
		},
	}, nil
}
