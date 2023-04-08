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

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if (err != nil) {
		return nil, err	
	}
	return &pb.Category {
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesResponse []*pb.Category

	for _, category := range categories {
		categoryResponse := &pb.Category {
			Id: category.ID,
			Name: category.Name,
			Description: category.Description,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryList { Categories: categoriesResponse }, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category {
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}, nil
}
