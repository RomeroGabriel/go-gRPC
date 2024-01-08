package service

import (
	"context"

	"github.com/RomeroGabriel/go-gRPC/internal/db"
	"github.com/RomeroGabriel/go-gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CatetgoryDb db.CategoryDB
}

func NewCategoryService(categoryDb db.CategoryDB) *CategoryService {
	return &CategoryService{
		CatetgoryDb: categoryDb,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CatetgoryDb.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, err
}
