package api

import "github.com/google/uuid"

type CheckCurrentPositionApiResponseDto struct {
	StartLatitude  *string
	StartLongitude *string
	EndLatitude    *string
	EndLongitude   *string
}

type CheckCurrentPositionApiDto struct {
	EntityId uuid.UUID
}

type ExternalTrackingApi struct{}

func (api *ExternalTrackingApi) CheckLastCoordinates(data *CheckCurrentPositionApiDto) CheckCurrentPositionApiResponseDto {
	return CheckCurrentPositionApiResponseDto{
		StartLatitude:  nil,
		StartLongitude: nil,
		EndLatitude:    nil,
		EndLongitude:   nil,
	}
}
