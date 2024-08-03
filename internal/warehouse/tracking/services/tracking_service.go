package services

import (
	"errors"

	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/tracking/api"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/tracking/dtos"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/delivery/tracking/models"
	"gorm.io/gorm"
)

type TrackingService struct {
	Repository          *gorm.DB
	ExternalTrackingApi api.ExternalTrackingApi
}

func (v *TrackingService) Update(updateRouteDto *dtos.UpdateRouteDto) (*models.Route, error) {
	var route *models.Route
	v.Repository.Find(&route, updateRouteDto.Id)

	if route == nil {
		return nil, errors.New("Route not found")
	}

	if updateRouteDto.CourierId != nil {
		route.CourierId = *updateRouteDto.CourierId
	}

	if updateRouteDto.OrderId != nil {
		route.OrderId = *updateRouteDto.OrderId
	}

	newCoordinates := v.ExternalTrackingApi.CheckLastCoordinates(&api.CheckCurrentPositionApiDto{EntityId: route.CourierId})

	route.StartLatitude = newCoordinates.StartLatitude
	route.EndLatitude = newCoordinates.EndLatitude
	route.StartLongitude = newCoordinates.StartLongitude
	route.EndLongitude = newCoordinates.EndLongitude

	updatedErr := v.Repository.Model(&models.Route{}).Where("id = ?", route.ID).Updates(route).Error
	if updatedErr != nil {
		return nil, updatedErr
	}
	return route, nil
}

func (v *TrackingService) Create(createRouteDto *dtos.CreateRouteDto) (*models.Route, error) {
	var route *models.Route
	if err := v.Repository.Create(models.Route{
		OrderId:   createRouteDto.OrderId,
		CourierId: createRouteDto.CourierId,
	}).Error; err != nil {
		return nil, err
	}
	return route, nil
}
