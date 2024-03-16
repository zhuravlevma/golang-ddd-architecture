package dal

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/board/offer/domain/entities"
)

func OfferToOrm(offer *entities.OfferEntity) *orm.OfferOrm {
	return &orm.OfferOrm{
		ID:                     offer.Id,
		Name:                   offer.Name,
		OrderId:                offer.OrderId,
		CurierId:               offer.CurierId,
		VehicleType:            offer.VehicleType,
		PreferredDeliveryAreas: offer.PreferredDeliveryAreas,
		WorkingHours:           offer.WorkingHours,
		Weight:                 offer.Weight,
		Bid:                    offer.Bid,
	}
}

func OfferToDomain(ormOffer *orm.OfferOrm) *entities.OfferEntity {

	return &entities.OfferEntity{
		Id:                     ormOffer.ID,
		Name:                   ormOffer.Name,
		OrderId:                ormOffer.OrderId,
		CurierId:               ormOffer.CurierId,
		VehicleType:            ormOffer.VehicleType,
		PreferredDeliveryAreas: ormOffer.PreferredDeliveryAreas,
		WorkingHours:           ormOffer.WorkingHours,
		Weight:                 ormOffer.Weight,
		Bid:                    ormOffer.Bid,
	}
}
