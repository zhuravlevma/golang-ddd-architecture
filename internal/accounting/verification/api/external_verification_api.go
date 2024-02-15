package api

type ReportApiDto struct {
	Id string
	ReportNumber int
}

type ExternalVerificationApi struct {}

func (api *ExternalVerificationApi) SimpleVerifyReport(report *ReportApiDto) {}
func (api *ExternalVerificationApi) FullVerifyReport(report *ReportApiDto) {}
func (api *ExternalVerificationApi) SignReport(report *ReportApiDto) {}
func (api *ExternalVerificationApi) Complete(report *ReportApiDto) {}
