package emailService

import context "context"

type EmailServiceImpl struct {
}

func (e *EmailServiceImpl) SendEmail(context.Context, *EmailServiceRequest) (*EmailServiceResponse, error) {
	// Implement the logic for sending email
}
