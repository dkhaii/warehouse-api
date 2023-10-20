package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
)

type categoryServiceImpl struct {
	categoryRepository repositories.CategoryRepository
	database           *sql.DB
}

func NewCategoryService(categoryRepository repositories.CategoryRepository, database *sql.DB) CategoryService {
	return &categoryServiceImpl{
		categoryRepository: categoryRepository,
		database:           database,
	}
}

func (service *categoryServiceImpl) Create(ctx context.Context, request models.CreateCategoryRequest) (models.CreateCategoryResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateCategoryResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateCategoryResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	createdAt := time.Now()
	request.CreatedAt = createdAt
	request.UpdatedAt = request.CreatedAt

	category := entity.Category{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		Location:    nil,
	}

	createdCategory, err := service.categoryRepository.Insert(ctx, tx, &category)
	if err != nil {
		return models.CreateCategoryResponse{}, err
	}

	response := models.CreateCategoryResponse{
		ID:          createdCategory.ID,
		Name:        createdCategory.Name,
		Description: createdCategory.Description,
		LocationID:  createdCategory.LocationID,
		CreatedAt:   createdCategory.CreatedAt,
		UpdatedAt:   createdCategory.UpdatedAt,
	}

	return response, nil
}

func (service *categoryServiceImpl) GetAll(ctx context.Context) ([]models.GetCategoryResponse, error) {
	categories, err := service.categoryRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetCategoryResponse, len(categories))

	for key, category := range categories {
		responses[key] = models.GetCategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			LocationID:  category.LocationID,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *categoryServiceImpl) GetByID(ctx context.Context, ctgID string) (models.GetCategoryResponse, error) {
	category, err := service.categoryRepository.FindByID(ctx, ctgID)
	if err != nil {
		return models.GetCategoryResponse{}, err
	}

	response := models.GetCategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		LocationID:  category.LocationID,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}

	return response, nil
}

func (service *categoryServiceImpl) GetByName(ctx context.Context, name string) ([]models.GetCategoryResponse, error) {
	categories, err := service.categoryRepository.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetCategoryResponse, len(categories))

	for key, category := range categories {
		responses[key] = models.GetCategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			LocationID:  category.LocationID,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *categoryServiceImpl) Update(ctx context.Context, request models.UpdateCategoryRequest) (models.CreateCategoryResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateCategoryResponse{}, err
	}

	category, err := service.categoryRepository.FindByID(ctx, request.ID)
	if err != nil {
		return models.CreateCategoryResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateCategoryResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	request.UpdatedAt = time.Now()

	updatedCategory := entity.Category{
		ID:          category.ID,
		Name:        request.Name,
		Description: request.Description,
		LocationID:  request.LocationID,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
	}

	categoryData, err := service.categoryRepository.Update(ctx, tx, &updatedCategory)
	if err != nil {
		return models.CreateCategoryResponse{}, err
	}

	response := models.CreateCategoryResponse{
		ID:          categoryData.ID,
		Name:        categoryData.Name,
		Description: categoryData.Description,
		LocationID:  categoryData.LocationID,
		CreatedAt:   categoryData.CreatedAt,
		UpdatedAt:   categoryData.UpdatedAt,
	}

	return response, nil
}

func (service *categoryServiceImpl) Delete(ctx context.Context, ctgID string) error {
	category, err := service.categoryRepository.FindByID(ctx, ctgID)
	if err != nil {
		return err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return err
	}
	defer helpers.CommitOrRollBack(tx)

	err = service.categoryRepository.Delete(ctx, tx, category.ID)
	if err != nil {
		return err
	}

	return nil
}
