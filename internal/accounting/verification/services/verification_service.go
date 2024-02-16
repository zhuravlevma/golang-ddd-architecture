package services

import (
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/verification/api"
	"gorm.io/gorm"
)

type VerificationService struct {
	Db *gorm.DB
	ExternalVerificationApi api.ExternalVerificationApi
	FindReportWithPositionsByIdInPort in.FindReportWithPositionsByIdInPort
}

func (v *VerificationService) create() {}
