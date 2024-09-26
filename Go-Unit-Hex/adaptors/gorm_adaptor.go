package adaptors

import (
	"github.com/thitipa-palm/go-Unit-Hex/core"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) core.OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) Save(order core.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		return result.Error
	}
	return nil
}
