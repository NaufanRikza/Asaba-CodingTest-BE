package controller

import (
	"context"

	"github.com/asaba/models"
	"gorm.io/gorm"
)

type ProductController interface {
	CreateProducts(ctx context.Context, req models.CreateProducts) error
	GetProducts(ctx context.Context) (models.Products, error)
	UpdateProducts(ctx context.Context, req models.UpdateProducts) error
}

type productController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) ProductController {
	return &productController{db: db}
}

func convertBoolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}

func (p *productController) CreateProducts(ctx context.Context, req models.CreateProducts) error {

	return p.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, r := range req {
			err := tx.Create(&models.Product{
				Code:        r.Code,
				Name:        r.Name,
				Quantity:    r.Quantity,
				Description: r.Description,
				IsActive:    convertBoolToInt(r.IsActive),
			}).Error

			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (p *productController) GetProducts(ctx context.Context) (models.Products, error) {
	var products models.Products
	err := p.db.WithContext(ctx).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productController) UpdateProducts(ctx context.Context, req models.UpdateProducts) error {
	return p.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, r := range req {
			err := tx.Where("code = ?", r.Code).Updates(&models.Product{
				Name:        r.Name,
				Quantity:    r.Quantity,
				Description: r.Description,
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}
