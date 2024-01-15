package controller

import (
	"context"

	"github.com/asaba/models"
	"gorm.io/gorm"
)

type ProductLogController interface {
	CreateProductLog(ctx context.Context, req models.ProductLog) error
}
type productLogController struct {
	db *gorm.DB
}

func NewProductLogController(db *gorm.DB) ProductLogController {
	return &productLogController{db: db}
}

func (pl *productLogController) CreateProductLog(ctx context.Context, req models.ProductLog) error {
	return pl.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&models.ProductLog{
			Code:       req.Code,
			ActionType: req.ActionType,
			Quantity:   req.Quantity,
		}).Error

		if err != nil {
			return err
		}
		return nil
	})
}
