package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/radiance822/time_memorizer/internal/app/provider"
	pb "github.com/radiance822/time_memorizer/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Implementation struct {
	pb.UnimplementedTimeMemorizerServer
	categoryProvider *provider.CategoryProvider
	timeProvider     *provider.TimeProvider
	userProvider     *provider.UserProvider
}

func NewImplementation(db *sqlx.DB) *Implementation {
	return &Implementation{
		categoryProvider: provider.NewCategoryProvider(db),
		timeProvider:     provider.NewTimeProvider(db),
		userProvider:     provider.NewUserProvider(db),
	}
}

func (imp *Implementation) AddTime(ctx context.Context, req *pb.AddTimeRequest) (*pb.AddTimeResponse, error) {
	// Получаем userID по имени пользователя
	userID, err := imp.userProvider.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get user: %v", err)
	}

	// Получаем или создаем категорию
	category, err := imp.categoryProvider.GetByName(ctx, req.Title)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get category: %v", err)
	}
	if category == nil {
		// Категория не существует, создаем новую
		categoryID, err := imp.categoryProvider.Store(ctx, req.Title)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to store category: %v", err)
		}
		category = &provider.Category{ID: categoryID, Name: req.Title}
	}

	// Добавляем время
	err = imp.timeProvider.Add(ctx, req.Amount, category.ID, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add time: %v", err)
	}

	return &pb.AddTimeResponse{Category: &pb.Category{Id: category.ID, Title: category.Name, Amount: req.Amount}}, nil
}

func (imp *Implementation) SubtractTime(ctx context.Context, req *pb.SubtractTimeRequest) (*pb.SubtractTimeResponse, error) {
	userID, err := imp.userProvider.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get user: %v", err)
	}

	category, err := imp.categoryProvider.GetByName(ctx, req.Title)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get category: %v", err)
	}
	if category == nil {
		return nil, status.Errorf(codes.NotFound, "Category not found")
	}

	existingTime, err := imp.timeProvider.GetByCategoryAndUser(ctx, category.ID, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get time: %v", err)
	}
	if existingTime == nil {
		return nil, status.Errorf(codes.NotFound, "Time record not found")
	}

	if existingTime.Amount < req.Amount {
		return nil, status.Errorf(codes.InvalidArgument, "Insufficient time to subtract")
	}

	err = imp.timeProvider.Subtract(ctx, req.Amount, category.ID, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to subtract time: %v", err)
	}

	return &pb.SubtractTimeResponse{Category: &pb.Category{Id: category.ID, Title: category.Name, Amount: req.Amount}}, nil
}

func (imp *Implementation) ShowTime(ctx context.Context, req *pb.ShowTimeRequest) (*pb.ShowTimeResponse, error) {
	userID, err := imp.userProvider.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get user: %v", err)
	}

	categories, err := imp.timeProvider.GetByUserId(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get time: %v", err)
	}

	var pbCategories []*pb.Category
	for _, category := range categories {
		amount, err := imp.timeProvider.GetByCategoryAndUser(ctx, category.ID, userID)
		if err != nil {
			return nil, err
		}
		pbCategories = append(pbCategories, &pb.Category{Id: category.ID, Title: category.Name, Amount: amount.Amount})
	}

	return &pb.ShowTimeResponse{Categories: pbCategories}, nil
}
