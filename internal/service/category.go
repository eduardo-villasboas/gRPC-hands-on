package service

import (
	"context"
	"github.com/eduardo-villasboas/gRPC-hands-on/internal/pb"
	"github.com/eduardo-villasboas/gRPC-hands-on/internal/database"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService {
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if (err != nil) {
		return nil, err	
	}
	categoryR := &pb.Category {
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}
	return &pb.CategoryResponse {
		Category: categoryR,
	}, nil
}
