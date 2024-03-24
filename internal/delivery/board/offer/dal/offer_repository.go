package dal

import (
	"github.com/google/uuid"
	relay "github.com/zhuravlevma/golang-ddd-architecture/internal/__relay__"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/entities"
	"gorm.io/gorm"
)

type OfferRepository struct {
	Db *gorm.DB
}

func (repo *OfferRepository) FindOfferById(id uuid.UUID) (*entities.OfferEntity, error) {
	var dbOffer orm.OfferOrm
	if err := repo.Db.First(&dbOffer, id).Error; err != nil {
		return nil, err
	}

	return OfferToDomain(&dbOffer), nil
}

func (repo *OfferRepository) FindOfferByOrderId(orderId uuid.UUID) (*entities.OfferEntity, error) {
	var dbOffer orm.OfferOrm
	if err := repo.Db.First(&dbOffer, repo.Db.Where("orderId = ?", orderId)).Error; err != nil {
		return nil, err
	}

	return OfferToDomain(&dbOffer), nil
}

func (repo *OfferRepository) UpdateOffer(offer *entities.OfferEntity) error {
	dbOffer := OfferToOrm(offer)
	var messages []*relay.MessageOrm
	for _, message := range offer.DomainMessages {
		messages = append(messages, relay.DomainToORM(message))
	}

	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&messages)

		err := tx.Model(&dbOffer).Where("id = ?", dbOffer.ID).Updates(dbOffer).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil
	}

	storedOffer, err := repo.FindOfferById(offer.Id)
	if err != nil {
		return err
	}
	*offer = *storedOffer
	return nil
}

func (repo *OfferRepository) CreateOffer(offer *entities.OfferEntity) error {
	dbOffer := OfferToOrm(offer)
	var messages []*relay.MessageOrm
	for _, message := range offer.DomainMessages {
		messages = append(messages, relay.DomainToORM(message))
	}

	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&messages)

		if err := repo.Db.Create(dbOffer).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil
	}

	storedOffer, err := repo.FindOfferById(dbOffer.ID)
	if err != nil {
		return err
	}
	*offer = *storedOffer
	return nil
}
