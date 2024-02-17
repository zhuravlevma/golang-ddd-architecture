package models

import (
	"errors"

	"github.com/google/uuid"
)

type Verfication struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	IsFull       bool
	Completed    bool
	Signed       bool
	ReportId     uuid.UUID
	ReportNumber int
}

func (report *Verfication) SignReport() error {
	if report.Completed == true {
		return errors.New("Cannot sign a report that has already been completed.")
	}
	report.Signed = true
	return nil
}

func (report *Verfication) CompleteVerification() error {
	if report.Signed == false {
		return errors.New("Cannot complete verification without signing the report.")
	}
	if report.ReportNumber < 0 {
		return errors.New("Report number cannot be negative.")
	}
	report.Completed = true
	return nil
}
