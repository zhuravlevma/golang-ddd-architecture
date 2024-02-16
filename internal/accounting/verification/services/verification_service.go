package services

import (
	"errors"

	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/reports/report/domain/ports/in"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/verification/api"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/verification/dtos"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/accounting/verification/models"
	"gorm.io/gorm"
)

type VerificationService struct {
	Repository *gorm.DB
	ExternalVerificationApi api.ExternalVerificationApi
	FindReportWithPositionsByIdInPort in.FindReportWithPositionsByIdInPort
}

func (v *VerificationService) Update(updateVerificationDto *dtos.UpdateVerificationDto) (*models.Verfication, error) {
	var verification *models.Verfication
	v.Repository.Find(&verification, updateVerificationDto.Id)

	if (verification == nil) {
		return nil, errors.New("Verification not found")
	}

	report, err := v.FindReportWithPositionsByIdInPort.Execute(&in.FindReportWithPositionsParams{
		Id: verification.ReportId,
	})

	if err != nil {
		return nil, err
	}

	if updateVerificationDto.IsFull != nil {
		v.ExternalVerificationApi.FullVerifyReport(&api.ReportApiDto{
			Id: report.ID.String(),
			ReportNumber: report.ReportNumber,
		})
	}

	if updateVerificationDto.Signed != nil {
		verification.SignReport()
		v.ExternalVerificationApi.SignReport(report.ReportNumber)
	}

	if updateVerificationDto.Completed != nil {
		verification.CompleteVerification()
		v.ExternalVerificationApi.Complete(report.ReportNumber)
	}

	updatedErr := v.Repository.Model(&models.Verfication{}).Where("id = ?", verification.ID).Updates(verification).Error
	if updatedErr != nil {
		return nil, updatedErr
	}
	return verification, nil
}

func (v *VerificationService) Create(createVerificationDto *dtos.CreateVerificationDto) (*models.Verfication, error) {
	var verification *models.Verfication
	if err := v.Repository.Create(models.Verfication{
		IsFull: createVerificationDto.IsFull,
	}).Error; err != nil {
		return nil, err
	}
	return verification, nil
}
