package business

import (
	"github.com/rithikjain/local-businesses-backend/pkg"
	"github.com/rithikjain/local-businesses-backend/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	AddBusiness(business *models.Business) error

	GetApprovedBusinesses(page, pageSize int) (*[]models.Business, error)

	GetBusinessesByCity(city string, page, pageSize int) (*[]models.Business, error)

	GetBusinessesByCityAndType(city, typ string, page, pageSize int) (*[]models.Business, error)
}

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) AddBusiness(business *models.Business) error {
	err := r.DB.Create(business).Error
	if err != nil {
		return pkg.ErrDatabase
	}
	return nil
}

func (r *repo) GetApprovedBusinesses(page, pageSize int) (*[]models.Business, error) {
	var bizs []models.Business

	err := r.DB.Where("approved=?", true).Scopes(pkg.Paginate(page, pageSize)).Find(&bizs).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}

	return &bizs, nil
}

func (r *repo) GetBusinessesByCity(city string, page, pageSize int) (*[]models.Business, error) {
	var bizs []models.Business

	err := r.DB.Where("approved=? and location_city=?", true, city).
		Scopes(pkg.Paginate(page, pageSize)).
		Find(&bizs).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}

	return &bizs, nil
}

func (r *repo) GetBusinessesByCityAndType(city, typ string, page, pageSize int) (*[]models.Business, error) {
	var bizs []models.Business

	err := r.DB.Where("approved=? and location_city=? and type=?", true, city, typ).
		Scopes(pkg.Paginate(page, pageSize)).
		Find(&bizs).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}

	return &bizs, nil
}
