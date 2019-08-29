package application

import (
	"time"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"github.com/sekky0905/random_lunch/src/infra/repository"
	"github.com/sekky0905/random_lunch/src/models"
)

// ShopApplicationService は、Shop のアプリケーションサービス。
type ShopApplicationService struct {
}

// ChooseShops は、お店を選択する。
func (s ShopApplicationService) ChooseShops(slackReq *models.SlackReq) (*models.SlackParams, error) {
	rep := repository.NewShopRepository()
	shops, err := rep.GetShops()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shops")
	}

	chosen := models.ChoiceShopsRandomly(shops)

	text := models.BuildChoiceText(chosen)
	return &models.SlackParams{
		Channel:  slackReq.ChannelID,
		Username: models.UserName,
		Text:     text,
	}, nil
}

// CreateShop は、お店を作成する。
func (s ShopApplicationService) CreateShop(slackReq *models.SlackReq, shop *models.Shop) (*models.SlackParams, error) {
	rep := repository.NewShopRepository()
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate new random")
	}

	shop.ID = u.String()
	shop.CreatedAt = time.Now()

	if err := rep.CreateShop(shop); err != nil {
		return nil, errors.Wrap(err, "failed to create shop")
	}

	text := models.BuildCreateText(shop)
	return &models.SlackParams{
		Channel:  slackReq.ChannelID,
		Username: models.UserName,
		Text:     text,
	}, nil
}

// ListShops は、お店の一覧を取得する。
func (s ShopApplicationService) ListShops(slackReq *models.SlackReq) (*models.SlackParams, error) {
	rep := repository.NewShopRepository()
	shops, err := rep.GetShops()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shops")
	}

	text := models.BuildListText(shops)
	return &models.SlackParams{
		Channel:  slackReq.ChannelID,
		Username: models.UserName,
		Text:     text,
	}, nil
}
