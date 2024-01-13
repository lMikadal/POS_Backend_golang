package repository

import (
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
)

type priceRepository struct {
	db *DB
}

func NewPriceRepository(db *DB) *priceRepository {
	return &priceRepository{db}
}

func (r priceRepository) Create(price *domain.Price) (*domain.Price, error) {
	err := r.db.Create(&price).Error
	if err != nil {
		return nil, err
	}

	return price, nil
}

func (r priceRepository) Update(price *domain.Price, id int) (*domain.Price, error) {
	var newPrice domain.Price

	err := r.db.First(&newPrice, id).Error
	if err != nil {
		return nil, err
	}

	newPrice.PriceAmount = price.PriceAmount
	newPrice.PricePrice = price.PricePrice

	err = r.db.Save(&newPrice).Error
	if err != nil {
		return nil, err
	}

	return &newPrice, nil
}

func (r priceRepository) Delete(id int) error {
	err := r.db.Delete(&domain.Price{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
