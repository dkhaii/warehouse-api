package services

// import (
// 	"context"

// 	"github.com/dkhaii/warehouse-api/models"
// )

// type userAdminServiceImpl struct {
// 	categoryService      CategoryService
// 	locationService      LocationService
// 	itemService          ItemService
// 	orderService         OrderService
// 	transferOrderService TransferOrderService
// }

// func NewUserAdminService(categoryService CategoryService, locationService LocationService, itemService ItemService, orderService OrderService, transferOrderService TransferOrderService) {
// 	return &userAdminServiceImpl{
// 		categoryService:      categoryService,
// 		locationService:      locationService,
// 		itemService:          itemService,
// 		orderService:         orderService,
// 		transferOrderService: transferOrderService,
// 	}
// }

// func (service *userAdminServiceImpl) CreateCategory(ctx context.Context, request models.CreateCategoryRequest) (models.CreateCategoryResponse, error) {
// 	category, err := service.categoryService.Create(ctx, request)
// 	if err != nil {
// 		return models.CreateCategoryResponse{}, err
// 	}

// 	return category, nil
// }

// func (service *userAdminServiceImpl) GetAllCategory(ctx context.Context) ([]models.GetCategoryResponse, error) {
// 	rows, err := service.categoryService.GetAll(ctx)
// 	if err !
// }
