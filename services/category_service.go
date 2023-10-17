package services

import (
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/internal/validationutil"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
)

type CategoryService interface {
	Create(request models.CreateCategoryRequest) (models.CreateCategoryResponse, error)
	GetAll() ([]models.GetCategoryResponse, error)
	GetByID(ctgID string) (models.GetCategoryResponse, error)
	GetByName(name string) ([]models.GetCategoryResponse, error)
	Update(request models.UpdateCategoryRequest) error
	Delete(ctgID string) error
}

type categoryServiceImpl struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryServiceImpl{
		categoryRepository: categoryRepository,
	}
}

func (service *categoryServiceImpl) Create(request models.CreateCategoryRequest) (models.CreateCategoryResponse, error) {
	err := validationutil.ValidateRequest(request)
	if err != nil {
		return models.CreateCategoryResponse{}, err
	}

	createdAt := time.Now()

	request.CreatedAt = createdAt
	request.UpdatedAt = request.CreatedAt

	category := entity.Category{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
	}

	_, err = service.categoryRepository.Insert(&category)
	if err != nil {
		return models.CreateCategoryResponse{}, err
	}

	response := models.CreateCategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}

	return response, nil
}

func (service *categoryServiceImpl) GetAll() ([]models.GetCategoryResponse, error) {
	categories, err := service.categoryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetCategoryResponse, len(categories))

	for key, category := range categories {
		responses[key] = models.GetCategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *categoryServiceImpl) GetByID(ctgID string) (models.GetCategoryResponse, error) {
	category, err := service.categoryRepository.FindByID(ctgID)
	if err != nil {
		return models.GetCategoryResponse{}, err
	}

	response := models.GetCategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}

	return response, nil
}

func (service *categoryServiceImpl) GetByName(name string) ([]models.GetCategoryResponse, error) {
	categories, err := service.categoryRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetCategoryResponse, len(categories))

	for key, category := range categories {
		responses[key] = models.GetCategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *categoryServiceImpl) Update(request models.UpdateCategoryRequest) error {
	err := validationutil.ValidateRequest(request)
	if err != nil {
		return err
	}

	category, err := service.categoryRepository.FindByID(request.ID)
	if err != nil {
		return err
	}

	request.UpdatedAt = time.Now()

	updatedCategory := entity.Category{
		ID:          category.ID,
		Name:        request.Name,
		Description: request.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
	}

	err = service.categoryRepository.Update(&updatedCategory)
	if err != nil {
		return err
	}

	return nil
}

func (service *categoryServiceImpl) Delete(ctgID string) error {
	isCategory, err := service.categoryRepository.FindByID(ctgID)
	if err != nil {
		return err
	}

	err = service.categoryRepository.Delete(isCategory.ID)
	if err != nil {
		return err
	}

	return nil
}
