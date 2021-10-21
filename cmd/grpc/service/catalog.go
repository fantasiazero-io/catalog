package service

import (
	"catalog/proto"
	"catalog/store"
	"context"

	"github.com/google/uuid"
)

type CatalogServer struct {
	proto.UnimplementedCatalogServer
}

func (s *CatalogServer) GetProducts(ctx context.Context, in *proto.GetProductsRequest) (*proto.GetProductsResponse, error) {
	products := store.GetProducts()
	var result []*proto.GetProductsResponse_Product
	for i := 0; i < len(products); i++ {
		result[i] = &proto.GetProductsResponse_Product{
			Title: products[i].Title,
			Image: products[i].Image,
			Id:    products[i].Id,
		}
	}
	return &proto.GetProductsResponse{
		Products: result,
	}, nil
}

func (s *CatalogServer) CreateProduct(ctx context.Context, in *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	var id = uuid.NewString()
	store.CreateProduct(id, in.Title, in.Image, int(in.Order))
	return &proto.CreateProductResponse{}, nil
}
