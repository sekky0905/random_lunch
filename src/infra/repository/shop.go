package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
	"github.com/sekky0905/random_lunch/src/models"
)

// ShopRepository は、Shop の Repository。
type ShopRepository struct {
	Table dynamo.Table
}

// NewShopRepository は、ShopRepository を生成して返す。
func NewShopRepository() *ShopRepository {
	db := dynamo.New(session.New(), &aws.Config{})
	return &ShopRepository{
		Table: db.Table("Shops"),
	}
}

// GetShops は、Shop の取得する。
func (r *ShopRepository) GetShops() ([]models.Shop, error) {
	shops := make([]models.Shop, 0)
	if err := r.Table.Scan().All(&shops); err != nil {
		return nil, errors.Wrap(err, "failed to scan all shops from dynamo db")
	}

	return shops, nil
}

// CreateShop は、Shop を作成する。
func (r *ShopRepository) CreateShop(shop *models.Shop) error {
	if err := r.Table.Put(shop).Run(); err != nil {
		return errors.Wrap(err, "failed to put shop to dynamo db")
	}

	return nil
}
